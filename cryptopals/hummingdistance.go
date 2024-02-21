package cryptopals

import (
	"fmt"
)

func StringToBinary(text string) string {
	var binaryString string
	for _, rune := range text {
		binaryString = fmt.Sprintf("%s%.8b", binaryString, rune)
	}
	return binaryString
}
func HummingDistance(word1, word2 string) int {
	hummingDistance := 0
	if len(word1) != len(word2) {
		panic("Strings with different lengh")
	}

	for i := range word1 {
		if word1[i] != word2[i] {
			hummingDistance++
		}
	}
	return hummingDistance
}

func KeySizeGuesser(text string) int {
	var (
		bestGuess                  int
		word1, word2, word3, word4 string
		smallDistance, normalized  float64
		currDist1, currDist2       float64
		currDist3, currDist4       float64
		currDist5, currDist6       float64
	)
	smallDistance = 2000
	for i := 2; i <= 40; i++ {
		word1 = text[0:i]
		word2 = text[i : 2*i]
		word3 = text[2*i : 3*i]
		word4 = text[3*i : 4*i]

		currDist1 = float64(HummingDistance(word1, word2)) / float64(i)
		currDist2 = float64(HummingDistance(word2, word3)) / float64(i)
		currDist3 = float64(HummingDistance(word3, word4)) / float64(i)
		currDist4 = float64(HummingDistance(word1, word3)) / float64(i)
		currDist5 = float64(HummingDistance(word1, word4)) / float64(i)
		currDist6 = float64(HummingDistance(word2, word4)) / float64(i)

		normalized = (currDist1 + currDist2 + currDist3 + currDist4 + currDist5 + currDist6) / 6

		if normalized < smallDistance {
			smallDistance = normalized
			bestGuess = i
		}
	}
	return bestGuess
}

func NTransposer(block []byte, amount int) [][]byte {
	transpostedBlock := make([][]byte, amount)
	for i := range transpostedBlock {
		transpostedBlock[i] = make([]byte, 0, 8)
	}
	for i, rune := range block {
		index := i % amount
		transpostedBlock[index] = append(transpostedBlock[index], rune)
	}
	return transpostedBlock
}
