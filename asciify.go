package asciify

import (
	"errors"
	"image"
	"image/color"
	"math"
)

// CharacterPalette is a type which represents a set of characters to use in art,
// ordered by the "darkest" character to the "lightest" character.
type CharacterPalette []string

// pick selects a value from the character palette based on the
// given brightness, whose range is assumed to be [0, 65535]
func (p CharacterPalette) pick(brightness uint32) string {
	// We want the inverse of the alpha value since the palette goes
	// from dark to light
	pct := 1.0 - (float64(brightness) / 65535.0)

	max := len(p) - 1
	idx := int(math.Round(float64(max) * pct))

	return p[idx]
}

func (p CharacterPalette) color(character string) (color.Color, error) {
	// TODO: memoize this
	idx := -1
	for i, char := range p {
		if char == character {
			idx = i
			break
		}
	}

	if idx == -1 {
		return nil, errors.New("character not found in palette")
	}

	col := uint8(math.Round((float64(idx+1) / float64(len(p))) * 255))

	return color.RGBA{
		col,
		col,
		col,
		255,
	}, nil
}

// DefaultCharacterPalette is a default set of ASCII characters to be used and is usually fine
// for most ASCII art.
//
// Source: http://mewbies.com/geek_fun_files/ascii/ascii_art_light_scale_and_gray_scale_chart.htm
var DefaultCharacterPalette = CharacterPalette{
	"$", "@", "B", "%", "8", "&", "W", "M", "#", "*", "o", "a", "h", "k", "b", "d", "p", "q", "w",
	"m", "Z", "O", "0", "Q", "L", "C", "J", "U", "Y", "X", "z", "c", "v", "u", "n", "x", "r", "j",
	"f", "t", "/", "|", "(", ")", "1", "{", "}", "[", "]", "?", "-", "_", "+", "~", "<", ">", "i",
	"!", "l", "I", ";", ":", ",", "\"", "^", "`", "'", ".", " ",
}

// rgbaBrightness calculates the luminance of a color.Color
// value. The luminance range is [0, 65535).
func rgbaBrightness(col color.Color) uint32 {
	r, g, b, a := col.RGBA()

	// https://stackoverflow.com/a/596243/2449940
	rawLum := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
	lum := rawLum * (float64(a) / 65535.0)

	return uint32(math.Round(lum))
}

// ASCIIArt represents a string matrix of generated ASCII art
// which is basically just a [][]string but has some useful methods
// for things like turning it into a single string separated by line
// breaks.
type ASCIIArt [][]string

// String turns the matrix into a single string,
// with each row separated by a line break.
func (a ASCIIArt) String() string {
	result := ""

	for _, x := range a {
		for _, y := range x {
			result += y
		}
		result += "\n"
	}

	return result
}

// Asciify takes an image + a character palette and produces a string matrix
// that represents the corresponding ASCII art. If you would like to resize
// the image before ASCIIfying, do that beforehand. Maybe with something
// like this: https://github.com/nfnt/resize
func Asciify(image image.Image, palette CharacterPalette) ASCIIArt {
	bounds := image.Bounds()
	art := make(ASCIIArt, bounds.Max.Y)

	// Iterate over all the pixels
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		art[y] = make([]string, bounds.Max.X)
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			art[y][x] = palette.pick(rgbaBrightness(image.At(x, y)))
		}
	}

	return art
}

// Imagify takes an existing ASCIIArt and turns it into an image.
// Note that this is basically just for fun, is extremely lossy (well,
// that depends on the palette you use), and shouldn't be used in
// production.
func Imagify(art ASCIIArt, palette CharacterPalette) (image.Image, error) {
	bounds := image.Rect(0, 0, len(art[0]), len(art))
	img := image.NewRGBA(bounds)

	for y, col := range art {
		for x, row := range col {
			col, err := palette.color(row)
			if err != nil {
				return nil, err
			}

			img.Set(x, y, col)
		}
	}

	return img, nil
}
