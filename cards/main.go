package main

import "fmt"




func main(){


	fmt.Println("Hello world")

	//variable declaration
	
	var card string = "Ace"
	fmt.Println(card)

	card2 := "bingo"
	fmt.Println(card2)

	//function call 

	var card_test string= cardName()
	fmt.Println(card_test)


	//Slice test
	fmt.Println("------------------Slice print statements------------------")
	testCards := deck{"Ace", cardName()}

	fmt.Println(testCards)
	fmt.Println("------------------Print statements using deck.go------------------")

	cards := newDeck()


	//cards = append(cards, "six of spades")
	
	hand, remainingDeck := deal(cards, 5)

	cards.print()

	fmt.Println("---------Printing HAND---------")

	hand.print()

	fmt.Println("---------Printing Remaining deck ---------")
	remainingDeck.print()
	fmt.Println(remainingDeck.toString())
	remainingDeck.saveToFile("remaining")
	fmt.Println("Reading from file")
	newDeck := readFile("remaining")
	//newDeck.print()	
	newDeck.shuffle().print()
}

func cardName() string{
	return "five of daimonds"
}