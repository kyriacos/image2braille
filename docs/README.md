## image2braille Documentation

Welcome to the documentation for `image2braille`, a tiny Go CLI that converts images into Unicode Braille art for terminal output or files.

- See CLI usage and options: [CLI Reference](./cli.md)
- Developer internals and function reference: [Developer Reference](./developer-reference.md)
- End-to-end examples and tips: [Examples](./examples.md)

### Installation

- Ensure Go is installed (go1.21+ recommended). See the official docs: `https://golang.org/doc/install`.
- Build the binary:

```bash
go build -o image2braille
```

### Quickstart

Convert an image to Braille and print to your terminal:

```bash
./image2braille -image path/to/photo.jpg
```

Tweak the threshold (0-255) for better contrast (higher values often work better for bright images):

```bash
./image2braille -image path/to/photo.jpg -threshold 180
```

Save output to a file:

```bash
./image2braille -image path/to/photo.jpg > output.txt
```

View with a pager that preserves Unicode:

```bash
./image2braille -image path/to/photo.jpg | less -R
```

### What it does (at a glance)

The tool samples pixels in 2x4 blocks, computes grayscale values, compares against a configurable threshold, and emits the corresponding Unicode Braille character for each block. See [Developer Reference](./developer-reference.md) for details.

