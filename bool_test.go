package is_test

import (
	"testing"

	"github.com/Alonza0314/is"
)

func TestBool(t *testing.T) {
	testTrue(t)
	testFalse(t)
}

var testBoolTrueCases = []struct {
	name     string
	input    bool
	expected bool
}{
	{
		name:     "True input in is.True should return true",
		input:    true,
		expected: true,
	},
	{
		name:     "False input in is.True should return false",
		input:    false,
		expected: false,
	},
}

func testTrue(t *testing.T) {
	for _, tc := range testBoolTrueCases {
		t.Run(tc.name, func(t *testing.T) {
			result := is.True(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

var testFalseCases = []struct {
	name     string
	input    bool
	expected bool
}{
	{
		name:     "True input in is.False should return false",
		input:    true,
		expected: false,
	},
	{
		name:     "False input in is.False should return true",
		input:    false,
		expected: true,
	},
}

func testFalse(t *testing.T) {
	for _, tc := range testFalseCases {
		t.Run(tc.name, func(t *testing.T) {
			result := is.False(tc.input)
			if result != tc.expected {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
