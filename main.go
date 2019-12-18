package main

import (
	"flag"
	"fmt"
	"image"
	"log"
	"os"
	"strings"

	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

/*
   Braile Unicode codepoints (0x2800-0x28FF)

   Using the extended set there is a character for a every possible combination
   8 dots in a 2x4 grid.
   Originally Braille characters had 6 dots so the two lowest dots were added later
   that's why the numbering is off.

   Braile Patterns:
   --------------------------------------------------
   #######
   # 1 4 #
   # 2 5 #
   # 3 6 #
   # 7 8 #
   #######

   Braille Pattern hexadecimal value of the braille dots:
   ------------------------------------------------------
   #########
   # 1   8 #
   # 2  10 #
   # 4  20 #
   # 40 80 #
   #########


   Enconding to a bitstring however is simple enough:
   --------------------------------------------------
   #######
   # * * #
   # * * #
   # - - #
   # * - #
   #######
   1 = *, 4 = *, 2 = *, 5 = *, 3 = -, 6 = -, 7 = *, 8 = -

   Braille Pattern:		    **-**-*-
   Dot Numbers available:	12345678
   Dots Raised:			    11011010
   =================================
   Binary Value:			01011011 	(by reversing the order of the dots raised)
   Decimal (base 10): 		91
   Hex (base 16):			0x5b
   Braille Character:		0x5b + 0x2800 = 0x285B = '⡛'

   We can compute the mapping by adding together the hexadecimal values for the dots raised.
   Either computing directly in hexadecimals or in binary the result is then added to 0x2800
   the offset for the Braille Patterns Unicode block.
*/

/*
   Pixelmap stores the hex values for dots
   #########
   # 1   8 #
   # 2  10 #
   # 4  20 #
   # 40 80 #
   #########
*/
var pixelMap = [4][2]int{
	{0x1, 0x8},
	{0x2, 0x10},
	{0x4, 0x20},
	{0x40, 0x80},
}

var (
	imageFile = flag.String("image", "", "path to image file (*required)")
	threshold = flag.Int("threshold", 128, "the color (0-255) threshold to convert a pixel to a dot")
)

func main() {
	flag.Parse()

	if *imageFile == "" {
		flag.Usage()
		os.Exit(1)
	}

	readImage()
}

func readImage() {
	reader, err := os.Open(*imageFile)
	if err != nil {
		log.Fatalln("failed to load image:", err)
	}
	defer reader.Close()

	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatalln("failed to decode image:", err)
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	str := new(strings.Builder)
	for y := 0; y < height; y += 4 {
		for x := 0; x < width; x += 2 {
			str.WriteString(getChar(x, y, img))
		}
		str.WriteString("\n")
	}
	fmt.Println(str.String())
}

func getGrayPixel(x, y int, img image.Image) uint8 {
	p := img.At(x, y)

	// grayscale conversion
	// Y = 0.299 * R +  0.587 * G + 0.114 * B
	// https://stackoverflow.com/a/42518487

	r, _, _, _ := color.GrayModel.Convert(p).RGBA()
	return uint8(r) // gray >> 8
}

func getChar(imgx, imgy int, img image.Image) string {
	val := 0
	for y := 0; y < 4; y++ {
		for x := 0; x < 2; x++ {
			if getGrayPixel(imgx+x, imgy+y, img) < uint8(*threshold) {
				val += 0
			} else {
				val += pixelMap[y%4][x%2]
			}
		}
	}
	return string(val + 0x2800)
}
