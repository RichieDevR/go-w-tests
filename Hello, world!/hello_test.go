package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Rich")
	want := "Hello, Rich"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
