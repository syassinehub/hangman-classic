package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "--startWith" {
		if len(os.Args) < 3 {
			fmt.Println("Veuillez fournir un nom de fichier pour charger la partie enregistrée.")
			return
		}
		word, wordArray, attempts, err := loadGame(os.Args[2])
		if err != nil {
			fmt.Println("Erreur lors du chargement du jeu :", err)
			return
		}
		guess(word, wordArray, attempts)
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Veuillez fournir un fichier de mots en tant qu'argument.")
		return
	}
	nomFichier := os.Args[1]
	motChoisi, err := motAuHasardDansFichier(nomFichier)
	if err != nil {
		fmt.Printf("Erreur : %v\n", err)
		os.Exit(1)
	}
	fmt.Println(motChoisi)
	guess(motChoisi, makeWordArray(len(motChoisi)), 10)
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

func makeWordArray(length int) []string {
	wordArray := make([]string, length)
	for i := range wordArray {
		wordArray[i] = "_"
	}
	return wordArray
}

func guess(word string, wordArray []string, attempts int) {
	numRevealed := len(word)/2 - 1
    revealedIndices := rand.Perm(len(word))
    revealedIndices = revealedIndices[:numRevealed]
    for _, index := range revealedIndices {
        wordArray[index] = string(word[index])
    }
    fmt.Println("Bonne chance, vous avez", attempts, "essais pour trouver le mot !!!")
    PrintWordArray(wordArray)

    var lettersGuessed []string

    for attempts > 0 {
        var guess string
        fmt.Print("\nChoix : ")
        fmt.Scanln(&guess)

        if guess == "STOP" {
            err := saveGame(word, wordArray, attempts)
            if err != nil {
                fmt.Println("Erreur lors de l'enregistrement du jeu :", err)
                return
            }
            fmt.Println("Jeu enregistré dans save.txt.")
            return
        }

        if len(guess) == 1 {
            letterGuessed := false
            if !in(lettersGuessed, guess) {
                lettersGuessed = append(lettersGuessed, guess)

                letterInWord := false
                for index, char := range word {
                    if string(char) == guess {
                        wordArray[index] = guess
                        letterGuessed = true
                        letterInWord = true
                    }
                }

                if !letterInWord {
                    attempts--
                    pose_hangman(10 - attempts)
                } else {
                    PrintWordArray(wordArray)
                }

                if !letterGuessed {
                    fmt.Println("Vous avez déjà deviné cette lettre. Essayez-en une différente.")
                }
            }
        } else {
            fmt.Println("Vous ne pouvez devinez qu'une lettre à la fois.")
        }

        if strings.Join(wordArray, "") == word {
            PrintWordArray(wordArray)
            fmt.Println("Félicitations ! Vous avez deviné le mot :", word)
            return
        }
    }

    fmt.Println("Vous avez épuisé tous vos essais. Le mot était :", word)
}

func in(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func saveGame(word string, wordArray []string, attempts int) error {
	gameState := struct {
		Word     string
		WordArray []string
		Attempts int
	}{
		Word:     word,
		WordArray: wordArray,
		Attempts: attempts,
	}

	jsonData, err := json.Marshal(gameState)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("save.txt", jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func loadGame(filename string) (string, []string, int, error) {
	jsonData, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", nil, 0, err
	}

	var gameState struct {
		Word     string
		WordArray []string
		Attempts int
	}

	err = json.Unmarshal(jsonData, &gameState)
	if err != nil {
		return "", nil, 0, err
	}

	return gameState.Word, gameState.WordArray, gameState.Attempts, nil
}

func PrintWordArray(array []string) {
	for _, str := range array {
		fmt.Printf("%s ", str)
	}
	fmt.Println()
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
