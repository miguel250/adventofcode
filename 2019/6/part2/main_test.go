package main

import (
	"os"
	"testing"
)

func TestCountDirectAndIndirect(t *testing.T) {
	input, err := os.Open("testdata/input")

	if err != nil {
		t.Errorf("Failed to open input file with: %w", err)
	}

	defer input.Close()

	o := New(input)
	got := o.countDirectAndIndirect()

	if got != 42 {
		t.Errorf("wants 42 but got (%d)", got)
	}
}

func TestOrbitalTranferCount(t *testing.T) {
	input, err := os.Open("testdata/orbital-transfers")

	if err != nil {
		t.Errorf("Failed to open input file with: %w", err)
	}

	defer input.Close()

	o := New(input)
	got := o.orbitalTransferCount("YOU", "SAN")
	if got != 4 {
		t.Errorf("wants 42 but got (%d)", got)
	}
}
