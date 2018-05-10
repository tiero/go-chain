package main

import "testing"

func TestCalculateHash(t *testing.T) {
	want := "9af15b336e6a9619928537df30b2e6a2376569fcf9d7e773eccede65606529a0"
	got := calculateHash(0,"0", Transaction{}, 0)
	if got != want {
		t.Errorf("Got == %q, want %q", got, want)
	}
}