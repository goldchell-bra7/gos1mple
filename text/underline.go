package text

func Underline(prompt string) string {
	return "\033[4m" + prompt + "\033[0m"
}
