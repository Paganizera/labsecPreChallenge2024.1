package cryptopals

import (
	"encoding/hex"
	"strings"
)

/*
We score a string accoring to how it contains chars
following the englosh letter frequency

# An example of usage is

fmt.Println(cryptopals.EvaluateWord("Hello World!"))

output: 76.81
*/
func EvaluateWord(word *string) (totalsum float64) {

	//Built according to <http://www.norvig.com/mayzner.html>
	frequencyTable := map[rune]float64{
		'e': 12.49, 't': 9.28, 'a': 8.04, 'o': 7.64, 'i': 7.57,
		'n': 7.23, 's': 6.51, 'r': 6.28, 'h': 5.05, 'l': 4.07,
		'd': 3.82, 'c': 3.34, 'u': 2.73, 'm': 2.51, 'f': 2.40,
		'p': 2.14, 'g': 1.87, 'w': 1.68, 'y': 1.66, 'b': 1.48,
		'v': 1.05, 'k': 0.54, 'x': 0.23, 'j': 0.16, 'q': 0.12,
		'z': 0.09, ' ': 20}

	// For each char we sum its vallue from the table
	// If the char isn't there, we sum 0 (default value)
	for _, rune := range strings.ToLower(*word) {
		totalsum += frequencyTable[rune]
	}
	return
}

// We xor cipher two words with same length
func FixedWordXor(word1, word2 string) []byte {
	if len(word1) != len(word2) {
		panic("Different word sizes")
	}
	// For each char, we XOR'em and save the result
	wordarray := make([]byte, len(word1))
	for i := 0; i < len(word1); i++ {
		wordarray[i] = word1[i] ^ word2[i]
	}
	return wordarray
}

// We cipher and decipher XOR's encrypted with a single key
func SingleCharXor(word []byte, key byte) string {
	xordWord := []byte{}
	// All characters are encoded against the same character
	for i := range word {
		xordWord = append(xordWord, word[i]^(key))
	}
	return string(xordWord)
}

// Cracks Single Character XORs analyzing it's probability over
func XORSingleCharCracker(word string) (string, rune, float64) {
	var highscoreWord, currWord string
	var highscoreKey rune
	var highscoreValue, currScore float64 = 0, 0
	var currKey byte

	// Stores a curent key and a current vallue. For each
	// better score, update the highest score, key and value
	for i := 0; i < 256; i++ {
		currKey = byte(i)
		// XOR with the current key and analyze it's frequency
		currWord = SingleCharXor([]byte(word), currKey)
		currScore = EvaluateWord(&currWord)
		if currScore > highscoreValue {
			highscoreValue = currScore
			highscoreWord = currWord
			highscoreKey = rune(currKey)
		}

	}
	return highscoreWord, highscoreKey, highscoreValue
}

// We use the repeating XOR here, each key is gonna be
// used periodically to the key's lengh when XORing
func RepeatingXOREncript(word, key string) string {

	var result []byte

	for i := range word {
		result = append(result, word[i]^key[i%len(key)])
	}

	return hex.EncodeToString(result)
}
