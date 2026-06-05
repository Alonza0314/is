package is

import (
	"unicode"
	"unicode/utf8"
)

// Upper reports whether input is an upper-case Unicode letter.
func Upper(input rune) bool {
	return unicode.IsUpper(input)
}

// Lower reports whether input is a lower-case Unicode letter.
func Lower(input rune) bool {
	return unicode.IsLower(input)
}

// AllUpper reports whether input is valid UTF-8 and all its cased letters are
// upper case. It returns false when input contains no cased letters.
func AllUpper(input string) bool {
	return allCase(input, unicode.IsUpper)
}

// AllLower reports whether input is valid UTF-8 and all its cased letters are
// lower case. It returns false when input contains no cased letters.
func AllLower(input string) bool {
	return allCase(input, unicode.IsLower)
}

func allCase(input string, matchesCase func(rune) bool) bool {
	if !utf8.ValidString(input) {
		return false
	}

	hasCasedLetter := false
	for _, r := range input {
		switch {
		case matchesCase(r):
			hasCasedLetter = true
		case unicode.IsUpper(r), unicode.IsLower(r), unicode.IsTitle(r):
			return false
		}
	}

	return hasCasedLetter
}
