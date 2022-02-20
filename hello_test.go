package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Roma")
	want := "Hello, Roma"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

