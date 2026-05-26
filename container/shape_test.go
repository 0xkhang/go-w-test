package container

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}
	got := rect.Perimiter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTest := []struct {
		shape Shape
		want  float64
	} {
        {Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
        // {Triangle{12, 6}, }
    }

    for _, tt := range areaTest {
        got := tt.shape.Area()
        if got != tt.want {
            t.Errorf("got %g want %g", got, tt.want)
        }
    }

	// checkArea := func(t testing.TB, shape Shape, want float64) {
	// 	t.Helper()
	// 	got := shape.Area()
	// 	if got != want {
	// 		t.Errorf("got %g want %g", got, want)
	// 	}
	// }
	//
	// t.Run("rectangles", func(t *testing.T) {
	// 	rect := Rectangle{12, 6}
	// 	checkArea(t, rect, 72.0)
	// })
	//
	// t.Run("circles", func(t *testing.T) {
	// 	c := Circle{10}
	// 	checkArea(t, c, 314.1592653589793)
	// })
}
