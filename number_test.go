package is_test

import (
	"testing"

	"github.com/Alonza0314/is"
)

func TestNumber(t *testing.T) {
	testPowerOf(t)
	testPowerOfTwo(t)
}

var testPowerOfCases = []struct {
	name     string
	input    int
	base     int
	expected bool
}{
	{
		name:     "One input in is.PowerOf should return true",
		input:    1,
		base:     2,
		expected: true,
	},
	{
		name:     "Power of base in is.PowerOf should return true",
		input:    8,
		base:     2,
		expected: true,
	},
	{
		name:     "Different power of base in is.PowerOf should return true",
		input:    81,
		base:     3,
		expected: true,
	},
	{
		name:     "Non-power of base in is.PowerOf should return false",
		input:    12,
		base:     2,
		expected: false,
	},
	{
		name:     "Zero input in is.PowerOf should return false",
		input:    0,
		base:     2,
		expected: false,
	},
	{
		name:     "Negative input in is.PowerOf should return false",
		input:    -8,
		base:     2,
		expected: false,
	},
	{
		name:     "Zero base in is.PowerOf should return false",
		input:    8,
		base:     0,
		expected: false,
	},
	{
		name:     "One base in is.PowerOf should return false",
		input:    8,
		base:     1,
		expected: false,
	},
	{
		name:     "Negative base in is.PowerOf should return false",
		input:    8,
		base:     -2,
		expected: false,
	},
}

func testPowerOf(t *testing.T) {
	for _, tc := range testPowerOfCases {
		t.Run(tc.name, func(t *testing.T) {
			result := is.PowerOf(tc.input, tc.base)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

var testPowerOfTwoCases = []struct {
	name     string
	input    int
	expected bool
}{
	{
		name:     "One input in is.PowerOfTwo should return true",
		input:    1,
		expected: true,
	},
	{
		name:     "Power of two in is.PowerOfTwo should return true",
		input:    16,
		expected: true,
	},
	{
		name:     "Non-power of two in is.PowerOfTwo should return false",
		input:    12,
		expected: false,
	},
	{
		name:     "Zero input in is.PowerOfTwo should return false",
		input:    0,
		expected: false,
	},
	{
		name:     "Negative input in is.PowerOfTwo should return false",
		input:    -16,
		expected: false,
	},
}

func testPowerOfTwo(t *testing.T) {
	for _, tc := range testPowerOfTwoCases {
		t.Run(tc.name, func(t *testing.T) {
			result := is.PowerOfTwo(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
