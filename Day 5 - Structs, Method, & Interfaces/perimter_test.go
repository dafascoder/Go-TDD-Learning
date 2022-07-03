package main

import "testing"



func TestPerimeter(t *testing.T){
	rectangle := Rectangle{10.0,10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("want %.2f but got %.2f", got, want)
	}
}

func TestArea(t *testing.T){

	/* checkArea := func(t testing.TB, shape Shape, want float64){
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("want %g but got %g", got , want)
		}
	}


	t.Run("circles", func(t* testing.T){
		circle := Circle{10}
		want := 314.1592653589793
		checkArea(t,circle,want)
	})

	t.Run("Rectangle", func(t* testing.T){
		rectangle := Rectangle{10.0,10.0}
		want := 100.0

		checkArea(t,rectangle,want)
	})
 */

	areaTests := []struct {
		name string
		shape Shape
		hasArea float64
	}{
		{name: "Rectangle",shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle",shape: Circle{Radius: 10},hasArea: 314.1592653589793},
		{name: "Triangle",shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
	  t.Run(tt.name, func(t *testing.T){
		got := tt.shape.Area()
		if got != tt.hasArea {
			t.Errorf("%#v got %g but want %g",tt.shape,got, tt.hasArea)
		}	
	  })
		
	}
	

}