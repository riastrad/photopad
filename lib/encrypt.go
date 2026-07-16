package lib

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/riastrad/photopad/utils"
)

func EncodeLetter(letter rune, shift uint8) rune {
	if letter > unicode.MaxASCII {
		return letter
	}

	newUnicodePoint := (int(letter) + int(shift)) % unicode.MaxASCII
	return rune(newUnicodePoint)
}

func ShiftTxtWithPad(txt string, pads []uint8) string {
	outputString := make([]rune, 0, len(txt))

	for ix, val := range txt {
		shiftIndex := ix % len(pads)
		pad := pads[shiftIndex]

		encoded := EncodeLetter(val, pad)
		outputString = append(outputString, encoded)
	}

	return string(outputString)
}

func Encrypt(inputPath string, imageKeyPath string) {
	inputText := utils.ReadTxt(inputPath)
	fmt.Printf("✔ loaded text from: %s\n", inputPath)

	pads := utils.GetPadFromImage(imageKeyPath)
	fmt.Printf("✔ generated pad from: %s\n", imageKeyPath)

	output := ShiftTxtWithPad(inputText, pads)
	fmt.Println("✔ ciphered text with pad")

	outputPath := strings.Replace(inputPath, ".txt", ".encrypted.txt", 1)
	utils.SaveTxt(outputPath, output)
	fmt.Printf("💾 saved output to: %s\n", outputPath)
}
