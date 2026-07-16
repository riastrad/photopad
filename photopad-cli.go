package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/riastrad/photopad/lib"
	"github.com/riastrad/photopad/utils"
)

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
	// encrypt: options
	encryptCmd.StringVar(&inputPath, "input", "", "path to the .txt input file")
	encryptCmd.StringVar(&inputPath, "i", "", "path to the .txt input file (short)")
	encryptCmd.StringVar(&imageKeyPath, "key", "", "path to the .jpg file to use as a key")
	encryptCmd.StringVar(&imageKeyPath, "k", "", "path to the .jpg file to use as a key (short)")
	encryptCmd.StringVar(&outputPath, "output", "", "path to save the .txt output file")
	encryptCmd.StringVar(&outputPath, "o", "", "path to save the .txt output file (short")

	decryptCmd := flag.NewFlagSet("decrypt", flag.ExitOnError)
	// decrypt: options
	decryptCmd.StringVar(&inputPath, "input", "default", "path to the .txt input file")
	decryptCmd.StringVar(&inputPath, "i", "default", "path to the .txt input file (short)")
	decryptCmd.StringVar(&imageKeyPath, "key", "", "path to the .jpg file to use as a key")
	decryptCmd.StringVar(&imageKeyPath, "k", "", "path to the .jpg file to use as a key (short)")
	decryptCmd.StringVar(&outputPath, "output", "", "path to save the .txt output file")
	decryptCmd.StringVar(&outputPath, "o", "", "path to save the .txt output file (short")

	switch os.Args[1] {
	case "help":
		utils.PrintHelp()
		return
	case "encrypt":
		encryptCmd.Parse(os.Args[2:])
		lib.Encrypt(inputPath, imageKeyPath, outputPath)
		return
	case "decrypt":
		decryptCmd.Parse(os.Args[2:])
		lib.Decrypt(inputPath, imageKeyPath, outputPath)
		return
	}
}
