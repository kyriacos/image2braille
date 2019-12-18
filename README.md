# Convert Images to Braille

Super simple command line program that converts an image to braille and outputs it to the terminal or a file.

Just started learning Go so i was trying to find something interesting to do. The code is pretty straight forward and it was fun to code.

## Compiling and running the program

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

### Resources used in making this:

- Braille Patterns - https://en.wikipedia.org/wiki/Braille_Patterns
- Braille Unicode Pixelation - https://blog.jverkamp.com/2014/05/30/braille-unicode-pixelation/
