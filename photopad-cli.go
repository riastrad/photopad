package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/riastrad/photopad/lib"
)

func printHelp() {
	fmt.Print(`A picture's worth a thousand pads.™

cmds:
  help   : print this help message
  encrypt: convert plaintext to ciphertext based on a provided .jpg file
  decrypt: convert ciphertext to plaintext based on a provided .jpg file

options:
  --input, -i : path to a plaintext file
  --key, -k   : path to a valid .jpg file
`)

}

var (
	inputPath    string
	imageKeyPath string
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

	decryptCmd := flag.NewFlagSet("decrypt", flag.ExitOnError)
	decryptCmd.StringVar(&inputPath, "input", "default", "path to the .txt input file")
	decryptCmd.StringVar(&inputPath, "i", "default", "path to the .txt input file (short)")
	decryptCmd.StringVar(&imageKeyPath, "key", "", "path to the .jpg file to use as a key")
	decryptCmd.StringVar(&imageKeyPath, "k", "", "path to the .jpg file to use as a key (short)")

	switch os.Args[1] {
	case "help":
		printHelp()
		return
	case "encrypt":
		encryptCmd.Parse(os.Args[2:])
		lib.Encrypt(inputPath, imageKeyPath)
		return
	case "decrypt":
		decryptCmd.Parse(os.Args[2:])
		lib.Decrypt(inputPath, imageKeyPath)
		return
	}
}
