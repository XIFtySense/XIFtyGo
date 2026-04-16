package main

import (
	"encoding/json"
	"fmt"
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

	fmt.Println("XIFty version:", xifty.Version())
	encoded, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(encoded))
}
