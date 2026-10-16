package testing

import (
	"testing"
)

func TestAdd(t *testing.T) {
	result := Add(20, 10)
	expected := 30

	if result != expected {
		t.Errorf("Expected %d got %d", expected, result)
	}

}
