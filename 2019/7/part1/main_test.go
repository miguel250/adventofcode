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
		{"3,9,8,9,10,9,4,9,99,-1,8", "3,9,8,9,10,9,4,9,99,1,8", 8, []int{1}},
		{"3,9,8,9,10,9,4,9,99,-1,8", "3,9,8,9,10,9,4,9,99,0,8", 1, []int{0}},
		{"3,9,7,9,10,9,4,9,99,-1,8", "3,9,7,9,10,9,4,9,99,1,8", 1, []int{1}},
		{"3,9,7,9,10,9,4,9,99,-1,8", "3,9,7,9,10,9,4,9,99,0,8", 9, []int{0}},
		{"3,3,1108,-1,8,3,4,3,99", "3,3,1108,1,8,3,4,3,99", 8, []int{1}},
		{"3,3,1107,-1,8,3,4,3,99", "3,3,1107,1,8,3,4,3,99", 7, []int{1}},
	} {
		var b bytes.Buffer
		b.WriteString(test.input)
		got, output, err := exec(&b, []int{test.inputValue})

		if err != nil {
			t.Fatalf("Failed input (%s) with err: %s", test.input, err)
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

func TestExecSequence(t *testing.T) {
	for _, test := range []struct {
		input, want        string
		inputValue, output []int
	}{
		{
			"3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0",
			"3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,43210,43210",
			[]int{4, 3, 2, 1, 0},
			[]int{43210},
		},
		{
			"3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0",
			"3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,54321,54320",
			[]int{0, 1, 2, 3, 4},
			[]int{54321},
		},
		{
			"3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0",
			"3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,65210,65210,0",
			[]int{1, 0, 4, 3, 2},
			[]int{65210},
		},
	} {
		var b bytes.Buffer
		b.WriteString(test.input)

		p := &program{}
		err := p.parser(&b)
		if err != nil {
			t.Fatalf("failed to parse opscode with %s\n", err)
		}

		got, output, err := execSequence(p, test.inputValue)

		if err != nil {
			t.Fatalf("Failed input (%s) with err: %s", test.input, err)
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
