package menu

import (
	"fmt"
)

func clearScreen(lines int) {
	const escape = "\033"
	// ANSI escape code to move the cursor up one line
	const moveCursorUp = escape + "[F"
	// ANSI escape code to clear the entire line
	const clearLine = escape + "[2K"

	for i := 0; i < lines; i++ {
		fmt.Print(moveCursorUp + clearLine)
	}
}
