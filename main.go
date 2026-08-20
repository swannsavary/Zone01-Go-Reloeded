package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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
	f := formatageComplet(file)
	err3 := os.WriteFile(args[1], []byte(f), 0644)
	if err3 != nil {
		return
	}

}

// appel toutes les fonctions de formatage
func formatageComplet(file []byte) string {
	f := strings.Fields(string(file))
	f = modif(f)
	return formatage(f)
}

// modifie la forme de la phrase pour qu'elle soit conforme au regles
func formatage(phrase []string) string {
	var resultat []rune
	phrase2 := slicestringToString(formatageOfA(phrase))
	phrase2 = formatageApostrophe(phrase2)
	for index, cara := range phrase2 {
		if index == len(phrase2)-1 {
			resultat = append(resultat, cara)
			break
		}
		switch cara {
		case ' ':
			suivant := rune(phrase2[index+1])
			estPonctuation := suivant == '.' || suivant == ',' || suivant == ';' || suivant == '?' || suivant == '!' || suivant == ':'
			if suivant != ' ' && index != 0 && !estPonctuation {
				resultat = append(resultat, cara)
			}
		case ',', ';', '.', '!', '?', ':':
			resultat = append(resultat, cara)
			if isLettre(rune(phrase2[index+1])) {
				resultat = append(resultat, ' ')
			}
		default:
			resultat = append(resultat, cara)
		}
	}
	return string(resultat)
}

func formatageApostrophe(s string) string {
	dejavu := false
	indexvu := []int{}
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '\'' {
			for _, nb := range indexvu {
				if nb == i {
					dejavu = true
				}
			}
			if !dejavu {
				if i != 0 && isLettre(rune(s[i-1])) && isLettre(rune(s[i+1])) {
					indexvu = append(indexvu, i)
					continue
				} else {
					for e := i + 1; e < len(s)-1; e++ {
						if s[e] == ' ' {
							s = removeCara(s, e)
							e--
						} else {
							break
						}
					}
					for e := i + 1; e < len(s); e++ {
						if s[e] != '\'' {
							continue
						}
						if e != len(s)-1 && isLettre(rune(s[e-1])) && isLettre(rune(s[e+1])) {
							indexvu = append(indexvu, e)
							continue
						}
						for j := e - 1; j >= 0; j-- {
							if s[j] == ' ' {
								s = removeCara(s, j)
							} else {
								indexvu = append(indexvu, j+1)
								break
							}
						}
						break
					}
				}
			} else {
				dejavu = false
			}
		}
	}
	return s
}

func removeCara(s string, place int) string {
	var result []rune
	for i, cara := range s {
		if i != place {
			result = append(result, cara)
		}
	}
	return string(result)

}

// change le mot a en an si le mot suivant comment par une voyelle ou h
func formatageOfA(phrase []string) []string {
	for i, cara := range phrase {
		if cara == "a" {
			if beforeVoy(phrase, i) {
				phrase[i] = "an"
			}
		}
		if cara == "A" {
			if beforeVoy(phrase, i) {
				phrase[i] = "An"
			}
		}
	}
	return phrase
}

// renvoie true si la 1er lettre du mot suivant est une voyelle
func beforeVoy(phrase []string, place int) bool {
	voyelle := "aeiouhAEIOUH"
	for _, voy := range voyelle {
		if rune(phrase[place+1][0]) == voy {
			return true
		}
	}
	return false
}

// renvoie true si le caractere est une lettre
func isLettre(r rune) bool {
	if r <= 'Z' && r >= 'A' {
		return true
	}
	if r <= 'z' && r >= 'a' {
		return true
	}
	return false
}

// transform un slice de string en string
func slicestringToString(s []string) string {
	var s2 string
	for index, cara := range s {
		if index != 0 && index != len(s) {
			s2 += " "
		}
		s2 += cara
	}
	return s2
}

func modif(f []string) []string {
	for i := 0; i < len(f)-1; i++ {
		if notmot(f[i+1]) {
			if f[i+1][len(f[i+1])-1] != ')' {
				change(f, f[i+1]+f[i+2], i)
				f = remove(f, i+1)
				f = remove(f, i+1)
			} else if f[i+1][len(f[i+1])-1] == ')' && f[i+1][0] == '(' {
				change(f, f[i+1], i)
				f = remove(f, i+1)
			}
		}
	}
	return f
}

// enleve un mot d'un slice de string et le renvoie modifier
func remove(phrase []string, place int) []string {
	var phrase2 []string
	for i := 0; i < len(phrase); i++ {
		if i != place {
			phrase2 = append(phrase2, phrase[i])
		}
	}
	return phrase2
}

// renvoie tout les chiffres d'un string en un nombre
func nombreInString(para string) int {
	var nombre []rune
	for i := 0; i < len(para); i++ {
		if para[i] <= '9' && para[i] >= '0' {
			nombre = append(nombre, rune(para[i]))
		}
	}
	val, _ := strconv.Atoi(string(nombre))
	return val
}

// change un mot en fontion du parametre donner
func change(phrase []string, para string, place int) {
	if len(para) > 5 {
		parametre := para[:4]
		nombre := nombreInString(para)
		for i := place; i >= 0 && i >= place-nombre+1; i-- {
			switch parametre {
			case "(cap":
				phrase[i] = cap(phrase[i])
			case "(up,":
				phrase[i] = up(phrase[i])
			case "(low":
				phrase[i] = low(phrase[i])
			case "(bin":
				phrase[i] = bin(phrase[i])
			case "(hex":
				phrase[i] = hex(phrase[i])
			}
		}
	} else {
		switch para {
		case "(cap)":
			phrase[place] = cap(phrase[place])
		case "(up)":
			phrase[place] = up(phrase[place])
		case "(low)":
			phrase[place] = low(phrase[place])
		case "(bin)":
			phrase[place] = bin(phrase[place])
		case "(hex)":
			phrase[place] = hex(phrase[place])
		}
	}
}

// renvoi true si le mot est entre parentèse
func notmot(mot string) bool {
	if len(mot) == 0 {
		return false
	}
	if mot[0] == '(' || mot[len(mot)-1] == ')' {
		return true
	}
	return false
}

// renvoie un mot avec sa premiere lettre en majuscule
func cap(mot string) string {
	result := []rune(mot)
	if result[0] <= 'z' && result[0] >= 'a' {
		result[0] = rune(result[0] - 'a' + 'A')
	}
	return string(result)
}

// change un nombre d'une base en base 16
func hex(mot string) string {
	val, err := strconv.ParseInt(mot, 16, 64)
	if err != nil {
		return mot
	}
	return strconv.Itoa(int(val))
}

// change un nombre d'une base en base 2
func bin(mot string) string {
	val, err := strconv.ParseInt(mot, 2, 64)
	if err != nil {
		return mot
	}
	return strconv.Itoa(int(val))
}

// met en majuscule le mot
func up(mot string) string {
	var result []rune
	for _, cara := range mot {
		if cara <= 'z' && cara >= 'a' {
			result = append(result, cara-'a'+'A')
		} else {
			result = append(result, cara)
		}
	}
	return string(result)
}

// met en minuscule le mot
func low(mot string) string {
	var result []rune
	for _, cara := range mot {
		if cara <= 'Z' && cara >= 'A' {
			result = append(result, cara-'A'+'a')
		} else {
			result = append(result, cara)
		}
	}
	return string(result)
}
