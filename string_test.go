package is_test

import (
	"testing"

	"github.com/Alonza0314/is"
)

func TestString(t *testing.T) {
	testAllUpper(t)
	testAllLower(t)
	testPalindrome(t)
	testAlnumPalindrome(t)
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

var testPalindromeCases = []struct {
	name     string
	input    string
	expected bool
}{
	{name: "Odd-length palindrome in is.Palindrome should return true", input: "racecar", expected: true},
	{name: "Even-length palindrome in is.Palindrome should return true", input: "abba", expected: true},
	{name: "Unicode palindrome in is.Palindrome should return true", input: "世界世", expected: true},
	{name: "Palindrome with mirrored spaces in is.Palindrome should return true", input: "a b a", expected: true},
	{name: "Single rune in is.Palindrome should return true", input: "a", expected: true},
	{name: "Empty string in is.Palindrome should return true", input: "", expected: true},
	{name: "Non-palindrome in is.Palindrome should return false", input: "hello", expected: false},
	{name: "Case mismatch in is.Palindrome should return false", input: "Aba", expected: false},
	{name: "Unmirrored punctuation in is.Palindrome should return false", input: "a,b.a", expected: false},
	{name: "Invalid UTF-8 in is.Palindrome should return false", input: string([]byte{0xff}), expected: false},
}

func testPalindrome(t *testing.T) {
	for _, tc := range testPalindromeCases {
		t.Run(tc.name, func(t *testing.T) {
			result := is.Palindrome(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

var testAlnumPalindromeCases = []struct {
	name     string
	input    string
	expected bool
}{
	{name: "Exact palindrome in is.AlnumPalindrome should return true", input: "racecar", expected: true},
	{name: "Mixed-case alphanumeric palindrome in is.AlnumPalindrome should return true", input: "A1b1a", expected: true},
	{name: "Sentence palindrome in is.AlnumPalindrome should return true", input: "A man, a plan, a canal: Panama", expected: true},
	{name: "Unicode alphanumeric palindrome in is.AlnumPalindrome should return true", input: "É世é", expected: true},
	{name: "Non-alphanumeric characters only in is.AlnumPalindrome should return true", input: " ,! ", expected: true},
	{name: "Empty string in is.AlnumPalindrome should return true", input: "", expected: true},
	{name: "Non-palindrome in is.AlnumPalindrome should return false", input: "race a car", expected: false},
	{name: "Alphanumeric mismatch in is.AlnumPalindrome should return false", input: "0P", expected: false},
	{name: "Invalid UTF-8 in is.AlnumPalindrome should return false", input: string([]byte{0xff}), expected: false},
}

func testAlnumPalindrome(t *testing.T) {
	for _, tc := range testAlnumPalindromeCases {
		t.Run(tc.name, func(t *testing.T) {
			result := is.AlnumPalindrome(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
