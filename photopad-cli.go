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

func decode() 	{
	fmt.Println("DEcode!")
	// 1. func args should be the text loaded into memory & the pad
	// 2. iterate through the txt & pad array, and for every latin character (shift the letter rightward)
	// 3. save the encoded document next to the original
}

func encode() {
	fmt.Println("ENcode!")
	// 1. func args should be the text loaded into memory & the pad
	// 2. iterate through the txt & pad array, and for every latin character (shift the letter rightward)
	// 3. save the encoded document next to the original
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("No arguments, hence nothing done.")
		return
	}

	cmd := os.Args[1]


	if cmd == "encode" {
		encode()
		return
	}

	if cmd == "decode" {
		decode()
		return
	}
}
