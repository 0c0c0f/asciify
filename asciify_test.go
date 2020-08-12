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
