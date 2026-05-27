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

func decrypt(inputPath string, imageKeyPath string, outputPath string) {
	fmt.Printf(`🥚 would have unscrambled!
input: %v
key: %v
output: %v
`, inputPath, imageKeyPath, outputPath)
	// 1. func args should be the text loaded into memory & the pad
	// 2. iterate through the txt & pad array, and for every latin character (shift the letter rightward)
	// 3. save the encoded document next to the original
}

func encrypt(inputPath string, imageKeyPath string, outputPath string) {
	fmt.Printf(`🍳 would have scrambled!
input: %v
key: %v
output: %v
`, inputPath, imageKeyPath, outputPath)
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

var (
	inputPath    string
	imageKeyPath string
	outputPath   string
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Expected either 'encrypt' or 'decrypt' command")
		os.Exit(1)
	}

	encryptCmd := flag.NewFlagSet("encrypt", flag.ExitOnError)
	encryptCmd.StringVar(&inputPath, "input", "", "path to the .txt input file")
	encryptCmd.StringVar(&inputPath, "i", "", "path to the .txt input file (short)")
	encryptCmd.StringVar(&imageKeyPath, "key", "", "path to the .jpg file to use as a key")
	encryptCmd.StringVar(&imageKeyPath, "k", "", "path to the .jpg file to use as a key (short)")
	encryptCmd.StringVar(&outputPath, "output", fmt.Sprintf(`./%s.output.txt`, os.Args[1]), "path to write the resulting .txt file")
	encryptCmd.StringVar(&outputPath, "o", fmt.Sprintf(`./%sed.output.txt`, os.Args[1]), "path to write the resulting .txt file (short)")

	decryptCmd := flag.NewFlagSet("decrypt", flag.ExitOnError)
	decryptCmd.StringVar(&inputPath, "input", "default", "path to the .txt input file")
	decryptCmd.StringVar(&inputPath, "i", "default", "path to the .txt input file (short)")
	decryptCmd.StringVar(&imageKeyPath, "key", "", "path to the .jpg file to use as a key")
	decryptCmd.StringVar(&imageKeyPath, "k", "", "path to the .jpg file to use as a key (short)")
	decryptCmd.StringVar(&outputPath, "output", fmt.Sprintf(`./%s.output.txt`, os.Args[1]), "path to write the resulting .txt file")
	decryptCmd.StringVar(&outputPath, "o", fmt.Sprintf(`./%sed.output.txt`, os.Args[1]), "path to write the resulting .txt file (short)")

	switch os.Args[1] {
	case "help":
		printHelp()
		return
	case "encrypt":
		encryptCmd.Parse(os.Args[2:])
		encrypt(inputPath, imageKeyPath, outputPath)
		return
	case "decrypt":
		decryptCmd.Parse(os.Args[2:])
		decrypt(inputPath, imageKeyPath, outputPath)
		return
	}
}
