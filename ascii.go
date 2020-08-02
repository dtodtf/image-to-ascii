/*
 * Dylan Todtfeld
 * Package lets user convert image into ascii art.
 * Usage: ./ascii -i [image name] -r [width] [height]
 */
package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	// "strconv"
	"github.com/nfnt/resize"
)

//TODO: suport image resizing
//FIXME: crashes when you input negative numbers
func main() {
	//TODO: support formats other than png
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	imageName, width, height := commandLineArgs()
	openedFile, err := os.Open("./" + imageName)
	errorCheck(err)
	img, _, err := image.Decode(openedFile)
	errorCheck(err)
	img = resizeImage(img, width, height)
	convertToASCII(img)
}

/*
 * Parameter: Nothing. Reads command line args
 * Purpose: Returns name of file to convert and size the new image should be.
 * Returns: The image file name, the desired width, the desired height.
 */
func commandLineArgs() (string, int64, int64) {
	var imageNamePtr *string
	var widthPtr, heightPtr *int64

	//Default value is "" b/c this is required. Dummy so if the file name is 
	//still "" later in this function, we crash.
	imageNamePtr = flag.String("image", "", "Required: the path of the image to turn into ASCII art.")
	//Default value is 80 so it fits on a standard terminal window. May change.
	widthPtr = flag.Int64("width", 80, "The width of the resulting ASCII art.")
	//Default value is 0 so the image scales automatically.
	heightPtr = flag.Int64("height", 0, "The height of the resulting ASCII art.")
	flag.Parse()
	if *imageNamePtr == "" {
		fmt.Println("Error: must provide an image.")
		os.Exit(1)
	}
	return *imageNamePtr, *widthPtr, *heightPtr
}

/*
 * Parameters: image.Image to resize, width to resize to, height to resize to.
 * Purpose: Resize the image to the specified paramete,rs.
 * Returns: image.Image object.
 */ 
func resizeImage(img image.Image, width int64, height int64) image.Image {
	//User doesn't want to reseize
	if width == -1 && height == -1 {
		return img
	}
	return resize.Resize(uint(width), uint(height), img, resize.Lanczos3) //TODO: which algo should be here?
}

/*
 * Parameter: An image.Image object
 * Purpose: Iterate through pixels and determine brightness.
 *          of each pixel. Based on that, print the appropriate ASCII char.
 * Returns: Nothing. Prints to command line.
 */
func convertToASCII(img image.Image) {
	var luminosity float32
	var red, green, blue, alpha uint32
	var charPosition, red8bit, green8bit, blue8bit int

	//Got ascii characters from this link:
	//https://people.sc.fsu.edu/~jburkardt/data/ascii_art_grayscale/ascii_art_grayscale.html
	asciiChars := "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'. "

	// img, _, err := image.Decode(file)
	// errorCheck(err)

	//Y outer loop X inner loop so we print row by row. Otherwise image is
	//rotated
	width, height := img.Bounds().Max.X, img.Bounds().Max.Y
	for y := img.Bounds().Min.Y; y < height; y++ {
		for x := img.Bounds().Min.X; x < width; x++ {
			//.RGBA() premultiplies rgb values by the alpha.
			//To get just RGB vals out of 255, we bitshift right 8 times.
			red, green, blue, alpha = img.At(x, y).RGBA()
			red8bit, green8bit, blue8bit = int(red>>8), int(green>>8), int(blue>>8)
			//Luminosity Formula: https://stackoverflow.com/a/596241/12148894
			//make transparent pixels appear as transparent by giving them max
			//brightness
			if alpha == 0 { 
				luminosity = 255
			} else {
				luminosity = 0.2126*float32(red8bit) + 0.7152*float32(green8bit) +
					0.0722*float32(blue8bit)
			}
			
			//3.65 is the difference in luminosity needed to get a different character.
			charPosition = int(luminosity / 3.65)
			//We print character by character because that's faster than
			//adding every character to a string and then printing the string
			fmt.Print(string(asciiChars[charPosition]))
		}
		fmt.Print("\n")
	}
}

/*
 * Paramter: An error
 * Purpose: Crash the program if there is an error
 * Returns: Nothing. Just crashes program.
 */
func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
