## Examples

### Basic conversion

```bash
./image2braille -image samples/landscape.jpg
```

### Adjust threshold for contrast

```bash
./image2braille -image samples/portrait.png -threshold 192
```

### Save output to a file

```bash
./image2braille -image samples/logo.png > logo.txt
```

### Sample output snippet

```
⣀⣀⣀⣀⣀⣀⣀
⣿⣿⣿⣿⣿⣿⣿
⠛⠛⠛⠛⠛⠛⠛
```

Note: Actual output depends on the image content and chosen threshold.

