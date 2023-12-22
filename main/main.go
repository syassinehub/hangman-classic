package main

import (
	"Hangman"
	"os"
	"fmt"
)

func main() {
	nomFichier := os.Args[1]
	motChoisi, err := Hangman.MotAuHasardDansFichier(nomFichier)
	if err != nil {
		fmt.Printf("Erreur : %v\n", err)
		os.Exit(1)
	}
	Hangman.Guess(motChoisi)
}

