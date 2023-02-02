package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{17.0, 3.0}

	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, but wanted %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	t.Parallel() // marks TLog as capable of running in parallel with other tests

	//Table driven test, represented as a map so we can specify a name and test cases
	areaTests := map[string]struct {
		shape Shape
		want  float64
	}{
		"testing Area of rectangle": {
			shape: Rectangle{17.0, 3.0},
			want:  51.0,
		},
		"testing Area of circle": {
			shape: Circle{10.0},
			want:  314.1592653589793,
		},
	}

	//helper function
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()

		got := shape.Area()

		if got != want {
			t.Errorf("got %g, but wanted %g", got, want)
		}
	}

	//actual test runs
	for name, test := range areaTests {
		t.Run(name, func(t *testing.T) {
			// NOTE: https://github.com/golang/go/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables
			// copy variables since in go the address of a loop variable remains the same, leaving its value as the it was on the last iteration
			name, test := name, test

			t.Parallel()
			checkArea(t, test.shape, test.want)
			t.Log(name)
		})
	}
}
