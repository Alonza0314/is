package is_test

import (
	"testing"

	"github.com/Alonza0314/is"
)

func TestLetter(t *testing.T) {
	testUpper(t)
	testLower(t)
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
