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
		clearScreen()
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

func clearScreen() {
	fmt.Print("\033[H\033[2J") // ANSI escape codes to clear the screen
}
