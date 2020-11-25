package main

import (
	"testing"
)

func TestDummyString(t *testing.T) {
	cases := []struct {
		in   string
		want DummyStringVal
	}{
		{"foo bar", DummyStringVal{"foo bar"}},
	}

	for _, c := range cases {
		got := dummyString(c.in)
		if got != c.want {
			t.Errorf("dummString(%q) got %q, should be %q", c.in, got, c.want)
		}
	}
}
