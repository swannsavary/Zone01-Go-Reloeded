package main

import (
	"fmt"
	"go-reloeded/formatage"
	"log"
	"os"
)

// prend le fichier sample,le modifie puis renvoie le resultat dans result.txt
func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Error, need 2 arguments")
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
