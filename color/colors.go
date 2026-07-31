package color

// ANSI escape sequences used to apply foreground colors, background colors,
// and reset formatting in terminals that support ANSI color codes.
//
// Naming convention:
//   - Standard colors:      Red, Green, Blue, etc.
//   - Bright colors:        BrightRed, BrightGreen, BrightBlue, etc.
//   - Background colors:    FillRed, FillGreen, FillBlue, etc.
//   - Reset:                Restores the terminal's default formatting.
//
// These constants can be concatenated with strings to produce colored
// console output. Always append Reset after colored text to prevent
// formatting from affecting subsequent output.
const (
	// Standard foreground colors.
	Red   = "\033[31m"
	Green = "\033[32m"
	Blue  = "\033[34m"
	Black = "\033[30m"
	Yellow = "\033[33m"
	Magenta = "\033[35m"
	Cyan = "\033[36m"
	White = "\033[37m"

	// Bright foreground colors.
	BrightRed     = "\033[91m"
	BrightGreen   = "\033[92m"
	BrightBlue    = "\033[94m"
	BrightBlack   = "\033[90m"
	BrightYellow  = "\033[93m"
	BrightMagenta = "\033[95m"
	BrightCyan    = "\033[96m"
	BrightWhite   = "\033[97m"

	// Background colors.
	FillRed     = "\033[41m"
	FillGreen   = "\033[42m"
	FillBlue    = "\033[44m"
	FillBlack   = "\033[40m"
	FillYellow  = "\033[43m"
	FillMagenta = "\033[45m"
	FillCyan    = "\033[46m"
	FillWhite   = "\033[47m"

	// Reset all terminal formatting.
	Reset = "\033[0m"
)
