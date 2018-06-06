package main

import (
	"testing"
)

func TestExample(t *testing.T) {
	actual := 1
	expected := 1
    if actual != expected {
       t.Errorf("Failed, got: %d, want: %d.", actual, expected)
    }
}