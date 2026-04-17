package xifty

import (
	"path/filepath"
	"testing"
)

func fixture(name string) string {
	return filepath.Join("fixtures", name)
}

func fieldsByName(output map[string]any) map[string]map[string]any {
	normalized := output["normalized"].(map[string]any)
	fields := normalized["fields"].([]any)
	byName := map[string]map[string]any{}
	for _, fieldAny := range fields {
		field := fieldAny.(map[string]any)
		byName[field["field"].(string)] = field
	}
	return byName
}

func TestVersionLooksSemantic(t *testing.T) {
	if Version() == "" {
		t.Fatal("expected non-empty version")
	}
}

func TestProbeReturnsInputSummary(t *testing.T) {
	output, err := Probe(fixture("happy.jpg"))
	if err != nil {
		t.Fatalf("probe failed: %v", err)
	}

	input := output["input"].(map[string]any)
	if output["schema_version"] != "0.1.0" {
		t.Fatalf("unexpected schema version: %v", output["schema_version"])
	}
	if input["detected_format"] != "jpeg" {
		t.Fatalf("unexpected format: %v", input["detected_format"])
	}
	if input["container"] != "jpeg" {
		t.Fatalf("unexpected container: %v", input["container"])
	}
}

func TestExtractDefaultsToFullEnvelope(t *testing.T) {
	output, err := Extract(fixture("happy.jpg"), ViewFull)
	if err != nil {
		t.Fatalf("extract failed: %v", err)
	}

	if output["raw"] == nil || output["interpreted"] == nil || output["normalized"] == nil || output["report"] == nil {
		t.Fatal("expected full envelope sections")
	}
}

func TestRawViewPreservesMetadataEvidence(t *testing.T) {
	output, err := Extract(fixture("happy.jpg"), ViewRaw)
	if err != nil {
		t.Fatalf("extract failed: %v", err)
	}

	raw := output["raw"].(map[string]any)
	containers := raw["containers"].([]any)
	metadata := raw["metadata"].([]any)
	if containers[0].(map[string]any)["label"] != "jpeg" {
		t.Fatalf("unexpected container label: %v", containers[0])
	}
	if metadata[0].(map[string]any)["tag_name"] != "ImageWidth" {
		t.Fatalf("unexpected first tag: %v", metadata[0])
	}
}

func TestInterpretedViewExposesDecodedTags(t *testing.T) {
	output, err := Extract(fixture("happy.jpg"), ViewInterpreted)
	if err != nil {
		t.Fatalf("extract failed: %v", err)
	}

	entries := output["interpreted"].(map[string]any)["metadata"].([]any)
	names := map[string]bool{}
	for _, entryAny := range entries {
		entry := entryAny.(map[string]any)
		names[entry["tag_name"].(string)] = true
	}
	if !names["Make"] || !names["Model"] || !names["DateTimeOriginal"] {
		t.Fatalf("missing expected interpreted tags: %v", names)
	}
}

func TestNormalizedViewReturnsExpectedFields(t *testing.T) {
	output, err := Extract(fixture("happy.jpg"), ViewNormalized)
	if err != nil {
		t.Fatalf("extract failed: %v", err)
	}

	fields := fieldsByName(output)
	if fields["device.make"]["value"].(map[string]any)["value"] != "XIFtyCam" {
		t.Fatalf("unexpected make value: %v", fields["device.make"])
	}
	if fields["device.model"]["value"].(map[string]any)["value"] != "IterationOne" {
		t.Fatalf("unexpected model value: %v", fields["device.model"])
	}
	if fields["captured_at"]["value"].(map[string]any)["value"] != "2024-04-16T12:34:56" {
		t.Fatalf("unexpected captured_at: %v", fields["captured_at"])
	}
}

func TestReportViewStaysExplicitWhenEmpty(t *testing.T) {
	output, err := Extract(fixture("happy.jpg"), ViewReport)
	if err != nil {
		t.Fatalf("extract failed: %v", err)
	}

	report := output["report"].(map[string]any)
	if len(report["issues"].([]any)) != 0 || len(report["conflicts"].([]any)) != 0 {
		t.Fatalf("expected empty report: %v", report)
	}
}

func TestNamedViewSelectionWorks(t *testing.T) {
	output, err := ExtractNamed(fixture("happy.jpg"), "normalized")
	if err != nil {
		t.Fatalf("extract named failed: %v", err)
	}
	fields := fieldsByName(output)
	if fields["software"]["value"].(map[string]any)["value"] != "XIFtyTestGen" {
		t.Fatalf("unexpected software value: %v", fields["software"])
	}
}

func TestInvalidNamedViewReturnsError(t *testing.T) {
	if _, err := ExtractNamed(fixture("happy.jpg"), "bad-view"); err == nil {
		t.Fatal("expected error for invalid named view")
	}
}
