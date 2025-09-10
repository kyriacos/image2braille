## CLI Reference

### Name

`image2braille` — Convert images to Unicode Braille art

### Synopsis

```bash
./image2braille -image <path> [-threshold <0-255>]
```

### Options

- `-image string` (required): Path to the input image file. Supported formats include PNG, JPEG, and GIF (via Go's standard library decoders).
- `-threshold int` (default `128`): Grayscale threshold (0-255). Pixels with grayscale >= threshold are considered “raised” dots. Higher values generally produce sparser, higher-contrast output.

### Examples

- Convert an image with default threshold:

```bash
./image2braille -image images/cat.jpg
```

- Increase threshold for brighter images:

```bash
./image2braille -image images/cat.jpg -threshold 190
```

- Save to a file for later viewing:

```bash
./image2braille -image images/cat.jpg > cat.txt
```

- View with a pager that preserves ANSI/Unicode:

```bash
./image2braille -image images/cat.jpg | less -R
```

### Exit Codes

- `0`: Success
- `1`: Usage error (e.g., `-image` not provided) or I/O/decoding failures

### Help

Display available flags and defaults:

```bash
./image2braille --help
```

