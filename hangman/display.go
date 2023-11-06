package hangman

import (
	"fmt"
	"os"
)

var choice int
var choix int

func DrawWelcome() {
	fmt.Println(`
     ⠰⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠘⢿⣿⣆⠀⠀⠀⠀⠀⠀⠀⠀⠈⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠈⢻⣿⣇⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⣤⣶⣤⣤⣤⣤⣄⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⠛⠛⠛⠛⠻⠿⠿⣿⣿⣷⣦⣄⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⡀⣀⠀⠀⠀⠀⠀⠀⠈⠙⠻⠿⣿⣷⣦⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⣠⣾⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⣿⣿⣿⣿⣿⣷⣶⣦⣀⡀⠀⠀⠈⠛⢿⣿⣶⣄⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⢀⣾⣿⡿⠋⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⡐⢆⠲⣈⠭⣉⠛⠿⢿⣿⣷⣤⡀⠀⠀⠙⢿⣿⣧⡀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⣰⣿⡿⠋⠀⠀⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⡜⣌⠳⣌⠲⣡⢋⡜⢢⢍⠻⢿⣿⣷⣄⠀⠀⠙⣿⣿⣆⠀⠀⠀⠀
⠀⠀⠀⣰⣿⡿⠁⠀⠀⣴⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⡒⣌⠳⣌⠓⢦⠓⡜⢣⢎⢣⡃⢟⢿⣿⣦⠀⠀⠈⢻⣿⣆⠀⠀⠀
⠀⠀⣸⣿⡿⠀⠀⢀⣾⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⡱⣊⠵⣊⡝⢎⡹⡜⢣⢎⡳⢜⡣⢎⣻⣿⣷⡀⠀⠈⢿⣿⣇⠀⠀
⠀⢠⣿⣿⠃⠀⠀⣾⣿⢏⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⡱⣍⢮⡑⣎⠧⣓⡜⣣⠞⡜⣬⢱⢫⡔⣻⣿⣧⡀⠀⠈⣿⣿⡄⠀
⠀⣼⣿⡟⠀⠀⣸⣿⡿⣸⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⡱⢎⠶⣙⢦⡛⡴⣙⢦⡛⡜⢦⣋⠶⡹⣔⣻⣿⣇⠀⠀⢸⣿⣷⠀
⢀⣿⣿⠅⠀⠀⣿⣿⣇⢳⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⠵⣋⡞⡱⣎⠵⣣⠝⣦⢹⡜⣣⢎⡳⢵⢪⡼⣿⣿⡀⠀⠈⣿⣿⡀
⠈⣿⣿⠀⠀⠀⣿⣿⢬⢳⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⣋⢶⡩⢗⡜⣣⢧⡛⡴⣓⠮⣕⢮⠳⡭⢖⡳⣿⣿⡅⠀⠀⣿⣿⠆
⠰⣿⣿⡀⠀⠀⣿⣿⣎⢳⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⡱⢎⡵⢫⠼⡱⢎⡵⢣⡝⡺⡜⣎⢳⡙⢮⣱⣿⣿⠂⠀⠀⣿⣿⠅
⠀⣿⣿⡆⠀⠀⣿⣿⣎⢳⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⣱⢫⡜⣣⢏⡵⢫⡜⣣⢞⡱⡹⣌⢧⡙⢧⢺⣿⡿⠀⠀⢸⣿⡿⠀
⠀⢹⣿⣿⠀⠀⠘⣿⣿⣪⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⢆⡳⣜⡱⢎⡜⣣⢞⡱⢎⡵⠳⡜⢦⡛⣼⣿⣿⠇⠀⠀⣾⣿⡟⠀
⠀⠀⣿⣿⣇⠀⠀⠹⣿⣷⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⢎⠵⣪⢜⢣⠞⣡⢎⡵⢋⡴⢫⡜⡣⣝⣾⣿⠟⠀⠀⣰⣿⡿⠀⠀
⠀⠀⠈⢿⣿⡄⠀⠀⠹⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⠎⡵⢢⢫⡜⢎⡵⢊⠶⣩⢒⣇⡚⣵⣿⣿⠏⠀⠀⣠⣿⣿⠃⠀⠀
⠀⠀⠀⠘⢿⣿⣤⠀⠀⠈⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⡹⢜⣣⠳⡜⣣⠞⣩⠞⣡⠳⣼⣾⣿⡿⠁⠀⠀⣴⣿⡿⠁⠀⠀⠀
⠀⠀⠀⠀⠈⠻⣿⣧⣀⠠⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⣿⣿⣼⣧⣶⣽⣼⣵⣮⣵⣯⣶⣿⣿⣿⣯⣤⣤⣤⣾⣿⣿⣥⣤⣤⣄⡀
⠀⠀⠀⠀⠀⠀⠘⢿⣿⣧⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⠛⣿⣿⣿⠟⠁
⠀⠀⠀⠀⠀⠀⠀⠀⠉⣻⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣴⣿⠿⠋⠀⠀
⠀⠀⠀⠀⠀⠀⠀⢀⣾⣿⡿⠃⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣴⣿⡿⠋⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠰⣾⣟⣾⣽⣻⣟⣿⣻⣟⣿⣻⣟⣿⣻⣟⣿⣻⣟⣿⣻⣟⣿⣻⣟⣿⣻⣟⣿⣻⣟⡿⣿⣻⠝⠁⠀⠀⠀⠀`)
}

func Menu(g *Game) {
	fmt.Println("-----------------------")
	fmt.Println("1. Commencer à jouer")
	fmt.Println("2. Quitter")

	fmt.Scanln(&choice)

	switch choice {
	case 1:
		Draw(g, "")
	case 2:
		os.Exit(1)

	default:
		fmt.Println("Choix non valide. Veuillez choisir une autre option.")
	}
}

func Draw(g *Game, guess string) {
	drawTurns(g.TurnsLeft)
	drawState(g, guess)
	fmt.Println()
}

func drawTurns(l int) {
	var draw string
	switch l {
	case 1:
		draw = `
        +---+  
        |   |  
        O   |  
       /|\  |  
       / \  |  
            |  
	  =========
		`
	case 2:
		draw = `
        +---+  
        |   |  
        O   |  
       /|\  |  
       /    |  
            |  
	  =========
		`
	case 3:
		draw = `
        +---+  
        |   |  
        O   |  
       /|\  |  
            |  
            |  
	  =========
		`
	case 4:
		draw = `
        +---+  
        |   |  
    	O   |  
       /|   |  
            |  
            |  
      =========
			`
	case 5:
		draw = `
        +---+  
        |   |  
        O   |  
        |   |  
            |  
            |  
     =========
		`
	case 6:
		draw = `
        +---+  
        |   |  
        O   |  
            |  
            |  
            |  
      =========
		`
	case 7:
		draw = `
	+---+  
	|   |  
	    |  
	    |  
	    |  
	    |  
    =========
		`
	case 8:
		draw = `
    +---+  
	|  
	|  
	|  
	|  
	|  
  =========
		`
	case 9:
		draw = `
	|  
	|  
	|  
	|  
	|  
  =========
		`
	case 10:
		draw = `
	=========
	`
	case 11:
		draw = `

`
	}
	fmt.Println(draw)
}

func drawLetters(g []string) {
	for _, c := range g {
		fmt.Printf("%v ", c)
	}
	fmt.Println()
}

func drawState(g *Game, guess string) {
	fmt.Print("Déjà trouvé: ")
	drawLetters(g.FoundLetters)

	fmt.Print("Utilisée: ")
	drawLetters(g.UsedLetters)
	fmt.Println()

	switch g.State {
	case "goodGuess":
		fmt.Print("\x1b[32m GG! \x1b[0m")
	case "alreadyGuessed":
		fmt.Printf("T'as déjà utilisé le '%s' chef", guess)
	case "badGuess":
		fmt.Printf("\x1b[31m Mauvaise réponse, '%s' n'est pas dans le mot \x1b[0m", guess)
	case "lost":
		fmt.Print("\x1b[31m Mip Moop :( ! Le mot était : \x1b[0m")
		drawLetters(g.Letters)
	case "won":
		fmt.Print("\x1b[32m GG WP! Jgl canyon turbo sulfuro astro gap: \x1b[0m ")
		drawLetters(g.Letters)
	}

}
