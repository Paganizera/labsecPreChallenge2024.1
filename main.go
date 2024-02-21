package main

import (
	"challenge/cryptopals"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
)
// Fun function just for printing a cool intro
func init() {
	fmt.Println("###################################################")
	path := "./data/init/labsec.txt"
	ASCIIArt := cryptopals.ReadFile(&path)

	for i := range ASCIIArt {
		fmt.Print("#", ASCIIArt[i], "#\n")
	}
	fmt.Println("###################################################")
}

func Challenge1() {
	// We convert the hex string to binary
	// then the binary to decimal
	// and convert by mapping onto a premade list
	// of base64 characters
	fmt.Println("\nChallenge 1")
	path := "./data/challenge1/hexstring.txt"
	hexstring := cryptopals.ReadFile(&path)
	binstring := string(hexstring[0])
	binstring = cryptopals.HexToBinary(&binstring)
	base64result := cryptopals.BinaryToBase64(&binstring)
	fmt.Println(base64result)
}

func Challenge2() {
	// We read our strings and then convert to hex
	// Finally we XOR them and make them readable
	fmt.Println("\nChallenge 2")
	path1 := "./data/challenge2/xorword1.txt"
	path2 := "./data/challenge2/xorword2.txt"
	xorword1 := string(cryptopals.ReadFile(&path1)[0])
	xorword2 := string(cryptopals.ReadFile(&path2)[0])
	word1, _ := hex.DecodeString(xorword1)
	word2, _ := hex.DecodeString(xorword2)
	result := cryptopals.FixedWordXor(string(word1), string(word2))
	fmt.Printf("%x\n", string(result))
}

func Challenge3() {
	// We try all combinations of single char strings
	// and then return the best scored one
	fmt.Println("\nChallenge 3")
	path := "./data/challenge3/encripted.txt"
	encripted := cryptopals.ReadFile(&path)[0]
	word, _ := hex.DecodeString(encripted)

	var result string
	var key rune
	result, key, _ = cryptopals.XORSingleCharCracker(string(word))

	fmt.Println("Word:", result)
	fmt.Println("Key:", string(key))
}

func Challenge4() {
	// We use the challenge 3's idea and then 
	// score all the lines separetelly for furthermore
	// return the highest score one
	fmt.Println("\nChallenge 4")
	path := "./data/challenge4/4.txt"
	encripted := cryptopals.ReadFile(&path)

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
	// Basic repeating XOR encription, similar as 3 but 
	// rotating our chars
	fmt.Println("\nChallenge 5")
	path := "./data/challenge5/toencript.txt"
	var text string
	input := cryptopals.ReadFile(&path)

	for i := range input {
		text += input[i] + "\n"
	}
	encripted := cryptopals.RepeatingXOREncript(text[:len(text)-1], "ICE")
	fmt.Println(encripted)
}

func Challenge6() {
	// Here we follow all the steps from cryptopals' guide
	// but basically we guess the keysize from the encrypted 
	// text, then we reorder them transposing into KEYLENGH size
	// blocks, then guessing each char separetaly and concatening
	// the final key, the it's a simple repeating XOR
	fmt.Println("\nChallenge 6")
	path := "./data/challenge6/6.txt"
	var text string
	input := cryptopals.ReadFile(&path)
	for i := range input {
		text += input[i]
	}
	bytes, _ := base64.StdEncoding.DecodeString(text)
	aux := string(bytes)
	key := cryptopals.KeySizeGuesser(&aux)

	blocks := cryptopals.NTransposer(&bytes, key)

	var actualkey string
	var auxKey rune
	for i := 0; i < len(blocks); i++ {
		_, auxKey, _ = cryptopals.XORSingleCharCracker(string(blocks[i]))
		actualkey += string(auxKey)
	}
	result := cryptopals.RepeatingXOREncript(string(bytes), actualkey)
	finalString, _ := hex.DecodeString(result)
	result = string(finalString)
	fmt.Println("Text:", result)
	fmt.Println("Key", actualkey)
}

func main() {
	var i string
	fmt.Println("Welcome to cryptopals answers")
	fmt.Println("Insert the level you want to check the answers!")
	for {
		fmt.Println("Challenges 1-6 are available.\n0 for ending operations")
		fmt.Scanln(&i)
		switch i {
		case "1":
			Challenge1()
		case "2":
			Challenge2()
		case "3":
			Challenge3()
		case "4":
			Challenge4()
		case "5":
			Challenge5()
		case "6":
			Challenge6()
		case "0":
			fmt.Println("Thanks for using our program!")
			os.Exit(0)
		default:
			fmt.Println("Invalid input, please try again")
		}
	}
}
