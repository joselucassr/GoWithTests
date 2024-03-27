package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	expected := 40.0

	assertFloats(t, expected, got)
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the ˋt.Runˋ test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if tt.hasArea != got {
				t.Errorf("%#v expected %g but got %g", tt.shape, tt.hasArea, got)
			}
		})
	}
}

func assertFloats(t testing.TB, expected, got float64) {
	t.Helper()

	if expected != got {
		t.Errorf("expected %g but got %g", expected, got)
	}
}
