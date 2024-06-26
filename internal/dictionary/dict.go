package dict

import (
	"math/rand"
	"strings"
)

func GetPatchName(numberOfSelections int) string {
	word := ""
	for i := 0; i < numberOfSelections; i++ {
		word += words[rand.Intn(len(words))]
		word += " "
	}
	word = strings.TrimRight(word, " ")
	if len(word) > 16 {
		word = word[:16]
	}
	if len(word) < 16 {
		pad := 16 - len(word)
		for i := 0; i < pad; i++ {
			word += " "
		}
	}
	return word
}
