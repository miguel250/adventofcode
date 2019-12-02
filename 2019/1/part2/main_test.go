package main

import (
	"bufio"
	"bytes"
	"os"
	"testing"
)

func TestCalculateFuel(t *testing.T) {
	input, err := os.Open("testdata/input")

	if err != nil {
		t.Errorf("Failed to open input file with: %w", err)
	}

	defer input.Close()

	var b bytes.Buffer
	calculateFuel(input, &b)

	expected, err := os.Open("testdata/output")

	if err != nil {
		t.Errorf("Failed to open output file with: %w", err)
	}

	defer expected.Close()

	expectedScanner := bufio.NewScanner(expected)
	resultScanner := bufio.NewScanner(&b)

	for expectedScanner.Scan() && resultScanner.Scan() {
		resultValue := resultScanner.Text()
		expectedValue := expectedScanner.Text()

		if resultValue != expectedValue {
			t.Errorf("Values don't match expected %s got: %s", expectedValue, resultValue)
		}
	}
}
