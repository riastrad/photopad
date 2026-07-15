package main

import (
	"flag"
	"fmt"
	"image"
	"image/color/palette"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
	"strings"
	"unicode"
)

func readTxt(path string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func saveTxt(path string, txt string) {
	err := os.WriteFile(path, []byte(txt), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func readJpegPhoto(path string) image.Image {
	if path[len(path)-3:] != "jpg" {
		log.Fatal("Unsupported image format. Only .jpg files supported currently.")
	}

	imageFile, readErr := os.Open(path)
	if readErr != nil {
		log.Fatal(readErr)
	}

	defer imageFile.Close()

	image, decodeErr := jpeg.Decode(imageFile)
	if decodeErr != nil {
		log.Fatal(decodeErr)
	}
	return image
}

func getPixelColorIndices(img image.Image) []uint8 {
	// using Plan9 means index values will always be between 0-255
	paletted := image.NewPaletted(img.Bounds(), palette.Plan9)

	// this could be a bottleneck for large images
	draw.Draw(paletted, paletted.Rect, img, img.Bounds().Min, draw.Src)

	return paletted.Pix
}

func getPadFromImage(path string) []uint8 {
	image := readJpegPhoto(path)
	pads := getPixelColorIndices(image)

	return pads
}

func shiftTxtWithPad(txt string, pads []uint8) string {
	outputString := make([]rune, 0, len(txt))

	for ix, val := range txt {
		shiftIndex := ix % len(pads)
		pad := pads[shiftIndex]

		if val <= unicode.MaxASCII {
			newIndex := (int(val) + int(pad)) % unicode.MaxASCII
			outputString = append(outputString, rune(newIndex))
		} else {
			outputString = append(outputString, val)
		}
	}

	return string(outputString)
}

func encodeLetter(letter rune, shift int) {}
func decodeLetter(letter rune, shift int) {}

func unshiftTxtWithPad(txt string, pads []uint8) string {
	outputString := make([]rune, 0, len(txt))

	for ix, val := range txt {
		shiftIndex := ix % len(pads)
		pad := pads[shiftIndex]

		if val <= unicode.MaxASCII {
			oldIndex := (int(val) + (unicode.MaxASCII - (int(pad) % unicode.MaxASCII))) % unicode.MaxASCII
			outputString = append(outputString, rune(oldIndex))
		} else {
			outputString = append(outputString, val)
		}

	}

	return string(outputString)
}

func decrypt(inputPath string, imageKeyPath string) {
	inputText := readTxt(inputPath)
	fmt.Printf("✔ loaded text from: %s\n", inputPath)

	pads := getPadFromImage(imageKeyPath)
	fmt.Printf("✔ generated pad from: %s\n", imageKeyPath)

	output := unshiftTxtWithPad(inputText, pads)
	fmt.Println("✔ deciphered text with pad")

	outputPath := strings.Replace(inputPath, ".txt", ".decrypted.txt", 1)
	saveTxt(outputPath, output)
	fmt.Printf("💾 saved output to: %s\n", outputPath)
}

func encrypt(inputPath string, imageKeyPath string) {
	inputText := readTxt(inputPath)
	fmt.Printf("✔ loaded text from: %s\n", inputPath)

	pads := getPadFromImage(imageKeyPath)
	fmt.Printf("✔ generated pad from: %s\n", imageKeyPath)

	output := shiftTxtWithPad(inputText, pads)
	fmt.Println("✔ ciphered text with pad")

	outputPath := strings.Replace(inputPath, ".txt", ".encrypted.txt", 1)
	saveTxt(outputPath, output)
	fmt.Printf("💾 saved output to: %s\n", outputPath)
}

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
		encrypt(inputPath, imageKeyPath)
		return
	case "decrypt":
		decryptCmd.Parse(os.Args[2:])
		decrypt(inputPath, imageKeyPath)
		return
	}
}
