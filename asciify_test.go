package asciify_test

import (
	"fmt"
	"image/png"
	"io/ioutil"
	"os"
	"testing"

	"github.com/tjhorner/asciify"
)

func oops(t *testing.T, err error) {
	if err != nil {
		if t != nil {
			t.Errorf("fatal error while testing: %s", err)
		} else {
			panic(err)
		}
	}
}

func ExampleAsciify() {
	img, err := os.Open("./test_fixtures/gopher.png")
	oops(nil, err)

	gopher, err := png.Decode(img)
	oops(nil, err)

	result := asciify.Asciify(gopher, asciify.DefaultCharacterPalette)

	fmt.Println(result.String())
}

func TestAsciify_LargeImage(t *testing.T) {
	img, err := os.Open("./test_fixtures/sam.png")
	oops(t, err)

	expectedASCII, err := ioutil.ReadFile("./test_fixtures/sam.txt")
	oops(t, err)

	sam, err := png.Decode(img)
	oops(t, err)

	result := asciify.Asciify(sam, asciify.DefaultCharacterPalette)

	// Correct width
	if len(result) != sam.Bounds().Max.Y {
		t.Errorf("len(result) = %d; want %d", len(result), sam.Bounds().Max.Y)
	}

	// Correct height
	if len(result[0]) != sam.Bounds().Max.X {
		t.Errorf("len(result[0]) = %d; want %d", len(result[0]), sam.Bounds().Max.X)
	}

	// Actual ASCII matches what we want
	if result.String() != string(expectedASCII) {
		t.Error(("result.String() did not match expected value"))
	}
}

func TestAsciify_DefaultValues(t *testing.T) {
	img, err := os.Open("./test_fixtures/gopher.png")
	oops(t, err)

	expectedASCII, err := ioutil.ReadFile("./test_fixtures/default_gopher.txt")
	oops(t, err)

	gopher, err := png.Decode(img)
	oops(t, err)

	result := asciify.Asciify(gopher, asciify.DefaultCharacterPalette)

	// Correct width
	if len(result) != gopher.Bounds().Max.Y {
		t.Errorf("len(result) = %d; want %d", len(result), gopher.Bounds().Max.Y)
	}

	// Correct height
	if len(result[0]) != gopher.Bounds().Max.X {
		t.Errorf("len(result[0]) = %d; want %d", len(result[0]), gopher.Bounds().Max.X)
	}

	// Actual ASCII matches what we want
	if result.String() != string(expectedASCII) {
		t.Error(("result.String() did not match expected value"))
	}
}

func TestAsciify_CustomPalette(t *testing.T) {
	img, err := os.Open("./test_fixtures/gopher.png")
	oops(t, err)

	expectedASCII, err := ioutil.ReadFile("./test_fixtures/custom_gopher.txt")
	oops(t, err)

	gopher, err := png.Decode(img)
	oops(t, err)

	customPalette := asciify.CharacterPalette{"A", "B", "C", "1", "2", "3"}

	result := asciify.Asciify(gopher, customPalette)

	// Correct width
	if len(result) != gopher.Bounds().Max.Y {
		t.Errorf("len(result) = %d; want %d", len(result), gopher.Bounds().Max.Y)
	}

	// Correct height
	if len(result[0]) != gopher.Bounds().Max.X {
		t.Errorf("len(result[0]) = %d; want %d", len(result[0]), gopher.Bounds().Max.X)
	}

	// Actual ASCII matches what we want
	if result.String() != string(expectedASCII) {
		t.Error(("result.String() did not match expected value"))
	}
}

func TestImagify(t *testing.T) {
	img, err := os.Open("./test_fixtures/gopher.png")
	oops(t, err)

	gopher, err := png.Decode(img)
	oops(t, err)

	expectedImg, err := os.Open("./test_fixtures/gopher.png")
	oops(t, err)

	demonGopher, err := png.Decode(expectedImg)
	oops(t, err)

	result := asciify.Asciify(gopher, asciify.DefaultCharacterPalette)

	imagified, err := asciify.Imagify(result, asciify.DefaultCharacterPalette)
	oops(t, err)

	// Correct width
	if imagified.Bounds().Max.Y != demonGopher.Bounds().Max.Y {
		t.Errorf("imagified.Bounds().Max.Y = %d; want %d", imagified.Bounds().Max.Y, demonGopher.Bounds().Max.Y)
	}

	// Correct height
	if imagified.Bounds().Max.X != demonGopher.Bounds().Max.X {
		t.Errorf("imagified.Bounds().Max.X = %d; want %d", imagified.Bounds().Max.X, demonGopher.Bounds().Max.X)
	}

	// I don't think the rest is worth testing, but if you do,
	// feel free to finish the test lol.
}
