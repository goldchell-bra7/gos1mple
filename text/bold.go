package text

func Bold(prompt string) string {
	return "\033[1m" + prompt + "\033[0m"
}
