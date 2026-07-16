package utils

import "fmt"

func PrintHelp() {
	fmt.Print(`A picture's worth a thousand pads.™

cmds:
  help   : print this help message
  encrypt: convert plaintext to ciphertext based on a provided .jpg file
  decrypt: convert ciphertext to plaintext based on a provided .jpg file

options:
  --input, -i  : path to a .txt file
  --output, -o : path to write the .txt file
  --key, -k    : path to a valid .jpg file
`)

}
