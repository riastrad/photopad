package lib

import (
	"fmt"
	"strings"

	"github.com/riastrad/photopad/utils"
)

func DecodeLetter(letter rune, shift uint8) rune {
	if letter > utils.BOUNDARY {
		return letter
	}

	oldUnicodePoint := (int(letter) + (int(utils.BOUNDARY) - (int(shift) % int(utils.BOUNDARY)))) % int(utils.BOUNDARY)
	return rune(oldUnicodePoint)
}

func UnshiftTxtWithPad(txt string, pads []uint8) string {
	outputString := make([]rune, 0, len(txt))

	for ix, val := range txt {
		shiftIndex := ix % len(pads)
		pad := pads[shiftIndex]

		decoded := DecodeLetter(val, pad)
		outputString = append(outputString, decoded)
	}

	return string(outputString)
}

func Decrypt(inputPath string, imageKeyPath string, outputPath string) {
	inputText := utils.ReadTxt(inputPath)
	fmt.Printf("\n✔ loaded text from: %s\n", inputPath)

	pads := utils.GetPadFromImage(imageKeyPath)
	fmt.Printf("✔ generated pad from: %s\n", imageKeyPath)

	output := UnshiftTxtWithPad(inputText, pads)
	fmt.Println("✔ deciphered text with pad")

	if outputPath == "" {
		outputPath = strings.Replace(inputPath, ".txt", ".decrypted.txt", 1)
	}
	utils.SaveTxt(outputPath, output)
	fmt.Printf("\n💾 saved output to: %s\n", outputPath)
}
