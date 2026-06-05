package is

import (
	"unicode"
	"unicode/utf8"
)

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

// Palindrome reports whether input is valid UTF-8 and reads the same forward
// and backward. It compares the input exactly, including case, spaces, and
// punctuation.
func Palindrome(input string) bool {
	if !utf8.ValidString(input) {
		return false
	}

	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}

	return true
}

// AlnumPalindrome reports whether input is valid UTF-8 and reads the same
// forward and backward after ignoring non-alphanumeric runes and normalizing
// letters to lower case.
func AlnumPalindrome(input string) bool {
	if !utf8.ValidString(input) {
		return false
	}

	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; {
		if !unicode.IsLetter(runes[i]) && !unicode.IsDigit(runes[i]) {
			i++
			continue
		}

		if !unicode.IsLetter(runes[j]) && !unicode.IsDigit(runes[j]) {
			j--
			continue
		}

		if unicode.ToLower(runes[i]) != unicode.ToLower(runes[j]) {
			return false
		}

		i++
		j--
	}

	return true
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
