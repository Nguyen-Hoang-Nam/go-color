# Go Color

I love the way [fatih/color](https://github.com/fatih/color) work and the
power of [BurntSushi/termcolor](https://github.com/BurntSushi/termcolor).

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
warning.Println("Some waring")
```

One more thing, `New` function is actually variadic function so you can
add as many parameters as you like.

### Add

Sometime, you may want to add or change color then this function is for you

```go
warning := gocolor.New(gocolor.AnsiEscape(gocolor.FgRed)).Add(gocolor.BgYellow)
warning.Println("Some waring")
```

### RGB

RGB color is the reason why I create this package

```go
warning := gocolor.New(gocolor.FgRGB(255, 0, 0))
warning.Println("Some waring")
```

In case, your terminal is not support rgb color then I will try to fallback to
256 colors or ansi color base on your flag.

### Flags

This topic is quite complicated, there are 5 flags you can set:

- TrueColor (Only work with true color terminal, otherwise weird character
  will print in screen)
- Color256 (Only work with terminal support 256 colors)
- Ansi (Work in most terminal event Window but look really ugly)
- NoColor (Really, why use with package with this flag :v)
- Auto (Try to guest with flag should be use by checking $COLORTERM, $TERM,
$NO_COLOR)

```go
truecolor := gocolor.New(gocolor.FgRBG(255, 0, 255)).setFlag(gocolor.TrueColor)
```

### Distance

This also a tough topic, when I try to fallback rgb aka true color to support
256 color or Ansi color terminal. I might need some way to find the color
which nearest original color.

## Contributing

Pull requests are welcome. For major changes,
please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
