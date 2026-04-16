package xifty

import (
	"path/filepath"
	"testing"
)

func fixture(name string) string {
	return filepath.Join("fixtures", name)
}

func TestVersionIsNonEmpty(t *testing.T) {
	if Version() == "" {
		t.Fatal("expected non-empty version")
	}
}

func TestProbeReturnsDetectedFormat(t *testing.T) {
	output, err := Probe(fixture("happy.jpg"))
	if err != nil {
		t.Fatalf("probe failed: %v", err)
	}

	input := output["input"].(map[string]any)
	if input["detected_format"] != "jpeg" {
		t.Fatalf("unexpected format: %v", input["detected_format"])
	}
}

func TestExtractNormalizedReturnsExpectedField(t *testing.T) {
	output, err := Extract(fixture("happy.jpg"), ViewNormalized)
	if err != nil {
		t.Fatalf("extract failed: %v", err)
	}

	normalized := output["normalized"].(map[string]any)
	fields := normalized["fields"].([]any)
	for _, fieldAny := range fields {
		field := fieldAny.(map[string]any)
		if field["field"] == "device.make" {
			value := field["value"].(map[string]any)
			if value["value"] != "XIFtyCam" {
				t.Fatalf("unexpected make value: %v", value["value"])
			}
			return
		}
	}

	t.Fatal("device.make field not found")
}

