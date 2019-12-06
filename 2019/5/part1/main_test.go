package main

import (
	"bytes"
	"testing"
)

func TestExec(t *testing.T) {
	for _, test := range []struct {
		input, want string
		inputValue  int
		output      []int
	}{
		{"1,0,0,0,99", "2,0,0,0,99", 0, []int{}},
		{"2,3,0,3,99", "2,3,0,6,99", 0, []int{}},
		{"2,4,4,5,99,0", "2,4,4,5,99,9801", 0, []int{}},
		{"1,1,1,4,99,5,6,0,99", "30,1,1,4,2,5,6,0,99", 0, []int{}},
		{"1,9,10,3,2,3,11,0,99,30,40,50", "3500,9,10,70,2,3,11,0,99,30,40,50", 0, []int{}},
		{"3,0,4,0,99", "4,0,4,0,99", 4, []int{4}},
		{"1002,4,3,4,33", "1002,4,3,4,99", 0, []int{}},
		{"1101,100,-1,4,0", "1101,100,-1,4,99", 0, []int{}},
	} {
		var b bytes.Buffer
		b.WriteString(test.input)
		got, output, err := exec(&b, test.inputValue)

		if err != nil {
			t.Fatalf("Failed with err: %s", err)
		}

		if got != test.want {
			t.Errorf("Got %s but expected: %s", got, test.want)
		}

		for i, v := range test.output {
			if v != output[i] {
				t.Errorf("Got %d output but expected: %d", output[i], v)
			}
		}
	}
}
