package itteration

import (
	"testing"
)

func TestRepeat(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Repeat("Roma", 3)
		want := "Roma Roma Roma"
		assertCorrectMessage(t, got, want)
	})
}
