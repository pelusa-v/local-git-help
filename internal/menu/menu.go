package menu

import (
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
)

func MenuPrinter() {

	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	option := 0
	limit := 7
	ShowMenu(option, limit)

	for {
		key_value, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		if key == keyboard.KeyArrowDown && option < limit {
			option++
		} else if key == keyboard.KeyArrowUp && option > 0 {
			option--
		} else if key == keyboard.KeyCtrlC || key == keyboard.KeyCtrlD || key_value == 'Q' || key_value == 'q' {
			break
		}

		// ShowMenu(option, limit)
		clearScreen(limit + 2)
		ShowMenu(option, limit)
	}
}

func ShowMenu(option int, limit int) {
	fmt.Println("=============(↑ or ↓)==============")
	for i := 0; i < limit; i++ {
		if i == option-1 {
			fmt.Printf("%d) Opción %d   <--------\n", i, i)
		} else {
			fmt.Printf("%d) Opción %d\n", i, i)
		}
	}
	fmt.Println("===================================")
}

// const escape = "\033"
// const clearLine = escape + "[2K" // ANSI escape code to clear the current line

// func moveCursorUp(lines int) string {
// 	return fmt.Sprintf("%s[%dA", escape, lines)
// }

// func clearScreen(lines int) {
// 	cmd := exec.Command("clear")
// 	cmd.Stdout = os.Stdout
// 	cmd.Run()

// 	// fmt.Print("\033[H")

// 	fmt.Print(moveCursorUp(lines))
// 	fmt.Print(clearLine)
// }

const escape = "\033"

// ANSI escape code to clear the entire screen
const clearScreenStr = escape + "[2J"

// ANSI escape code to move the cursor to the beginning of the line
const moveToBeginningOfLine = escape + "[0G"

//
//
//
//
//
//
//

// ANSI escape code to move the cursor up one line
const moveCursorUp1 = escape + "[F"

// ANSI escape code to clear the entire line
const clearLine = escape + "[2K"

// ANSI escape code to move the cursor up by a specified number of lines
func moveCursorUp(lines int) string {
	return fmt.Sprintf("%s[%dA", escape, lines)
}

func clearScreen(lines int) {
	for i := 0; i < lines; i++ {
		fmt.Print(moveCursorUp1 + clearLine)
	}
	// fmt.Print(clearScreenStr)
	// fmt.Print(moveCursorUp(lines))
}
