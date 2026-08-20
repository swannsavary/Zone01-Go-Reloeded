package formatage

import (
	"go-reloeded/tools"
)

// formate avec les bonnes regles la position des apostrophe
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
				if i != 0 && tools.IsLettre(rune(s[i-1])) && tools.IsLettre(rune(s[i+1])) {
					indexvu = append(indexvu, i)
					continue
				} else {
					for e := i + 1; e < len(s)-1; e++ {
						if s[e] == ' ' {
							s = tools.RemoveCara(s, e)
							e--
						} else {
							break
						}
					}
					for e := i + 1; e < len(s); e++ {
						if s[e] != '\'' {
							continue
						}
						if e != len(s)-1 && tools.IsLettre(rune(s[e-1])) && tools.IsLettre(rune(s[e+1])) {
							indexvu = append(indexvu, e)
							continue
						}
						for j := e - 1; j >= 0; j-- {
							if s[j] == ' ' {
								s = tools.RemoveCara(s, j)
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
