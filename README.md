# image2braille - convert images to braille

Super simple command line program that converts an image to braille and outputs it to the terminal or a file.

Just started learning Go so i was trying to find something interesting to do. The code is pretty straight forward and it was fun to code.

##### Table of contents

- [Usage](#usage)
- [How it works](#how-it-works)
- [Resources](#resources)

---

## Usage

Assuming you have Go installed (https://golang.org/doc/install) go ahead and compile the code:

```bash
go build
```

Run it with --help to see all the available options:

```bash
./image2braille --help

Usage of ./image2braille:
  -image string
    	path to image file (*required)
  -threshold int
    	the color (0-255) threshold to convert a pixel to a dot (default 128)
```

Just fyi you do need to play around with the threshold to figure out what works and what doesn't for some images. I found that although most recommend 128 which i did set as a default higher values such as 150-199 worked better.

---

## How it works

#### The Braile Unicode codepoints (0x2800-0x28FF)

Using the extended set there is a character for a every possible combination 8 dots in a 2x4 grid.

#### Braile Patterns:

```ascii
Braille Pattern        Hexadecimal value of each dot
    #######                     #########
    # 1 4 #                     # 1   8 #
    # 2 5 #                     # 2  10 #
    # 3 6 #                     # 4  20 #
    # 7 8 #                     # 40 80 #
    #######                     #########
```

Originally Braille characters had 6 dots. The two lowest dots were added later that's why the numbering is off for dot 7 and dot 8.

#### Enconding to a bit string is simple enough:

```ascii
#######
# * * #
# * * #
# - - #
# * - #
#######
1 = *, 4 = *, 2 = *, 5 = *, 3 = -, 6 = -, 7 = *, 8 = -

Braille Pattern:        **-**-*-
Dot Numbers available:  12345678
Dots Raised:            11011010
=================================
Binary Value:			01011011 	(by reversing the order of the dots raised)
Decimal (base 10): 		91
Hex (base 16):			0x5b
Braille Character:		0x5b + 0x2800 = 0x285B = '⡛'
```

We can compute the mapping by adding together the hexadecimal values for the dots raised.
Either computing directly in hexadecimals or in binary the result is then added to `0x2800`
the offset for the Braille Patterns Unicode block.

In the code i used a pixel map which is just a two dimensional int array (`[4][2]int`) to store the mapping for the hexadecimal values.

Going through the image we grab each pixel and find a grayscale value for it using some of Go's built-in functions. Depending on the threshold we then decide whether a dot is raised or not.

That's basically it :)

Anyway hope you enjoy using it.

---

## Resources

- Braille Patterns - https://en.wikipedia.org/wiki/Braille_Patterns
- Braille Unicode Pixelation - https://blog.jverkamp.com/2014/05/30/braille-unicode-pixelation/
