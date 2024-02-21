package cryptopals

import (
	"fmt"
)

/*
Convert a literal string to it's binary form
# An exemple of usage is

fmt.Println(cryptopals.StringToBinary("foo"))

output: 011001100110111101101111
*/
func StringToBinary(text string) (binaryString string) {

	// For each char we convert to bin and append it
	for _, rune := range text {
		binaryString = fmt.Sprintf("%s%.8b", binaryString, rune)
	}
	return
}

/*
Evaluate how many bits are different from two strings
An exemple of usage is

fmt.Println(cryptopals.HammingDistance("this is a test", wokka wokka!!!))

output: 37
*/
func HammingDistance(word1, word2 string) (hammingDistance int) {
	if len(word1) != len(word2) {
		panic("Strings with different length")
	}

	// Evaluate equality for each bit pair and score for each
	// different ones
	for i := range word1 {
		if word1[i] != word2[i] {
			hammingDistance++
		}
	}
	return
}

/*
Uses the HammingDistance for guessing the most probable KEYSYZE
With the minimun size of at least 160 characters
var text = "[...]"
fmt.Println(cryptopals.KeySizeGuesser(yourText)

output: 5
*/
func KeySizeGuesser(text string) (bestGuess int) {
	var (
		word1, word2, word3, word4 string
		smallDistance, normalized  float64
		currDist1, currDist2       float64
		currDist3, currDist4       float64
		currDist5, currDist6       float64
	)
	smallDistance = 2000

	// Find the minimum normalized distance from all keysizes
	// by taking the average value from a few interations of
	// haming distance
	for i := 2; i <= 40; i++ {
		word1 = text[0:i]
		word2 = text[i : 2*i]
		word3 = text[2*i : 3*i]
		word4 = text[3*i : 4*i]

		currDist1 = float64(HammingDistance(word1, word2)) / float64(i)
		currDist2 = float64(HammingDistance(word2, word3)) / float64(i)
		currDist3 = float64(HammingDistance(word3, word4)) / float64(i)
		currDist4 = float64(HammingDistance(word1, word3)) / float64(i)
		currDist5 = float64(HammingDistance(word1, word4)) / float64(i)
		currDist6 = float64(HammingDistance(word2, word4)) / float64(i)

		normalized = (currDist1 + currDist2 + currDist3 + currDist4 + currDist5 + currDist6) / 6

		if normalized < smallDistance {
			smallDistance = normalized
			bestGuess = i
		}
	}
	return
}

/*
Transpose a slice in N other blocks such as the first block
receives all the i % N positions from our block

# An exemple of usage is

block := make([]byte, 3)
block = append(block, 1, 2, 3)
fmt.Println(cryptopals.NTransposer(block, 2))

output: [[0 0 2] [0 1 3]]
*/
func NTransposer(block []byte, amount int) [][]byte {
	transpostedBlock := make([][]byte, amount)

	// We make each slice contain a new one
	for i := range transpostedBlock {
		transpostedBlock[i] = make([]byte, 0, 8)
	}

	//We start reindexing and reordering all blocks
	for i, rune := range block {
		index := i % amount
		transpostedBlock[index] = append(transpostedBlock[index], rune)
	}
	return transpostedBlock
}
