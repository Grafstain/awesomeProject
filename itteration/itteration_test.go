package itteration

import (
	"fmt"
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
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
func ExampleRepeat() {
	res := Repeat("also", 2)
	fmt.Println(res)
	// Output: also also
}
