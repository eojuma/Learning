package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people",func(t *testing.T){

	got := Hello("Evans!")
	want := "Hello, Evans!"

	assertMessage(t,got,want)
})
t.Run("saying 'Hello,World',when an empty string is passed",func(t *testing.T){
	got:=Hello("")
	want:="Hello, World!"
		assertMessage(t,got,want)

})
}
func assertMessage(t testing.TB,got,want string){
	t.Helper()
if got !=want{
		t.Errorf("got %s,want %s",got,want)
	}
}

