package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	t.Run("Rectangle perimeter", func(t *testing.T) {
		got := Perimeter(Rectangle{12.0, 8.0})
		want := 40.0
		assertMessage(t, got, want)
	})

	t.Run("Circle perimeter", func(t *testing.T) {
		got := CirlePeri(Circle{float64(3.5)})
		want := 21.991148575128552
		assertMessage(t, got, want)
	})
}

func assertMessage(t testing.TB, got float64, want float64) {
	t.Helper()
	if want != got {
		t.Errorf("got %g,want %g", got, want)
	}

}

func TestArea(t *testing.T) {
	t.Run("Rectangle area", func(t *testing.T) {
		got := Area(Rectangle{5.0, 6.0})
		want := 30.0
		assertMessage(t, got, want)
	})

	t.Run("circle area", func(t *testing.T) {
		got := CirleArea(Circle{3.5})
		want := 38.48451000647496

		assertMessage(t, got, want)
	})

}
