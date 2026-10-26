package formatage

import (
	"go-reloeded/tools"
	"strings"
)

// Appelle toutes les fonctions de formatage
func FormatageComplet(file []byte) string {
	f := strings.Fields(string(file))
	f = modif(f)
	f = formatageOfA(f)
	f2 := tools.SlicestringToString(f)
	f2 = formatageApostrophe(f2)
	f2 = formatage(f2)
	return f2
}
