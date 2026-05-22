package main

import (
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

func encrypt() {
	fmt.Println("🍳 message scrambled!")
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

	if cmd == "help" {
		printHelp()
		return
	}

	if cmd == "encode" {
		encrypt()
		return
	}

	if cmd == "decode" {
		decrypt()
		return
	}
}
