package tools

import "strconv"

// renvoie un mot avec sa premiere lettre en majuscule
func Cap(mot string) string {
	result := []rune(mot)
	if result[0] <= 'z' && result[0] >= 'a' {
		result[0] = rune(result[0] - 'a' + 'A')
	}
	return string(result)
}

// change un nombre d'une base en base 16 (Hexadecimal)
func Hex(mot string) string {
	val, err := strconv.ParseInt(mot, 16, 64)
	if err != nil {
		return mot
	}
	return strconv.Itoa(int(val))
}

// change un nombre d'une base en base 2 (binaire)
func Bin(mot string) string {
	val, err := strconv.ParseInt(mot, 2, 64)
	if err != nil {
		return mot
	}
	return strconv.Itoa(int(val))
}

// met en majuscule le mot
func Up(mot string) string {
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
func Low(mot string) string {
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
