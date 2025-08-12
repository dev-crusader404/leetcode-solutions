package exercise

import "testing"

type labelTestCase struct {
	description string
	input       []string
	expected    string
}

var labelTestCases = []labelTestCase{
	{
		description: "Orange and orange and black",
		input:       []string{"orange", "orange", "black"},
		expected:    "33 ohms",
	},
	{
		description: "Blue and grey and brown",
		input:       []string{"blue", "grey", "brown"},
		expected:    "680 ohms",
	},
	{
		description: "Red and black and red",
		input:       []string{"red", "black", "red"},
		expected:    "2 kiloohms",
	},
	{
		description: "Green and brown and orange",
		input:       []string{"green", "brown", "orange"},
		expected:    "51 kiloohms",
	},
	{
		description: "Yellow and violet and yellow",
		input:       []string{"yellow", "violet", "yellow"},
		expected:    "470 kiloohms",
	},
	{
		description: "Blue and violet and blue",
		input:       []string{"blue", "violet", "blue"},
		expected:    "67 megaohms",
	},
	{
		description: "Minimum possible value",
		input:       []string{"black", "black", "black"},
		expected:    "0 ohms",
	},
	{
		description: "Maximum possible value",
		input:       []string{"white", "white", "white"},
		expected:    "99 gigaohms",
	},
	{
		description: "First two colors make an invalid octal number",
		input:       []string{"black", "grey", "black"},
		expected:    "8 ohms",
	},
	{
		description: "Ignore extra colors",
		input:       []string{"blue", "green", "yellow", "orange"},
		expected:    "650 kiloohms",
	},
}

func TestLabel(t *testing.T) {
	for _, tc := range labelTestCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := Label(tc.input)
			if actual != tc.expected {
				t.Fatalf("Label(%+q): expected %q, actual %q", tc.input, tc.expected, actual)
			}
		})
	}
}
