package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Monark")
	want := "Hello, Monark"

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
