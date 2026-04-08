package main
import ("testing"

)
func TestPerimeter(t *testing.T) {
	got:=Perimeter(Rectangle{12.0,8.0})
	want:=40.0
	assertMessage(t,got,want)
}

func assertMessage(t testing.TB,want float64,got float64){
	t.Helper()
	if want != got {
		t.Errorf("got %f,want %f",got,want)
	}

}

func TestArea(t *testing.T){
	got:=Area(Rectangle{5.0,6.0})
	want:=30.0
	assertMessage(t,got,want)
}

// Circle perimeter

func TestCirlePeri(t *testing.T){
	got:=CirlePeri(Circle{3.5})
	want:=21.994000

	assertMessage(t,got,want)
}

func TestCircleArea(t *testing.T){
	got:=CirleArea(Circle{3.5})
	want:=76.979000

	assertMessage(t,got,want)
}