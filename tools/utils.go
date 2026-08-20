package tools

import "strconv"

// renvoi true si le mot est entre parentèse
func Notmot(mot string) bool {
	if len(mot) == 0 {
		return false
	}
	if mot[0] == '(' || mot[len(mot)-1] == ')' {
		return true
	}
	return false
}

// renvoie tout les chiffres d'un string en un nombre
func NombreInString(para string) int {
	var nombre []rune
	for i := 0; i < len(para); i++ {
		if para[i] <= '9' && para[i] >= '0' {
			nombre = append(nombre, rune(para[i]))
		}
	}
	val, _ := strconv.Atoi(string(nombre))
	return val
}

// enleve un mot d'un slice de string et le renvoie modifier
func Remove(phrase []string, place int) []string {
	var phrase2 []string
	for i := 0; i < len(phrase); i++ {
		if i != place {
			phrase2 = append(phrase2, phrase[i])
		}
	}
	return phrase2
}

// transform un slice de string en string
func SlicestringToString(s []string) string {
	var s2 string
	for index, cara := range s {
		if index != 0 && index != len(s) {
			s2 += " "
		}
		s2 += cara
	}
	return s2
}

// renvoie true si le caractere est une lettre
func IsLettre(r rune) bool {
	if r <= 'Z' && r >= 'A' {
		return true
	}
	if r <= 'z' && r >= 'a' {
		return true
	}
	return false
}

// renvoie true si la 1er lettre du mot suivant est une voyelle
func BeforeVoy(phrase []string, place int) bool {
	voyelle := "aeiouhAEIOUH"
	for _, voy := range voyelle {
		if rune(phrase[place+1][0]) == voy {
			return true
		}
	}
	return false
}

func RemoveCara(s string, place int) string {
	var result []rune
	for i, cara := range s {
		if i != place {
			result = append(result, cara)
		}
	}
	return string(result)
}
