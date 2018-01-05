// Implement the Schinke Latin stemming as described in
// http://snowball.tartarus.org/otherapps/schinke/intro.html
package schinke

import (
	"strings"
)

var queSuffixWords = map[string]bool{
	"absque":        true,
	"abusque":       true,
	"adaeque":       true,
	"adusque":       true,
	"apsque":        true,
	"atque":         true,
	"attorque":      true,
	"concoque":      true,
	"contorque":     true,
	"coque":         true,
	"cuique":        true,
	"cuiusque":      true,
	"decoque":       true,
	"denique":       true,
	"deque":         true,
	"detorque":      true,
	"excoque":       true,
	"extorque":      true,
	"incoque":       true,
	"intorque":      true,
	"itaque":        true,
	"neque":         true,
	"oblique":       true,
	"obtorque":      true,
	"optorque":      true,
	"peraeque":      true,
	"plenisque":     true,
	"praetorque":    true,
	"quaeque":       true,
	"quamque":       true,
	"quandoque":     true,
	"quaque":        true,
	"quarumque":     true,
	"quasque":       true,
	"quemque":       true,
	"quibusque":     true,
	"quique":        true,
	"quisque":       true,
	"quoque":        true,
	"quorumque":     true,
	"quosque":       true,
	"quotusquisque": true,
	"quousque":      true,
	"recoque":       true,
	"retorque":      true,
	"susque":        true,
	"torque":        true,
	"ubique":        true,
	"undique":       true,
	"usque":         true,
	"uterque":       true,
	"utique":        true,
	"utribique":     true,
	"utroque":       true,
}

var removeSuffixNoun = []string{"ibus", "ius", "ae", "am", "as", "em", "es", "ia", "is", "nt", "os", "ud", "um", "us", "a", "e", "i", "o", "u"}
var removeSuffixVerb = []string{"iuntur", "beris", "erunt", "untur", "iunt", "mini", "ntur", "stis", "bor", "ero", "mur", "mus", "ris", "sti", "tis", "tur", "unt", "bo", "ns", "nt", "ri", "m", "r", "s", "t"}

var convertSuffix = map[string]string{
	"iuntur": "i",
	"erunt":  "i",
	"untur":  "i",
	"iunt":   "i",
	"unt":    "i",
	"beris":  "bi",
	"bor":    "bi",
	"bo":     "bi",
	"ero":    "eri",
}

// Stem each word to two forms, noun and verbs
func Stem(s string) (string, string) {
	var noun, verb string

	s = strings.Replace(s, "j", "i", -1)
	s = strings.Replace(s, "v", "u", -1)

	if strings.HasSuffix(s, "que") {
		if queSuffixWords[s] {
			return s, s
		}
		s = strings.TrimSuffix(s, "que")
	}

	noun = s
	for _, suffix := range removeSuffixNoun {
		if strings.HasSuffix(noun, suffix) {
			noun = strings.TrimSuffix(noun, suffix)
			break
		}
	}

	verb = s
	for _, suffix := range removeSuffixVerb {
		if strings.HasSuffix(verb, suffix) {
			verb = strings.TrimSuffix(verb, suffix)
			if newSuffix, ok := convertSuffix[suffix]; ok {
				verb += newSuffix
			}
			break
		}
	}

	if len(noun) < 2 {
		noun = s
	}
	if len(verb) < 2 {
		verb = s
	}

	return noun, verb
}
