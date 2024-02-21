package cryptopals

import (
	"encoding/hex"
	"strings"
)

func EvaluateWord(word string) float64 {

	//Built according to <http://www.norvig.com/mayzner.html>
	frequencyTable := map[rune]float64{
		'e': 12.49, 't': 9.28, 'a': 8.04, 'o': 7.64, 'i': 7.57,
		'n': 7.23, 's': 6.51, 'r': 6.28, 'h': 5.05, 'l': 4.07,
		'd': 3.82, 'c': 3.34, 'u': 2.73, 'm': 2.51, 'f': 2.40,
		'p': 2.14, 'g': 1.87, 'w': 1.68, 'y': 1.66, 'b': 1.48,
		'v': 1.05, 'k': 0.54, 'x': 0.23, 'j': 0.16, 'q': 0.12,
		'z': 0.09, ' ': 20, '!': 0.001, '"': 0.001, '\'': 0.001,
		'(': 0.001, ')': 0.001, ',': 0.001, '.': 0.001,
		':': 0.001, ';': 0.001, '?': 0.001}

	var totalsum float64

	for _, rune := range strings.ToLower(word) {
		totalsum += frequencyTable[rune]
	}

	return totalsum
}

func FixedWordXor(word1, word2 string) []byte {
	wordarray := make([]byte, len(word1))
	for i := 0; i < len(word1); i++ {
		wordarray[i] = word1[i] ^ word2[i]
	}
	return wordarray
}

func SingleBitXor(word []byte, key byte) string {
	xordWord := []byte{}
	for i := range word {
		xordWord = append(xordWord, word[i]^(key))
	}
	return string(xordWord)
}

func XORSingleBitCracker(word string) (string, rune, float64) {
	var highscoreWord, currWord string
	var highscoreKey rune
	var highscoreValue, currScore float64 = 0, 0
	var currKey byte
	for i := 0; i < 256; i++ {
		currKey = byte(i)
		currWord = SingleBitXor([]byte(word), currKey)
		currScore = EvaluateWord(currWord)
		if currScore > highscoreValue {
			highscoreValue = currScore
			highscoreWord = currWord
			highscoreKey = rune(currKey)
		}

	}
	return highscoreWord, highscoreKey, highscoreValue
}

func RepeatingXOREncript(word, key string) string {

	var result []byte

	for i := range word {
		result = append(result, word[i]^key[i%len(key)])
	}

	return hex.EncodeToString(result)
}
