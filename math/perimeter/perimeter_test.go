package perimeter

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles perimeter", func(t *testing.T) {
		rectangle := Rectangle{10, 6}
		expect := 32.0

		checkPerimeter(t, rectangle, expect)
	})

	t.Run("circles perimeter", func(t *testing.T) {
		circle := Circle{10}
		expect := math.Pi * 20
		checkPerimeter(t, circle, expect)
	})
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		want := 72.0

		checkArea(t, rectangle, want)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		want := math.Pi * 100

		checkArea(t, circle, want)
	})
}

func TestAreaWithAnonymousStruct(t *testing.T) {

	areaTests := []struct {
		name   string
		shape  Shape
		expect float64
	}{
		{"Rectangle", Rectangle{12, 6}, 72.0},
		{"Circle", Circle{10}, 100 * math.Pi},
		{"SameSidedTriangle", SameSidedTriangle{12, 6}, 36.0},
	}

	for _, testCase := range areaTests {
		t.Run(testCase.name, func(t *testing.T) {
			got := testCase.shape.Area()
			if got != testCase.expect {
				t.Errorf("got %g want %g", got, testCase.expect)
			}
		})
	}
}
