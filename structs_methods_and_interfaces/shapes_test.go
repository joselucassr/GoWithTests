package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	expected := 40.0

	assertFloats(t, expected, got)
}

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		got := rectangle.Area()
		expected := 72.0

		assertFloats(t, expected, got)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		expected := 314.1592653589793

		assertFloats(t, expected, got)
	})
}

func assertFloats(t testing.TB, expected, got float64) {
	t.Helper()

	if expected != got {
		t.Errorf("expected %g but got %g", expected, got)
	}
}
