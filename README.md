# Go Color

I love the way [fatih/color](https://github.com/fatih/color) work and the
power of [BurntSushi/termcolor](https://github.com/BurntSushi/termcolor).

![Kitty terminal](https://raw.githubusercontent.com/Nguyen-Hoang-Nam/readme-image/main/go-color/go-color.jpg)

## Installation

```bash
go get github.com/Nguyen-Hoang-Nam/go-color
```

## Usage

First you need to create new instance with `New` function then you can use
this instance to run function like fmt package. (Currently this library
only support Println function)

```go
warning := gocolor.New(gocolor.AnsiEscape(gocolor.FgRed))
warning.Println("Some warning")
```

One more thing, `New` function is actually variadic function so you can
add as many parameters as you like.

```go
warning := gocolor.New(gocolor.AnsiEscape(gocolor.FgRed), gocolor.AnsiEscape(gocolor.FgYellow))
warning.Println("Some warning")
```

### Add

Sometime, you may want to add or change color then this function is for you

```go
warning := gocolor.New(gocolor.AnsiEscape(gocolor.FgRed)).Add(gocolor.AnsiEscape(gocolor.BgYellow))
warning.Println("Some warning")
```

### RGB

RGB color is the reason why I create this package

```go
warning := gocolor.New(gocolor.FgRGB(255, 0, 0))
warning.Println("Some warning")
```

In case, your terminal is not support RGB color then I will try to fallback to
256 colors or ansi color base on your flag.

I also support HSV, too. (Thanks to [go-colorful](https://github.com/lucasb-eyer/go-colorful))

```go
warning := gocolor.New(gocolor.FgHSV(120, 1, 0.5))
warning.Println("Some warning")
```

### Flags

This topic is quite complicated, there are 5 flags you can set:

- TrueColor (Only work with true color terminal, otherwise weird character
  will print in screen)
- Color256 (Only work with terminal support 256 colors)
- Ansi (Work in most terminal event Window but look really ugly)
- NoColor (Really, why use this package with this flag :v)
- Auto (Default options, try to guest with flag should be use by checking
  $COLORTERM, $TERM, $NO_COLOR)

```go
truecolor := gocolor.New(gocolor.FgRBG(255, 0, 255)).setFlag(gocolor.TrueColor)
```

There is one caution, most terminal emulator supports users to customize Ansi
color (16 colors). This library simply can not get these customizable colors
so I choose default 16 colors of XTerm as fallback colors.

### Color256

Yes, this function allows you to choose between 256 colors.

```go
color256 := gocolor.New(gocolor.FgColor256(50))
```

### Ansi

This is quite madness, so you may want to read more in [wiki](https://en.wikipedia.org/wiki/ANSI_escape_code#Colors)

There are 8 colors, you can combine with 4 prefix (Fg-, Bg-, FgBright-, BgBright-)

- Black
- Red
- Green
- Yellow
- Blue
- Magenta
- Cyan
- White

```go
ansiColor := gocolor.New(gocolor.AnsiEscape(gocolor.FgRed))
```

I create some constants so that you don't need to remember their codes.

- Reset
- Bold
- Dim
- Italic
- Underline
- Reverse

```go
ansiColor := gocolor.New(gocolor.AnsiEscape(gocolor.FgRed), gocolor.AnsiEscape(gocolor.Underline))
```

### Distance

This is also a tough topic, when I try to fallback RGB aka true color to support
256 color or Ansi color terminal. I might need some way to find the color
which is the nearest to original color.

- DistanceRgb (default)
- DistanceLab
- DistanceLuv
- DistanceCIE94
- DistanceCIEDE2000

```go
distanceLab := gocolor.New(gocolor.FgRBG(255, 0, 255)).SetDistance(gocolor.DistanceLab)
```

Each distances may give you a completely difference fallback color which is
unsurprising

### Complementary

Sometime you may want to make background or foreground to the complementary
color to make a high contrast message.

```go
complementary := gocolor.New(gocolor.FgRBG(255, 0, 255)).AddComplementary()
```

## TODO

- [ ] Pipeline
- [x] Add image to README
- [ ] Support change default Ansi color
- [ ] Add github CI
- [ ] Improve test cases

## Contributing

Pull requests are welcome. For major changes,
please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
