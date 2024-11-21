package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Bosch")
	want := "Hello, Bosch"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
