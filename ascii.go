package main

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"log"
	"os"
)

func main() {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	imageName := getInput()
	imagePixels, err := os.Open("./" + imageName)
	if err != nil {
		log.Fatal(err)
	}
	convertToASCII(imagePixels)
}

func getInput() string {
	if len(os.Args) > 2 {
		fmt.Println("Error: Only one file allowed!")
		os.Exit(1)
	}
	fileName := os.Args[1]
	fmt.Println("You want a picture of " + fileName + ", right?")
	return fileName
}

func convertToASCII(file io.Reader) {
	var luminosity float32
	var red, green, blue uint32
	var charPosition, red8bit, green8bit, blue8bit int
	asciiChars := "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'. " //https://people.sc.fsu.edu/~jburkardt/data/ascii_art_grayscale/ascii_art_grayscale.html

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	// Y outer loop X inner loop so we print row by row. Otherwise image is
	// rotated
	width, height := img.Bounds().Max.X, img.Bounds().Max.Y
	for y := img.Bounds().Min.Y; y < height; y++ {
		for x := img.Bounds().Min.X; x < width; x++ {
			red, green, blue, _ = img.At(x, y).RGBA()
			red8bit, green8bit, blue8bit = int(red>>8), int(green>>8), int(blue>>8) //TODO: figure out how to bitshift automatically
			// fmt.Println("red: " + strconv.Itoa(red8bit) + " green: " + strconv.Itoa(green8bit) + " blue: " + strconv.Itoa(blue8bit))
			luminosity = 0.2126*float32(red8bit) + 0.7152*float32(green8bit) + 0.0722*float32(blue8bit) //https://stackoverflow.com/a/596241/12148894
			// fmt.Println(luminosity)
			charPosition = int(luminosity / 3.65) //3.65 is the difference in luminosity needed to get a different character.
			// fmt.Println(charPosition)
			fmt.Print(string(asciiChars[charPosition])) //We print character by character because that's faster than adding every character to a string and then printing the string
		}
		fmt.Print("\n")
	}
}
