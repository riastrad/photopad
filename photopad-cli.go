package main

import (
	"flag"
	"fmt"
	"os"
)

// TODO: all of it
func readTxt(path string) {}

func readPhoto(path string) {}

func squishPixelsIntoPad() {}

func shiftTxtWithPad() {}

func unshiftTxtWithPad() {}

func decrypt() {
	fmt.Println("🥚 message unscrambled!")
	// 1. func args should be the text loaded into memory & the pad
	// 2. iterate through the txt & pad array, and for every latin character (shift the letter rightward)
	// 3. save the encoded document next to the original
}

func encrypt(inputPath string, keyPath string, outputPath string) {
	fmt.Printf(`🍳 would have scrambled!
input: %v
key: %v
output: %v
`, inputPath, keyPath, outputPath)
	// 1. func args should be the text loaded into memory & the pad
	// 2. iterate through the txt & pad array, and for every latin character (shift the letter rightward)
	// 3. save the encoded document next to the original
}

func printHelp() {
	fmt.Print(`A picture's worth a thousand pads.™

cmds:
  help   : print this help message
  encrypt: convert plaintext to ciphertext based on a provided .jpg file
  decrypt: convert ciphertext to plaintext based on a provided .jpg file

options:
  --input, -i : path to a plaintext file
  --output, -o: path to write the results of the operation [default=./{cmd}.output.txt]
  --key, -k   : path to a valid .jpg file
`)

}

func main() {
	if len(os.Args) <= 1 {
		printHelp()
		return
	}
	cmd := os.Args[1]
	if cmd != "help" && cmd != "encrypt" && cmd != "decrypt" {
		fmt.Println("Invalid command. Must be one of: help, encrypt, decrypt")
		return
	}

	inputPath := flag.String("input", "", "path to an existing .txt file")
	// var keyPath = flag.String("key", "", "path to the .jpg file to use as a key")
	// var outputPath = flag.String("output", fmt.Sprintf(`./%s.output.txt`, cmd), "path to write the resulting .txt file")

	flag.Parse()

	switch cmd {
	case "help":
		printHelp()
		return
	case "encrypt":
		// encrypt(*inputPath, *keyPath, *outputPath)
		fmt.Println("input value:", *inputPath)
		return
	case "decrypt":
		decrypt()
		return
	}
}
