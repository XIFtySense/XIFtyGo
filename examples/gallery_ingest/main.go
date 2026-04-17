package main

import (
	"encoding/json"
	"log"
	"path/filepath"

	xifty "github.com/XIFtySense/XIFtyGo"
)

func main() {
	fixture := filepath.Join("fixtures", "happy.jpg")
	output, err := xifty.Extract(fixture, xifty.ViewNormalized)
	if err != nil {
		log.Fatal(err)
	}

	fields := map[string]any{}
	for _, fieldAny := range output["normalized"].(map[string]any)["fields"].([]any) {
		field := fieldAny.(map[string]any)
		value := field["value"].(map[string]any)
		fields[field["field"].(string)] = value["value"]
	}

	asset := map[string]any{
		"sourcePath":  fixture,
		"format":      output["input"].(map[string]any)["detected_format"],
		"capturedAt":  fields["captured_at"],
		"cameraMake":  fields["device.make"],
		"cameraModel": fields["device.model"],
		"width":       fields["dimensions.width"],
		"height":      fields["dimensions.height"],
		"software":    fields["software"],
	}

	encoded, err := json.MarshalIndent(asset, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	println(string(encoded))
}
