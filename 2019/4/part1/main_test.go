package main

import (
	"testing"
)

func TestValidPasswordCount(t *testing.T) {
	count, err := validPasswordCount(130254, 678275)

	if err != nil {
		t.Errorf("failed to get count of valid passwords with %s", err)
	}

	if count != 2090 {
		t.Errorf("should had gotten 2090 got: (%d)", count)
	}
}

func TestValidPasswords(t *testing.T) {
	for _, test := range []struct {
		input int
		want  bool
	}{
		{111111, true},
		{223450, false},
		{123789, false},
		{123345, true},
		{677959, false},
		{444775, false},
	} {
		got := isValidPassword(test.input)
		if got != test.want {
			t.Errorf("password %d should be (%t) but got (%t)", test.input, test.want, got)
		}
	}
}
