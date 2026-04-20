package main

import( "testing"
"reflect"
)
func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {

		got := Hello("Evans!", "english")
		want := "Hello, Evans!"

		assertMessage(t, got, want)
	})
	t.Run("saying 'Hello,World',when an empty string is passed", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertMessage(t, got, want)

	})

	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Eldoie", "spanish")
		want := "Hola, Eldoie"
		assertMessage(t, got, want)
	})

	t.Run("saying hello in french", func(t *testing.T) {
		got := Hello("Juma", "french")
		want := "Bonjour,Juma"
		assertMessage(t, got, want)
	})
}
func assertMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s,want %s", got, want)
	}
}

func TestAdd(t *testing.T) {
	got := Add(4, 5)
	want := 9

	if got != want {
		t.Errorf("got %d,want %d", got, want)
	}
}

func TestRepeat(t *testing.T) {
	got := Repeat()
	want := "AAAAA"

	assertMessage(t, got, want)
}

func TestSum(t *testing.T) {
	t.Run("testing the sum of five numbers", func(t *testing.T) {
		a := [5]int{1, 2, 3, 4, 5}
		got := Arrays(a)
		want := 15

		assertSum(t, got, want)
	})
}

func TestSlices(t *testing.T){
	a:=[]int{7,5,8}
	b:=[]int{1,2,3}
	got:=Slices(a,b)
	want:=[]int{20,6}
// go does not allow comparison of two slices 
// or arrays of different lengths so we use 
// reflect.DeepEqual from the reflect package to do the comparison
//it could also be used structs,maps
	if !reflect.DeepEqual(got,want) {
		t.Errorf("got %v,want %v",got,want)
	}
}




func assertSum(t testing.TB, got, want int) {
	if got != want {
		t.Errorf("got %d , want %d", got, want)
	}
}

