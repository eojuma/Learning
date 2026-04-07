package main

import "testing"

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
