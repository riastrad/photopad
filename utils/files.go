package utils

import (
	"image"
	"image/jpeg"
	"log"
	"os"
	"path/filepath"
)

func ReadTxt(path string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func SaveTxt(path string, txt string) {
	err := os.WriteFile(path, []byte(txt), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadJpegPhoto(path string) image.Image {
	extension := filepath.Ext(path)
	if extension != ".jpg" && extension != ".jpeg" {
		log.Fatalf("Image format %s is not supported. Only JPEG files supported currently.", extension)
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
