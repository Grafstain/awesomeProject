package geometry

import "testing"

func TestGeometry(t *testing.T) {
	t.Run("Test perimeter", func(t *testing.T) {
		r := Rectangle{10, 20}
		got := r.Perimeter()
		want := 60.0
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
	t.Run("Test square", func(t *testing.T) {
		r := Rectangle{10, 20}
		got := r.Square()
		want := 200.0
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
	t.Run("Test square circle", func(t *testing.T) {
		r := Circle{10}
		got := r.Square()
		want := 314.1592653589793
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
	t.Run("Test perimeter circle", func(t *testing.T) {
		r := Circle{10}
		got := r.Perimeter()
		want := 62.83185307179586
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}
