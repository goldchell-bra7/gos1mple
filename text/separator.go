package text

import "strings"

func Separator(symbol string, width int) string {
    return strings.Repeat(symbol, width)
}
