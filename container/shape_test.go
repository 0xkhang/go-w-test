package container

import "testing"


func TestPerimeter(t *testing.T) {
    rect := Rectangle{10.0, 10.0}
    got := Perimeter(rect) 
    want := 40.0

    if got != want {
        t.Errorf("got %.2f want %.2f", got, want)
    }
}

func TestArea(t *testing.T) {
    got := Area(Rectangle{18.0, 35.0})
    want := 630.0

    if got != want {
        t.Errorf("got %.2f want %.2f", got, want)
    }
}
