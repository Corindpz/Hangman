package main

import (
	"fmt"
	"hangman/dictionary"
	"hangman/hangman"
	"os"
)

func main() {

	err := dictionary.Load("words.txt")
	if err != nil {
		fmt.Printf("Could not load dictionary: %v\n", err)
		os.Exit(1)
	}

	g, err := hangman.New(11, dictionary.PickWord())
	if err != nil {
		fmt.Printf("Could not create game: %v\n", err)
		os.Exit(1)
	}

	hangman.DrawWelcome()
	hangman.Menu(g)
	fmt.Print("\033[H\033[2J")
	guess := ""

	gameOver := false

	for !gameOver {
		hangman.Draw(g, guess)

		switch g.State {
		case "won", "lost":
			gameOver = true

			var replay string
			fmt.Println("Voulez-vous rejouer ? (oui/non)")
			fmt.Scanln(&replay)

			if replay == "oui" {
				gameOver = false
				main()
			} else {
				os.Exit(0)
			}
		}

		l, err := hangman.ReadGuess()
		if err != nil {
			fmt.Printf("Could not read from terminal: %v", err)
			os.Exit(1)
		}
		guess = l

		g.MakeAGuess(guess)
		fmt.Print("\033[H\033[2J")
	}
}
