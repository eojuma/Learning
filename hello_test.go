package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello(" "+"Evans!")
	want := "Hello Evans!"

	if got != want {
		t.Errorf("got %s,want %s", got, want)
	}
}
