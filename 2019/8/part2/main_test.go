package main

import (
	"bytes"
	"testing"
)

func TestParse(t *testing.T) {
	var b bytes.Buffer
	b.WriteString("123456789012")
	img := &image{wide: 3, tall: 2}
	err := img.parse(&b)

	if err != nil {
		t.Fatalf("Failed input with err: %s", err)
	}

	wantMaxLayer := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
	}

	for i := range wantMaxLayer {
		for j, v := range wantMaxLayer[i] {
			got := img.minZeroLayer[i][j]
			if v != got {
				t.Errorf("Got %d output but expected: %d", got, v)
			}
		}
	}

	wantLayers := [][][]int{
		[][]int{
			[]int{1, 2, 3},
			[]int{4, 5, 6},
		},
		[][]int{
			[]int{7, 8, 9},
			[]int{0, 1, 2},
		},
	}

	for i := range wantLayers {
		for j := range wantLayers[i] {
			for k, v := range wantLayers[i][j] {
				got := img.layers[i][j][k]
				if v != got {
					t.Errorf("Got %d output but expected: %d", got, v)
				}
			}
		}
	}
}

func TestCheckCorruption(t *testing.T) {
	var b bytes.Buffer
	b.WriteString("123456789012")
	img := &image{wide: 3, tall: 2}
	err := img.parse(&b)

	if err != nil {
		t.Fatalf("Failed input with err: %s", err)
	}
	got := img.checkCorruption()

	if got != 1 {
		t.Errorf("failed to check corruption got (%d) wants (2)", got)
	}
}

func TestDecode(t *testing.T) {
	var b bytes.Buffer
	b.WriteString("0222112222120000")
	img := &image{wide: 2, tall: 2}
	err := img.parse(&b)

	if err != nil {
		t.Fatalf("Failed input with err: %s", err)
	}

	want := " █\n█ \n"
	got := img.decode()

	if got != want {
		t.Errorf("want (%s) but got (%s)", want, got)
	}
}
