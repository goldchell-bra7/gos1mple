package text

func Cursive(prompt string) string {
	return "\033[3m" + prompt + "\033[0m"
}
