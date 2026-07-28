package color

func Term(text, colorConst string) string {
	return colorConst + text + "\033[0m"
}
