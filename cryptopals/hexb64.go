package cryptopals

import (
	"math"
	"strings"
)

// We convert a hexadecimal string to it's binary form
// using a conversion table
func HexToBinary(hexString *string) (binaryString string) {

	hexBinaryMap := map[rune]string{
		'0': "0000", '1': "0001", '2': "0010", '3': "0011",
		'4': "0100", '5': "0101", '6': "0110", '7': "0111",
		'8': "1000", '9': "1001", 'a': "1010", 'b': "1011",
		'c': "1100", 'd': "1101", 'e': "1110", 'f': "1111"}

	// Iterates each character for conversion
	for _, rune := range *hexString {
		if hexBinaryMap[rune] != "" {
			binaryString += hexBinaryMap[rune]
		}
	}
	return
}

// Convert a binary string to base64 by making them decimal
// and checking their position into symbol list
func BinaryToBase64(binaryString *string) (base64String string) {

	symbolsBase64 := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	var index int

	// We averiguate if there's enough character for conversion
	// If not, we pad 0's to the MSBs until enough
	if math.Mod(float64(len(*binaryString)), 6) != 0 {
		*binaryString = PadString("0", 6, binaryString, true)
	}

	// The conversion is realized using our table
	for i := 0; i < len(*binaryString); i += 6 {
		index = BinaryToDecimal((*binaryString)[i : i+6])
		base64String += string([]rune(symbolsBase64)[index])
	}
	// We check if the string has a 4 multiple length
	// if not, we pad it right with "="
	base64String = PadString("=", 4, &base64String, true)
	return
}

// For padding right or left, the amount needed with a specific character
func PadString(char string, amount int, text *string, right bool) string {

	var remainder, amountNeeded int
	remainder = len(*text) % amount

	// If theres no need to pad, return
	if remainder == 0 {
		return *text
	}

	// If we need to pad, calculate and pad it
	amountNeeded = amount - remainder

	if right {
		return *text + strings.Repeat(char, amountNeeded)
	}

	return strings.Repeat(char, amountNeeded) + *text
}

// Convert bit strings to decimal notation
func BinaryToDecimal(text string) (sum int) {
	var max int = len(text) - 1
	var aux int
	for pow, rune := range text {
		if rune == '1' {
			aux = max - pow
			sum += int(math.Pow(2.0, float64(aux)))
		}
	}
	return
}
