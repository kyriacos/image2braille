## Developer Reference

This document explains how `image2braille` works internally and documents the public-facing functions in `main.go` for contributors.

### High-level flow

1. Parse CLI flags (`-image`, `-threshold`).
2. Open and decode the image using Go's standard `image` package.
3. Iterate the image in 2×4 pixel blocks.
4. For each block, compute a Unicode Braille character using grayscale thresholding and a precomputed `pixelMap`.
5. Emit a line of Braille for each vertical step of 4 pixels.

### Data structures

- `pixelMap [4][2]int`: Maps each of the 8 dot positions in a 2×4 block to its hexadecimal value per the Unicode Braille specification. Indexing is `[row][col]`.

### Functions

#### `main()`

Parses flags, validates that `-image` is provided, and invokes `readImage()`. Exits with status code `1` if input is missing.

#### `readImage()`

Opens the file at the path provided by `-image`, decodes it into an `image.Image`, iterates over the bounds in steps of 2 (x) and 4 (y), computes a Braille character for each block via `getChar`, appends to a `strings.Builder`, and prints the final string.

Errors on open/decode are fatal and logged via `log.Fatalln`.

#### `getGrayPixel(x, y int, img image.Image) uint8`

Converts `img.At(x, y)` to grayscale using `color.GrayModel.Convert` and returns the 8-bit intensity. Values are in the range `[0, 255]` where 0 is black and 255 is white.

#### `getChar(imgx, imgy int, img image.Image) string`

Computes the Unicode Braille character corresponding to the 2×4 pixel block whose top-left corner is `(imgx, imgy)`.

For each pixel in the block, it compares the grayscale value to the `-threshold`. If the value is greater than or equal to the threshold, the corresponding dot value from `pixelMap` is added to an accumulator. Finally, the accumulator is offset by `0x2800` to obtain the Unicode code point and converted to a string.

### Performance notes

- Time complexity is proportional to the number of pixels sampled (roughly `width/2 * height/4`).
- For very large images, consider preprocessing (downscaling) or streaming approaches.

### Extensibility ideas

- Add an `-invert` flag to invert threshold logic for dark-on-light vs light-on-dark images.
- Support resizing or aspect-ratio correction before sampling.
- Add colorized output by mapping raised dots to terminal colors.

