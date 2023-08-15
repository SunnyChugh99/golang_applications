package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

//create a new type deck which is slice of string (think as deck is a class which extends string)
// Ace Two Three , Spades, Hearts, Diamonds

type deck []string


func newDeck() deck {
	
	cards := deck{}

	cardValues := []string{"Ace", "Two", "Three", "Four"}
	cardSuites := []string{"Spades","Hearts", "Diamonds", "Clubs"}

	for _,suit := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}
	return cards
}


func deal(d deck, handSize int)  (deck, deck){
	return d[:handSize], d[handSize:]
}


func (d deck) toString() string {
	return strings.Join([]string(d),",")
}


func (d deck) print(){

	for i,card := range d{
		fmt.Println(i, card)
	}

}


func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}


func readFile(filename string) deck{

	content, err := (ioutil.ReadFile(filename))
	if err != nil {
		fmt.Println("Error - ", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(content), ","))
}


func (d deck) shuffle() deck{
		rand.Shuffle(len(d), func (i int, j int){
			d[i], d[j] = d[j], d[i]
		})
		
		return d 

}
