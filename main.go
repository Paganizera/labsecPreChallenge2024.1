package main

import (
	"challenge/cryptopals"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func Challenge1() {
	fmt.Println("Challenge 1")
	hexstring := cryptopals.ReadFile("./data/challenge1/hexstring.txt")
	binstring := string(hexstring[0])
	binstring = cryptopals.HexToBinary(binstring)
	base64result := cryptopals.BinaryToBase64(binstring)
	fmt.Println(base64result)
}

func Challenge2() {
	fmt.Println("\nChallenge 2")

	xorword1 := string(cryptopals.ReadFile("./data/challenge2/xorword1.txt")[0])
	xorword2 := string(cryptopals.ReadFile("./data/challenge2/xorword2.txt")[0])
	word1, _ := hex.DecodeString(xorword1)
	word2, _ := hex.DecodeString(xorword2)
	result := cryptopals.FixedWordXor(string(word1), string(word2))
	fmt.Printf("%x\n", string(result))
}

func Challenge3() {
	fmt.Println("\nChallenge 3")

	encripted := cryptopals.ReadFile("./data/challenge3/encripted.txt")[0]
	word, _ := hex.DecodeString(encripted)

	var result string
	var key rune
	result, key, _ = cryptopals.XORSingleCharCracker(string(word))

	fmt.Println("Word:", result)
	fmt.Println("Key:", string(key))
}

func Challenge4() {
	fmt.Println("\nChallenge 4")

	encripted := cryptopals.ReadFile("./data/challenge4/4.txt")

	var (
		score, currScore     float64 = 0, 0
		currText, actualText string
	)

	for i := 0; i < len(encripted); i++ {
		text, _ := hex.DecodeString(encripted[i])
		currText, _, currScore = cryptopals.XORSingleCharCracker(string(text))
		if currScore > score {
			score = currScore
			actualText = currText
		}
	}
	fmt.Printf("Text: %v", actualText)
}

func Challenge5() {
	fmt.Println("\nChallenge 5")

	var text string
	input := cryptopals.ReadFile("./data/challenge5/toencript.txt")

	for i := range input {
		text += input[i] + "\n"
	}
	encripted := cryptopals.RepeatingXOREncript(text[:len(text)-1], "ICE")
	fmt.Println(encripted)
}

func Challenge6() {
	fmt.Println("\nChallenge 6")
	var text string
	input := cryptopals.ReadFile("./data/challenge6/6.txt")
	for i := range input {
		text += input[i]
	}
	bytes, _ := base64.StdEncoding.DecodeString(text)
	key := cryptopals.KeySizeGuesser(string(bytes))

	blocks := cryptopals.NTransposer(bytes, key)

	var actualkey string
	var auxKey rune
	for i := 0; i < len(blocks); i++ {
		_, auxKey, _ = cryptopals.XORSingleCharCracker(string(blocks[i]))
		actualkey += string(auxKey)
	}
	result := cryptopals.RepeatingXOREncript(string(bytes), actualkey)
	aa, _ := hex.DecodeString(result)
	result = string(aa)
	fmt.Print(result)
}

func main() {
	Challenge1()
	Challenge2()
	Challenge3()
	Challenge4()
	Challenge5()
	Challenge6()
}
