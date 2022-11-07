package structs

import "testing"

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Width: 10, Height: 6}, want: 60.0},
		{shape: Circle{Radius: 10}, want: 314.1592653589793},
		{shape: Triangle{Base: 10, Height: 6}, want: 30},
	}

	for _, test := range areaTests {
		t.Run(test.shape.GetName(), func(t *testing.T) {
			got := test.shape.Area()

			if got != test.want {
				t.Errorf("got %g want %g", got, test.want)
			}
		})
	}
}
