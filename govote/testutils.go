package govote

import (
	"strconv"
	"testing"
)

//some simple function for testing the output
func testOutput(t *testing.T, result []string, expected []string) {
	if expected == nil && len(result) != 0 {
		t.Fatalf("We expected no output but got %+v", result)
	}
	if len(result) != len(expected) {
		t.Fatalf("The result has len %d we expected %d", len(result), len(expected))
	}
	for index, value := range result {
		if value != expected[index] {
			t.Errorf("Expected value on index %d was %s but got %s", index, expected[index], value)
		}
	}
}
