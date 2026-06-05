package is_test

import (
	"testing"

	"github.com/Alonza0314/is"
)

func TestLetter(t *testing.T) {
	testUpper(t)
	testLower(t)
	testAllUpper(t)
	testAllLower(t)
}

var testUpperCases = []struct {
	name     string
	input    rune
	expected bool
}{
	{name: "ASCII upper-case letter in is.Upper should return true", input: 'A', expected: true},
	{name: "Unicode upper-case letter in is.Upper should return true", input: 'É', expected: true},
	{name: "ASCII lower-case letter in is.Upper should return false", input: 'a', expected: false},
	{name: "Unicode lower-case letter in is.Upper should return false", input: 'é', expected: false},
	{name: "Unicode title-case letter in is.Upper should return false", input: 'ǅ', expected: false},
	{name: "Uncased letter in is.Upper should return false", input: '呱', expected: false},
	{name: "Digit in is.Upper should return false", input: '1', expected: false},
	{name: "Invalid rune in is.Upper should return false", input: rune(0x110000), expected: false},
}

func testUpper(t *testing.T) {
	for _, tc := range testUpperCases {
		t.Run(tc.name, func(t *testing.T) {
			result := is.Upper(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

var testLowerCases = []struct {
	name     string
	input    rune
	expected bool
}{
	{name: "ASCII lower-case letter in is.Lower should return true", input: 'a', expected: true},
	{name: "Unicode lower-case letter in is.Lower should return true", input: 'é', expected: true},
	{name: "ASCII upper-case letter in is.Lower should return false", input: 'A', expected: false},
	{name: "Unicode upper-case letter in is.Lower should return false", input: 'É', expected: false},
	{name: "Unicode title-case letter in is.Lower should return false", input: 'ǅ', expected: false},
	{name: "Uncased letter in is.Lower should return false", input: '呱', expected: false},
	{name: "Digit in is.Lower should return false", input: '1', expected: false},
	{name: "Invalid rune in is.Lower should return false", input: rune(0x110000), expected: false},
}

func testLower(t *testing.T) {
	for _, tc := range testLowerCases {
		t.Run(tc.name, func(t *testing.T) {
			result := is.Lower(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

var testAllUpperCases = []struct {
	name     string
	input    string
	expected bool
}{
	{name: "ASCII upper-case letters in is.AllUpper should return true", input: "ABC", expected: true},
	{name: "Unicode upper-case letters in is.AllUpper should return true", input: "ÉÖ", expected: true},
	{name: "Upper-case letters with uncased characters in is.AllUpper should return true", input: "ABC 123! 世界", expected: true},
	{name: "ASCII lower-case letters in is.AllUpper should return false", input: "abc", expected: false},
	{name: "Mixed-case letters in is.AllUpper should return false", input: "ABc", expected: false},
	{name: "Mixed Unicode case letters in is.AllUpper should return false", input: "Éö", expected: false},
	{name: "Title-case letter in is.AllUpper should return false", input: "ǅ", expected: false},
	{name: "Empty string in is.AllUpper should return false", input: "", expected: false},
	{name: "Uncased characters only in is.AllUpper should return false", input: "123 木頭人", expected: false},
	{name: "Invalid UTF-8 in is.AllUpper should return false", input: string([]byte{0xff}), expected: false},
}

func testAllUpper(t *testing.T) {
	for _, tc := range testAllUpperCases {
		t.Run(tc.name, func(t *testing.T) {
			result := is.AllUpper(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

var testAllLowerCases = []struct {
	name     string
	input    string
	expected bool
}{
	{name: "ASCII lower-case letters in is.AllLower should return true", input: "abc", expected: true},
	{name: "Unicode lower-case letters in is.AllLower should return true", input: "éö", expected: true},
	{name: "Lower-case letters with uncased characters in is.AllLower should return true", input: "abc 123! 世界", expected: true},
	{name: "ASCII upper-case letters in is.AllLower should return false", input: "ABC", expected: false},
	{name: "Mixed-case letters in is.AllLower should return false", input: "abC", expected: false},
	{name: "Mixed Unicode case letters in is.AllLower should return false", input: "éÖ", expected: false},
	{name: "Title-case letter in is.AllLower should return false", input: "ǅ", expected: false},
	{name: "Empty string in is.AllLower should return false", input: "", expected: false},
	{name: "Uncased characters only in is.AllLower should return false", input: "123 木頭人", expected: false},
	{name: "Invalid UTF-8 in is.AllLower should return false", input: string([]byte{0xff}), expected: false},
}

func testAllLower(t *testing.T) {
	for _, tc := range testAllLowerCases {
		t.Run(tc.name, func(t *testing.T) {
			result := is.AllLower(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
