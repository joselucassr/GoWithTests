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
		shape    Shape
		expected float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if tt.expected != got {
			t.Errorf("expected %g but got %g", tt.expected, got)
		}
	}
}

func assertFloats(t testing.TB, expected, got float64) {
	t.Helper()

	if expected != got {
		t.Errorf("expected %g but got %g", expected, got)
	}
}
