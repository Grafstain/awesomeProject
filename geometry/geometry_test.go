package geometry

import "testing"

func TestGeometry(t *testing.T) {
	areaTests := []struct {
		shape         Shape
		wantArea      float64
		wantPerimeter float64
	}{
		{Rectangle{10, 20}, 200.0, 60},
		{Circle{10}, 314.1592653589793, 62.83185307179586},
		{Triangle{4, 3}, 6.0, 12.0},
	}

	for _, tt := range areaTests {
		if got := tt.shape.Area(); got != tt.wantArea {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.wantArea)
		}
		if got := tt.shape.Perimeter(); got != tt.wantPerimeter {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.wantPerimeter)
		}
	}

}
