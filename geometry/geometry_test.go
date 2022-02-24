package geometry

import "testing"

func TestGeometry(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	checkPerimeter := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}

	}
	t.Run("Test rectangle", func(t *testing.T) {
		rectangle := Rectangle{10, 20}
		checkPerimeter(t, rectangle, 60.0)
		checkArea(t, rectangle, 200.0)
	})

	t.Run("Test circle", func(t *testing.T) {
		rectangle := Circle{10}
		checkPerimeter(t, rectangle, 62.83185307179586)
		checkArea(t, rectangle, 314.1592653589793)
	})

}
