package tests

import (
	"testing"

	"github.com/njmaeys/go_become_something/cmd/go_become_something/v1"
)

func TestDummyString(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"foo bar"}
	}

	for _, c := range cases {
		got := dummyString("blah")
		fmt.Println(got)
	}
}