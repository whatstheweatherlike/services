package geocode

import (
	"strings"
	"testing"
)

func TestReverse(t *testing.T) {
	result, err := Reverse("51.9530150,7.9908162")
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	if result == "" {
		t.Errorf("No result")
	}
	if result == "" {
		t.Errorf("No result")
	}
	if !strings.Contains(result, "Emsstraße 8, 48231 Warendorf, Germany") {
		t.Errorf("Unexpected value: %v", err)
	}
}
