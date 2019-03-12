package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Size of new deck is expected to be 52, but got %v", len(d))
	}
}
