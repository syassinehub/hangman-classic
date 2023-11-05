package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {
	nomFichier := os.Args[1]
	motChoisi, err := motAuHasardDansFichier(nomFichier)
	if err != nil {
		fmt.Printf("Erreur : %v\n", err)
		os.Exit(1)
	}
	guess(motChoisi)
}

var lines []string

func tab_hangman() {
	originalFile, err := os.Open("../hangman.txt")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier d'origine :", err)
		return
	}
	defer originalFile.Close()

	scanner := bufio.NewScanner(originalFile)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Erreur lors de la lecture du fichier d'origine :", err)
		return
	}
}

func pose_hangman(nb int) {
	const groupSize = 8
	start := (nb - 1) * groupSize
	end := nb * groupSize
	if start < 0 || end > len(lines) {
		fmt.Println("Groupe invalide.")
		return
	}
	for i := start; i < end; i++ {
		fmt.Println(lines[i])
	}
}

func motAuHasardDansFichier(nomFichier string) (string, error) {
	fichier, err := os.Open(nomFichier)
	if err != nil {
		return "", err
	}
	defer fichier.Close()

	scanner := bufio.NewScanner(fichier)

	var mots []string
	for scanner.Scan() {
		mots = append(mots, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}
	motAuHasard := mots[rand.Intn(len(mots))]
	return motAuHasard, nil
}

func guess(word string) {
	tab_hangman()
	a := len(word)
	var word_array []string
	for i := 0; i < a; i++ {
		word_array = append(word_array, "_")
	}
	if len(os.Args) != 2 {
		fmt.Println("Utilisation : ./main.go <nom_fichier>")
		os.Exit(1)
	}
	fmt.Println("Good luck, you have 10 attempts to find the word !!!")
	essaiRestant := 10
	var lettersGuessed []string
	for essaiRestant > 0 {
		for i := 0; i < len(word_array); i++ {
			fmt.Print(word_array[i])
			fmt.Print(" ")
		}

		var guess string
		fmt.Print("\nGuess (letter or word): ")
		fmt.Scanln(&guess)

		if len(guess) == 1 {
			letterGuessed := false
			if !in(lettersGuessed, guess) {
				lettersGuessed = append(lettersGuessed, guess)

				for index, char := range word {
					if string(char) == guess {
						word_array[index] = guess
						letterGuessed = true
					}
				}

				if !letterGuessed {
					essaiRestant--
					pose_hangman(10 - essaiRestant)
				}
			} else {
				fmt.Println("You already guessed that letter. Try a different one.")
			}
		} else if len(guess) >= 2 {
			if guess == word {
				fmt.Println("Congratulations! You guessed the word:", word)
				os.Exit(0)
			} else {
				essaiRestant -= 2
				pose_hangman(10 - essaiRestant)
			}
		}

		if strings.Join(word_array, "") == word {
			fmt.Println("Congratulations! You guessed the word:", word)
			os.Exit(0)
		}
	}

	fmt.Println("You've exhausted all your attempts. The word was:", word)
}

func in(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

var lines2 []string

func tab_ascii() {
	originalFile, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier d'origine :", err)
		return
	}
	defer originalFile.Close()

	scanner := bufio.NewScanner(originalFile)
	for scanner.Scan() {
		line := scanner.Text()
		lines2 = append(lines2, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Erreur lors de la lecture du fichier d'origine :", err)
		return
	}

}
