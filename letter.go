package is

import "unicode"

// Upper reports whether input is an upper-case Unicode letter.
func Upper(input rune) bool {
	return unicode.IsUpper(input)
}

// Lower reports whether input is a lower-case Unicode letter.
func Lower(input rune) bool {
	return unicode.IsLower(input)
}
