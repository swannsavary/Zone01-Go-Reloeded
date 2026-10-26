package main

import (
	"go-reloeded/formatage"
	"log"
	"os"
)

// prend le 1er fichier,le formate puis renvoie le resultat dans le second fichier
func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		os.Stdout.WriteString("Error, need 2 arguments\n")
		return
	}
	file, err := os.ReadFile(args[0])
	if err != nil {
		log.Fatal(err)
	}
	f := formatage.FormatageComplet(file)
	err3 := os.WriteFile(args[1], []byte(f), 0644)
	if err3 != nil {
		return
	}
}
