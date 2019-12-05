package main

import (
	"bytes"
	"testing"
)

func TestShortersDistance(t *testing.T) {
	for _, test := range []struct {
		input string
		want  int
	}{
		{"R8,U5,L5,D3\nU7,R6,D4,L4", 6},
		{"R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83", 159},
		{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7", 135},
	} {
		var b bytes.Buffer
		b.WriteString(test.input)

		p, err := New(&b)

		if err != nil {
			t.Errorf("failed to create panel with %s", err)
		}

		got := p.shortersDistance()

		if got != test.want {
			t.Errorf("Result don't match want: %d got: (%d)", test.want, got)
		}
	}
}
