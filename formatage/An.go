package formatage

import (
	"go-reloeded/tools"
)

// change le mot "a" en "an" si le mot suivant comment par une voyelle ou "h"
func formatageOfA(phrase []string) []string {
	for i, cara := range phrase {
		if cara == "a" {
			if tools.BeforeVoy(phrase, i) {
				phrase[i] = "an"
			}
		}
		if cara == "A" {
			if tools.BeforeVoy(phrase, i) {
				phrase[i] = "An"
			}
		}
	}
	return phrase
}
