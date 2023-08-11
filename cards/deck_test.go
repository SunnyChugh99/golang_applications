package main

import (
	"os"
	"testing"
)


func TestNewDesk(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected length of 16 but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of ace of spaded but got %v", d[0])
	}
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of four of clubs but got %v", d[len(d)-1])
	}
}


func TestSaveToFilekAndReadFile(t *testing.T){
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := readFile("_decktesting")
	if len(loadedDeck) != 16 {
		t.Errorf("Expected length of 16 but got %v",len(loadedDeck))
	}
	os.Remove("_decktesting")
}