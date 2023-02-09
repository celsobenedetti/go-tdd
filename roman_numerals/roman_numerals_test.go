package roman_numerals

import (
	"fmt"
	"testing"
	"testing/quick"
)

func TestConvertToRoman(t *testing.T) {
	for _, test := range testCases {
		t.Run(fmt.Sprintf("%d gets converted to %q", test.Arabic, test.Roman), func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			assertNumberConversion(t, "ConvertToRoman", got, test.Roman)
		})
	}
}

func TestConvertToArabic(t *testing.T) {
	for _, test := range testCases[:4] {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			assertNumberConversion(t, "ConvertToArabic", fmt.Sprint(got), fmt.Sprint(test.Arabic))
		})
	}
}

func TestPropertiesOfConvertion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}

		t.Log("testing", arabic)

		roman := ConvertToRoman(int(arabic))
		fromRoman := ConvertToArabic(roman)
		return fromRoman == int(arabic)
	}

	if err := quick.Check(assertion, &quick.Config{MaxCount: 1000}); err != nil {
		t.Error("failed checks", err)
	}
}

func assertNumberConversion(t testing.TB, desc, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("%q: got %q, want %q", desc, got, want)
	}
}

type Numeral struct {
	Arabic int
	Roman  string
}

var testCases = []Numeral{
	{Arabic: 1, Roman: "I"},
	{Arabic: 2, Roman: "II"},
	{Arabic: 3, Roman: "III"},
	{Arabic: 4, Roman: "IV"},
	{Arabic: 5, Roman: "V"},
	{Arabic: 6, Roman: "VI"},
	{Arabic: 7, Roman: "VII"},
	{Arabic: 8, Roman: "VIII"},
	{Arabic: 9, Roman: "IX"},
	{Arabic: 10, Roman: "X"},
	{Arabic: 14, Roman: "XIV"},
	{Arabic: 18, Roman: "XVIII"},
	{Arabic: 20, Roman: "XX"},
	{Arabic: 39, Roman: "XXXIX"},
	{Arabic: 40, Roman: "XL"},
	{Arabic: 47, Roman: "XLVII"},
	{Arabic: 49, Roman: "XLIX"},
	{Arabic: 50, Roman: "L"},
	{Arabic: 100, Roman: "C"},
	{Arabic: 90, Roman: "XC"},
	{Arabic: 400, Roman: "CD"},
	{Arabic: 500, Roman: "D"},
	{Arabic: 900, Roman: "CM"},
	{Arabic: 1000, Roman: "M"},
	{Arabic: 1984, Roman: "MCMLXXXIV"},
	{Arabic: 3999, Roman: "MMMCMXCIX"},
	{Arabic: 2014, Roman: "MMXIV"},
	{Arabic: 1006, Roman: "MVI"},
	{Arabic: 798, Roman: "DCCXCVIII"},
}
