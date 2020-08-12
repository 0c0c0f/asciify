package main

import (
	"bytes"
	"encoding/base64"
	"image"
	"strings"
	"syscall/js"

	_ "image/jpeg"
	"image/png"
	_ "image/png"

	"github.com/tjhorner/asciify"
)

func main() {
	asciifyJs := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		dec := base64.NewDecoder(base64.StdEncoding, strings.NewReader(args[0].String()))
		img, _, err := image.Decode(dec)
		if err != nil {
			return err.Error()
		}

		result := asciify.Asciify(img, asciify.DefaultCharacterPalette)

		return result.String()
	})

	imagifyJs := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		artTemp := args[0].String()
		rows := strings.Split(artTemp, "\n")

		art := make(asciify.ASCIIArt, len(rows))
		for x, row := range rows {
			art[x] = strings.Split(row, "")
		}

		result, err := asciify.Imagify(art, asciify.DefaultCharacterPalette)
		if err != nil {
			return "error: " + err.Error()
		}

		w := bytes.NewBuffer([]byte{})
		enc := base64.NewEncoder(base64.StdEncoding, w)
		err = png.Encode(enc, result)
		if err != nil {
			return "error: " + err.Error()
		}

		return w.String()
	})

	js.Global().Set("asciify", asciifyJs)
	js.Global().Set("imagify", imagifyJs)

	<-make(chan bool)
}
