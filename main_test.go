package main

import (
	"go-reloeded/formatage"
	"testing"
)

// fonction test
func TestFormatageComplet(t *testing.T) {
	cas := []struct {
		nom     string
		entree  string
		attendu string
	}{
		// --- (hex) ---
		{
			nom:     "hex simple depuis enonce",
			entree:  "1E (hex) fichiers ont été ajoutés",
			attendu: "30 fichiers ont été ajoutés",
		},
		{
			nom:     "hex minuscules",
			entree:  "ff (hex) points ont été marqués",
			attendu: "255 points ont été marqués",
		},
		{
			nom:    "hex en début de phrase",
			entree: "A (hex) était le score",

			attendu: "10 était le score",
		},
		// --- (bin) ---
		{
			nom:     "bin simple depuis enonce",
			entree:  "Ça fait 10 (bin) ans",
			attendu: "Ça fait 2 ans",
		},
		{
			nom:     "bin plus long",
			entree:  "il reste 1010 (bin) items",
			attendu: "il reste 10 items",
		},
		{
			nom:     "bin/hex combinés (exemple usage)",
			entree:  "Simply add 42 (hex) and 10 (bin) and you will see the result is 68.",
			attendu: "Simply add 66 and 2 and you will see the result is 68.",
		},

		// --- (up) sans count ---
		{
			nom:     "up simple depuis enonce",
			entree:  "Ready, set, go (up) !",
			attendu: "Ready, set, GO!",
		},

		// --- (up) avec count ---
		{
			nom:     "up avec count 2",
			entree:  "This is so exciting (up, 2)",
			attendu: "This is SO EXCITING",
		},
		{
			nom:     "up avec count 1 (equiv sans count)",
			entree:  "hello world (up, 1)",
			attendu: "hello WORLD",
		},

		// --- (low) sans count ---
		{
			nom:     "low simple depuis enonce",
			entree:  "I should stop SHOUTING (low)",
			attendu: "I should stop shouting",
		},

		// --- (low) avec count ---
		{
			nom:     "low avec count 3",
			entree:  "IT WAS THE (low, 3) winter of despair",
			attendu: "it was the winter of despair",
		},

		// --- (cap) sans count ---
		{
			nom:     "cap simple depuis enonce",
			entree:  "Welcome to the Brooklyn bridge (cap)",
			attendu: "Welcome to the Brooklyn Bridge",
		},

		// --- (cap) avec count ---
		{
			nom:     "cap avec count 6",
			entree:  "it was the age of foolishness (cap, 6)",
			attendu: "It Was The Age Of Foolishness",
		},

		// --- exemple complet de l'énoncé (le gros test) ---
		{
			nom: "phrase complete usage",
			entree: "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, " +
				"it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, " +
				"it was the season of Light, it was the season of darkness, it was the spring of hope, " +
				"IT WAS THE (low, 3) winter of despair.",
			attendu: "It was the best of times, it was the worst of TIMES, it was the age of wisdom, " +
				"It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, " +
				"it was the season of Light, it was the season of darkness, it was the spring of hope, " +
				"it was the winter of despair.",
		},

		// --- ponctuation simple ---
		{
			nom:     "ponctuation espace avant virgule",
			entree:  "I was sitting over there ,and then BAMM !!",
			attendu: "I was sitting over there, and then BAMM!!",
		},
		{
			nom:     "ponctuation espaces multiples avant point",
			entree:  "Il fait beau   .   On sort ?",
			attendu: "Il fait beau. On sort?",
		},
		{
			nom:     "ponctuation deux points et point virgule",
			entree:  "Voici la liste :pommes ;poires ;bananes.",
			attendu: "Voici la liste: pommes; poires; bananes.",
		},

		// --- groupes de ponctuation ... et !? ---
		{
			nom:     "points de suspension deja corrects",
			entree:  "Je pensais... Tu avais raison",
			attendu: "Je pensais... Tu avais raison",
		},
		{
			nom:     "points de suspension mal espacés",
			entree:  "Je pensais . . . Tu avais raison",
			attendu: "Je pensais... Tu avais raison",
		},
		{
			nom:     "groupe !? mal espacé",
			entree:  "Quoi ! ? tu es sérieux",
			attendu: "Quoi!? tu es sérieux",
		},
		{
			nom:     "exemple usage points suspension",
			entree:  "Punctuation tests are ... kinda boring ,what do you think ?",
			attendu: "Punctuation tests are... kinda boring, what do you think?",
		},

		// --- apostrophes / guillemets simples ---
		{
			nom:     "apostrophe deja correcte",
			entree:  "Je suis exactement comme ils me décrivent : 'génial'",
			attendu: "Je suis exactement comme ils me décrivent: 'génial'",
		},
		{
			nom:     "apostrophe avec espaces a nettoyer",
			entree:  "Je suis exactement comme ils me décrivent : ' génial '",
			attendu: "Je suis exactement comme ils me décrivent: 'génial'",
		},
		{
			nom:     "apostrophe multi-mots avec elisions (exemple enonce)",
			entree:  "Comme l'a dit Elton John : ' Je suis l'homosexuel le plus connu au monde '",
			attendu: "Comme l'a dit Elton John: 'Je suis l'homosexuel le plus connu au monde'",
		},
		{
			nom:     "deux paires de guillemets dans la meme phrase",
			entree:  "il a dit : ' ok ' puis ' stop '",
			attendu: "il a dit: 'ok' puis 'stop'",
		},
		{
			nom:     "elision seule sans guillemet",
			entree:  "qu'il vienne demain",
			attendu: "qu'il vienne demain",
		},

		// --- a / an ---
		{
			nom:     "a devient an devant voyelle (exemple enonce)",
			entree:  "There it was. A amazing rock!",
			attendu: "There it was. An amazing rock!",
		},
		{
			nom:     "a inchange devant consonne",
			entree:  "I want a banana",
			attendu: "I want a banana",
		},
		{
			nom:     "a devient an devant h",
			entree:  "I want a house",
			attendu: "I want an house",
		},
		{
			nom:     "A majuscule devient An",
			entree:  "A elephant walked in",
			attendu: "An elephant walked in",
		},

		// --- cas limites génériques ---
		{
			nom:     "chaine vide",
			entree:  "",
			attendu: "",
		},
		{
			nom:     "un seul mot",
			entree:  "bonjour",
			attendu: "bonjour",
		},
		{
			nom:     "espaces multiples entre mots normaux",
			entree:  "salut    tout   le monde",
			attendu: "salut tout le monde",
		},
	}

	for _, c := range cas {
		t.Run(c.nom, func(t *testing.T) {
			resultat := formatage.FormatageComplet([]byte(c.entree))
			if resultat != c.attendu {
				t.Errorf("entrée:   %q\nattendu:  %q\nobtenu:   %q", c.entree, c.attendu, resultat)
			}
		})
	}
}
