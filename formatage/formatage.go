package formatage

import (
	"go-reloeded/tools"
)

// modifie les espace et les ponctuation pour que la phrase soit comforme aux regles d'orthographe
func formatage(phrase string) string {
	var resultat []rune
	for index, cara := range phrase {
		if index == len(phrase)-1 {
			resultat = append(resultat, cara)
			break
		}
		switch cara {
		case ' ':
			suivant := rune(phrase[index+1])
			estPonctuation := suivant == '.' || suivant == ',' || suivant == ';' || suivant == '?' || suivant == '!' || suivant == ':'
			if suivant != ' ' && index != 0 && !estPonctuation {
				resultat = append(resultat, cara)
			}
		case ',', ';', '.', '!', '?', ':':
			resultat = append(resultat, cara)
			if tools.IsLettre(rune(phrase[index+1])) {
				resultat = append(resultat, ' ')
			}
		default:
			resultat = append(resultat, cara)
		}
	}
	return string(resultat)
}
