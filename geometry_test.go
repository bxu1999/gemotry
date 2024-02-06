package geometry

import (
        "testing"
)

func TestLength(t *testing.T) {

	point1 := Point{X: 3, Y: 4}
        expected := 5.0

        if got := point1.Length(); got != expected {
                t.Errorf("Point{X: 3, Y: 4}= %f, didn't return %f",  got, expected)
        }
}

