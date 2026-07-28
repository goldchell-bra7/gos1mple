package text

func Strike(prompt string) string {
	return "\033[9m" + prompt + "\033[0m"
}
