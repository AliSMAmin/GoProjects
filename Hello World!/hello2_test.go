package main

import "testing"

func TestHello2(t *testing.T) {
	got := Hello2()
	want := "Hello, world"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
