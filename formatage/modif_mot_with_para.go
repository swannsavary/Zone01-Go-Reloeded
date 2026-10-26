package formatage

import "go-reloeded/tools"

// Cherche dans la phrase donnée s'il existe un paramètre de la forme (paramètre), change les mots en fonction de celui-ci et supprime les paramètres de la phrase.
func modif(f []string) []string {
	for i := 0; i < len(f)-1; i++ {
		if tools.Notmot(f[i+1]) {
			if f[i+1][len(f[i+1])-1] != ')' {
				change(f, f[i+1]+f[i+2], i)
				f = tools.Remove(f, i+1)
				f = tools.Remove(f, i+1)
			} else if f[i+1][len(f[i+1])-1] == ')' && f[i+1][0] == '(' {
				change(f, f[i+1], i)
				f = tools.Remove(f, i+1)
			}
		}
	}
	return f
}

// Change un mot d'une phrase en fontion du parametre donner et de sa position
func change(phrase []string, para string, place int) {
	if len(para) > 5 {
		parametre := para[:4]
		nombre := tools.NombreInString(para)
		for i := place; i >= 0 && i >= place-nombre+1; i-- {
			switch parametre {
			case "(cap":
				phrase[i] = tools.Cap(phrase[i])
			case "(up,":
				phrase[i] = tools.Up(phrase[i])
			case "(low":
				phrase[i] = tools.Low(phrase[i])
			case "(bin":
				phrase[i] = tools.Bin(phrase[i])
			case "(hex":
				phrase[i] = tools.Hex(phrase[i])
			}
		}
	} else {
		switch para {
		case "(cap)":
			phrase[place] = tools.Cap(phrase[place])
		case "(up)":
			phrase[place] = tools.Up(phrase[place])
		case "(low)":
			phrase[place] = tools.Low(phrase[place])
		case "(bin)":
			phrase[place] = tools.Bin(phrase[place])
		case "(hex)":
			phrase[place] = tools.Hex(phrase[place])
		}
	}
}
