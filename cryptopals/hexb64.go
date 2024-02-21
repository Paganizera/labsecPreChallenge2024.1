package cryptopals

import (
	"math"
	"strings"
)

func HexToBinary(hexString string) string {

	binaryString := ""
	hexBinaryMap := map[rune]string{
		'0': "0000", '1': "0001", '2': "0010", '3': "0011",
		'4': "0100", '5': "0101", '6': "0110", '7': "0111",
		'8': "1000", '9': "1001", 'a': "1010", 'b': "1011",
		'c': "1100", 'd': "1101", 'e': "1110", 'f': "1111"}

	for _, rune := range hexString {
		if hexBinaryMap[rune] != "" {
			binaryString += hexBinaryMap[rune]
		}
	}

	return binaryString
}

func BinaryToBase64(binaryString string) string {

	symbolsBase64 := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	base64String := ""
	var index int

	if math.Mod(float64(len(binaryString)), 6) != 0 {
		binaryString = PadString("0", 6, binaryString, true)
	}

	for i := 0; i < len(binaryString); i += 6 {
		index = BinaryToDecimal(binaryString[i : i+6])
		base64String += string([]rune(symbolsBase64)[index])
	}
	base64String = PadString("=", 4, base64String, true)
	return base64String
}

func PadString(char string, amount int, text string, right bool) string {
	var remainder, amountNeeded int

	remainder = len(text) % amount
	amountNeeded = amount - remainder
	if remainder == 0{
		return text
	}

	if right {
		return text + strings.Repeat(char, amountNeeded)
	}

	return strings.Repeat(char, amountNeeded) + text
}

func BinaryToDecimal(text string) int {
	var sum int = 0
	var k int
	for pow, rune := range text {
		if rune == '1' {
			k = 5 - pow
			sum += int(math.Pow(2.0, float64(k)))
		}
	}
	return sum
}
