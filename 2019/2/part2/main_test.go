package main

import (
	"bytes"
	"testing"
)

func TestExec(t *testing.T) {
	for _, test := range []struct {
		input, want string
		expected    int
	}{
		{"1,0,0,0,99", "2,0,0,0,99", 0},
		{"2,3,0,3,99", "2,3,0,6,99", 0},
		{"2,4,4,5,99,0", "2,4,4,5,99,9801", 0},
		{"1,1,1,4,99,5,6,0,99", "30,1,1,4,2,5,6,0,99", 0},
		{"1,9,10,3,2,3,11,0,99,30,40,50", "3500,9,10,70,2,3,11,0,99,30,40,50", 0},
		{"1,1,1,0,99,5,6,0,99", "7,0,6,0,99,5,6,0,99", 7},
	} {
		var b bytes.Buffer
		b.WriteString(test.input)
		var (
			got string
			err error
		)
		if test.expected == 0 {
			got, err = exec(&b)
		} else {
			got, err = execExpected(&b, test.expected)
		}

		if err != nil {
			t.Fatalf("Failed with err: %w", err)
		}

		if got != test.want {
			t.Errorf("Got %s but expected: %s", got, test.want)
		}
	}
}
