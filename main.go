package main

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
	"time"
)

func motAleatoire(fichier string) (string, error) {

	fichierMots, err := os.Open(fichier)
	if err != nil {
		return "", err
	}
	defer fichierMots.Close()

	scanner := bufio.NewScanner(fichierMots)
	var mots []string
	for scanner.Scan() {
		mots = append(mots, scanner.Text())
	}

	rand.Seed(time.Now().Unix())
	motChoisi := mots[rand.Intn(len(mots))]

	return motChoisi, nil
}

func main() {
	/*mot, err := motAleatoire("words.txt")
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return
	}

	fmt.Println("Mot aléatoire choisi:", mot)*/

	content, _ := os.ReadFile("words.txt")
	lines := strings.Split(string(content), "\n")
	n1 := rand.Intn(len(lines) - 1)
	print(lines[n1])
}
