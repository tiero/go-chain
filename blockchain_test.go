package main

import "testing"

func TestCalculateHash(t *testing.T) {
	want := "460dcf402cac9cba808d32afad082caeb490f44f50a857545a8e857c70d700ea"
	got := calculateHash(0, 1, "0", Transaction{50 * 100000000, "COINBASE", "@tiero"}, 0)
	if got != want {
		t.Errorf("Got == %q, want %q", got, want)
	}
}
