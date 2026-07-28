<p align="center">
  <img src="https://raw.githubusercontent.com/goldchell-bra7/gos1mple/main/docs/gopher-gos1mple.png" alt="gos1mple banner">
</p>

<h1 align="center">gos1mple</h1>

<p align="center">
A lightweight Go library that simplifies console applications.
</p>

---

## ✨ Features

- 📥 Easy input functions
- 📝 Functions for text
- 🌈 Colors for text
- and more (*coming soon*)

---

## 📦 Installation

```bash
go get github.com/goldchell-bra7/gos1mple@latest
```

*or*

```bash
go get github.com/goldchell-bra7/gos1mple@v0.3.0
```
---

## 🚀 Usage

```go
package main

import (
    "fmt"

    "github.com/goldchell-bra7/gos1mple/input"
)

func main() {
    age := input.ReadInt("Age: ")
    fmt.Println(age)
}
```

---

# 📚 API

## Input

| Function | Description |
|----------|-------------|
| `ReadInt(prompt string)` | Read integer |
| `ReadIntln(prompt string)` | Read integer with prompt on new line |
| `ReadFloat(prompt string)` | Read float64 |
| `ReadFloatln(prompt string)` | Read float64 with prompt on new line |
| `ReadString(prompt string)` | Read string |
| `ReadStringln(prompt string)` | Read string with prompt on new line |
| `ReadBool(prompt string)` | Read bool |
| `ReadBoolln(prompt string)` | Read bool with prompt on new line |

## Text

| Function | Description |
|----------|-------------|
| `Header(text, symbol string, width int)` | Places your text in a frame |
| `Separator(symbol string, width int)` | Inserts a separator |
| `Bold(prompt string)` | Makes text bold |
| `Cursive(prompt string)` | Makes text cursive |
| `Underline(prompt string)` | Underlines the text |
| `Strike(prompt string)` | Strikes through the text |

## Color

### Functions

| Function | Description |
|----------|-------------|
| `Term(text, colorConst string)` | Prints terminal text in the selected color |

### Constants

- `Red`
- `BrightRed`
- `FillRed`
- `Green`
- `BrightGreen`
- `FillGreen`
- `Blue`
- `BrightBlue`
- `FillBlue`
- `Black`
- `BrightBlack`
- `FillBlack`
- `Yellow`
- `BrightYellow`
- `FillYellow`
- `Magenta`
- `BrightMagenta`
- `FillMagenta`
- `Cyan`
- `BrightCyan`
- `FillCyan`
- `White`
- `BrightWhite`
- `FillWhite`
- `Reset`

---

## 📄 License

MIT
