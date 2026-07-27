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
- 📝 Fuctions for text
- and more (*coming soon*)

---

## 📦 Installation

```bash
go get github.com/goldchell-bra7/gos1mple@latest
```

*or*

```bash
go get github.com/goldchell-bra7/gos1mple@v0.1.0
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
| `ReadInt()` | Read integer |
| `ReadIntln()` | Read integer with prompt on new line |
| `ReadFloat()` | Read float64 |
| `ReadFloatln()` | Read float64 with prompt on new line |
| `ReadString()` | Read string |
| `ReadStringln()` | Read string with prompt on new line |
| `ReadBool()` | Read bool |
| `ReadBoolln()` | Read bool with prompt on new line |

## Text

| Function | Description |
|----------|-------------|
| `Header()` | Places your text in a frame |
| `Separator()` | Inserts a separator |

---

## 📄 License

MIT
