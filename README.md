# asciify

[![GoDoc](https://godoc.org/github.com/tjhorner/makerbot-rpc?status.svg)](https://pkg.go.dev/github.com/tjhorner/asciify)

Asciify is an ASCII art library written in Go. It has no external dependencies and is fairly customizable.

You can use it to convert images to ASCII art.

## Usage

```shell
go get github.com/tjhorner/asciify
```

## Examples

### Asciify With Default Palette

Pass in an `image.Image` to `Asciify` and it will do the thing.

```go
// Pretend we have `img` defined somewhere
result := asciify.Asciify(img, asciify.DefaultCharacterPalette)

// `result` is of type `ASCIIArt`, which you can use like a `[][]string`
char := result[0][5]

// Or if you want the art as a single string, use the String() convenience method
fmt.Println(result.String())
```

<details>
  <summary>Result</summary>

  ### Input
  <img src="test_fixtures/gopher.png"/>

  ### Output
  ```
        'I]fzJYv|<`       
      '-cQZmmmmO0QY['^:.  
  ^tn?uOdbZ0mm0p#Mk0XvLn' 
  .rxjLm%$$@pOOp8$$$kOz}O< 
  `J-rOZfh$$W00U_k$$&0ZuO> 
  {z0Ou;O$$80Ov~b$$W0mJ}. 
    }mOmm8$$hCYOo@$BqOwC'  
    )mm0kWMbU<`[ZbdZ0mmO,  
    rmmmO00QL/}zL0OZmmmZ!  
    ummmmmm0wdwbZ0mmmmmm+  
    ummmmmmZCbpZLmmmmmmm?  
    xmmmmmmm0qmZZmmmmmmm[  
    tmmmmmmmZQQ0mmmmmmmm}  
    )mmmmmmmmmmmmmmmmmmm}  
    ]mmmmmmmmmmmmmmmmmmm{  
    _mmmmmmmmmmmmmmmmmmm1  
  .]mmmmmmmmmmmmmmmmmmmf^ 
  ?OJmmmmmmmmmmmmmmmmmmmCm1
  {f/mmmmmmmmmmmmmmmmmmmc)-
    -mmmmmmmmmmmmmmmmmmmn. 
    }mmmmmmmmmmmmmmmmmmmc. 
    |mmmmmmmmmmmmmmmmmmmz. 
    fmmmmmmmmmmmmmmmmmmmX. 
    rmmmmmmmmmmmmmmmmmmmz. 
    rmmmmmmmmmmmmmmmmmmmn  
    tmmmmmmmmmmmmmmmmmmm|  
    {mmmmmmmmmmmmmmmmmmm_  
    immmmmmmmmmmmmmmmmmO,  
    `Cmmmmmmmmmmmmmmmmmx.  
    ]mmmmmmmmmmmmmmmmOI   
    .jmmmmmmmmmmmmmwO{    
    ^cC0mmmmmmmmmmZYJC"   
    .uhr;_xJ0ZZOQX/i'_p1   
    }/.   `,;I:^'    ~,   
  ```
</details>

### Asciify With Custom Palette

You can also use custom character palettes if you want. The default palette is suitable for most cases, but you might want to use different characters for different scenarios.

To do so, just make a `CharacterPalette` with the characters you'd like to appear in the output, ordered from darkest to lightest.

The default palette comes from [here](http://mewbies.com/geek_fun_files/ascii/ascii_art_light_scale_and_gray_scale_chart.htm).

```go
// Let's use a custom palette
palette := asciify.CharacterPalette{"A", "B", "C", "1", "2", "3"}

result := asciify.Asciify(img, palette)

fmt.Println(result.String())
```

<details>
  <summary>Result</summary>

  ### Input
  <img src="test_fixtures/gopher.png"/>

  ### Output
  ```
  33333333221CCCC1233333333
  3333332CCBBBBBCCCC1332333
  33112CCBBBCBBCBBABCCCC133
  3111CBAAAABCCBAAAABCC1C23
  3C21CB1BAAACCC2BAAACBCC23
  31CCCC2CAAACCC2BAAACBC133
  331BCBBAAABCCCBAAABCBC333
  331BBCBAABC231BBBCCBBC333
  331BBBCCCCC11CCCCBBBBB233
  33CBBBBBBCBBBBBCBBBBBB233
  33CBBBBBBBCBBBCBBBBBBB233
  331BBBBBBBCBBBBBBBBBBB133
  331BBBBBBBBCCCBBBBBBBB133
  331BBBBBBBBBBBBBBBBBBB133
  331BBBBBBBBBBBBBBBBBBB133
  332BBBBBBBBBBBBBBBBBBB133
  332BBBBBBBBBBBBBBBBBBB133
  2CCBBBBBBBBBBBBBBBBBBBCB1
  111BBBBBBBBBBBBBBBBBBBC12
  332BBBBBBBBBBBBBBBBBBBC33
  331BBBBBBBBBBBBBBBBBBBC33
  331BBBBBBBBBBBBBBBBBBBC33
  331BBBBBBBBBBBBBBBBBBBC33
  331BBBBBBBBBBBBBBBBBBBC33
  331BBBBBBBBBBBBBBBBBBBC33
  331BBBBBBBBBBBBBBBBBBB133
  331BBBBBBBBBBBBBBBBBBB233
  332BBBBBBBBBBBBBBBBBBC333
  333CBBBBBBBBBBBBBBBBB1333
  3332BBBBBBBBBBBBBBBBC2333
  33331BBBBBBBBBBBBBBC13333
  3333CCCBBBBBBBBBBBCCC3333
  333CB1221CCCBCCC1232B1333
  3331133333322233333323333
  ```
</details>

### Imagify

If you have some ASCII art and know the palette that was used to create it, you can also turn it back into an image.

This is pretty much just for fun and shouldn't ever be used in any serious way, as the resulting image will be extremely lossy (depending on your palette) and in grayscale.

```go
// Turn an image into ASCII art
result := asciify.Asciify(img, asciify.DefaultCharacterPalette)

// Then re-imagify it
imagified, _ := asciify.Imagify(result, asciify.DefaultCharacterPalette)
```

<details>
  <summary>Result</summary>

  ### Input
  ```
        'I]fzJYv|<`       
      '-cQZmmmmO0QY['^:.  
  ^tn?uOdbZ0mm0p#Mk0XvLn' 
  .rxjLm%$$@pOOp8$$$kOz}O< 
  `J-rOZfh$$W00U_k$$&0ZuO> 
  {z0Ou;O$$80Ov~b$$W0mJ}. 
    }mOmm8$$hCYOo@$BqOwC'  
    )mm0kWMbU<`[ZbdZ0mmO,  
    rmmmO00QL/}zL0OZmmmZ!  
    ummmmmm0wdwbZ0mmmmmm+  
    ummmmmmZCbpZLmmmmmmm?  
    xmmmmmmm0qmZZmmmmmmm[  
    tmmmmmmmZQQ0mmmmmmmm}  
    )mmmmmmmmmmmmmmmmmmm}  
    ]mmmmmmmmmmmmmmmmmmm{  
    _mmmmmmmmmmmmmmmmmmm1  
  .]mmmmmmmmmmmmmmmmmmmf^ 
  ?OJmmmmmmmmmmmmmmmmmmmCm1
  {f/mmmmmmmmmmmmmmmmmmmc)-
    -mmmmmmmmmmmmmmmmmmmn. 
    }mmmmmmmmmmmmmmmmmmmc. 
    |mmmmmmmmmmmmmmmmmmmz. 
    fmmmmmmmmmmmmmmmmmmmX. 
    rmmmmmmmmmmmmmmmmmmmz. 
    rmmmmmmmmmmmmmmmmmmmn  
    tmmmmmmmmmmmmmmmmmmm|  
    {mmmmmmmmmmmmmmmmmmm_  
    immmmmmmmmmmmmmmmmmO,  
    `Cmmmmmmmmmmmmmmmmmx.  
    ]mmmmmmmmmmmmmmmmOI   
    .jmmmmmmmmmmmmmwO{    
    ^cC0mmmmmmmmmmZYJC"   
    .uhr;_xJ0ZZOQX/i'_p1   
    }/.   `,;I:^'    ~,   
  ```

  ### Output
  <img src="test_fixtures/demon_gopher.png"/>
</details>

## FAQ

### Can I resize the output image?

No, as that is outside the scope of this project. Since asciify takes in an `image.Image`, it's compatible with the rest of the Go ecosystem, so you could use a library like this to do the resizing before you asciify the image: https://github.com/nfnt/resize

### Shouldn't you be using runes instead of strings?

Probably. My excuse is that someone somewhere might want to use multiple characters to represent a pixel.

## License

```
Copyright 2020 TJ Horner

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
```