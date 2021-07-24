package shapes

import (
	"math"
	"reflect"
	"testing"
)

func TestShape(t *testing.T) {
	//test table
	tests := []struct {
		shape Shape
		want  float64
	}{{
		shape: Rectangle{length: 4, width: 4}, want: 16.0,
	},
		{
			shape: Circle{radius: 1}, want: math.Pi,
		},
	}

	for _, test := range tests {
		got := test.shape.Area()
		assertshapearea(t, got, test.want)

	}

}

func assertshapearea(t testing.TB, got float64, want float64) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %g, want %g", got, want)
	}

}
