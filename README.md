# bee

A simple Go test library with colorful output and easily locatable error locations.

## Methods

```golang
func Test(t *testing.T) {
    bee := bee.New(t)
    bee.Nil(errors.New("whoopsie"))  // whoopsie != <nil>
    bee.NotNil(nil)                  // <nil> == <nil>
    bee.True(false)                  // false != true
    bee.False(true)                  // true != false
    bee.Equal(1, 2)                  // 1 != 2
}
```

## Configure

```golang
func Test(t *testing.T) {
    bee := bee.New(
        t,
        bee.ActualColor(191, 29, 245),    // set <actual> color to rgb(191, 29, 245)
        bee.ExpectedColor(50, 168, 127),  // set <expected> color to rgb(50, 168, 127)
        bee.WhatColor(224, 154, 22),      // set <what> color to rgb(224, 154, 22)
        bee.ColumnWidth(60),              // set column width to 60
    )
}
```

### Disable color

1. set the `bee.NoColor()` option in `bee.New()`
2. pass the `-nocolor` flag to `go test`

## Output

### Basic types

- Format: `<actual> != <expected>`

```golang
func Test(t *testing.T) {
    bee := bee.New(t)
    bee.Equal(1, 2)
    // 1 != 2
}
```

### Complex types

- Format: `<actual> != <expected> (<what>)`

```golang
type Person struct {
    Name string
}

func Test(t *testing.T) {
    bee := bee.New(t)
    actual := []Person{{Name: "Obi-Wan Kenobi"}}
    expected := []Person{{Name: "Jar Jar Binks"}}
    bee.Equal(actual, expected)
    // Obi-Wan Kenobi != Jar Jar Binks ([0].Name)
}
```

### Expand

The length of `<actual>` and `<expected>` is limited to the column width.

If the output would be longer than twice the column width, the values of `<actual>` and `<expected>` are expanded side-by-side in an additional log message.

```golang
type Book struct {
    Chapters []Chapter
}

type Chapter struct {
    Text string
}

func Test(t *testing.T) {
    bee := bee.New(t, bee.ColumnWidth(30))
    actual := Book{
        Chapters: []Chapter{
            {
                Text: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam vehicula volutpat justo, scelerisque egestas nisl volutpat.",
            },
        },
    }
    expected := Book{
        Chapters: []Chapter{
            {
                Text: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Curabitur scelerisque, sapien eget mattis mattis, eros mi nisi.",
            },
        },
    }
    bee.Equal(actual, expected)
    // Lorem ipsum dolor sit amet,... != Lorem ipsum dolor sit amet,... (.Chapters[0].Text)
    //  
    //    Lorem ipsum dolor sit amet,    Lorem ipsum dolor sit amet,  
    //    consectetur adipiscing elit.   consectetur adipiscing elit. 
    //    Nam vehicula volutpat justo,   Curabitur scelerisque, sapien
    //    scelerisque egestas nisl       eget mattis mattis, eros mi  
    //    volutpat.                      nisi.
}
```
