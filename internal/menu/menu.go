package menu

import (
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
	"github.com/pelusa-v/local-git-help.git/internal/config"
)

type Menu struct {
	MenuPrinter  Printer
	ErrorPrinter Printer
}

func (menu *Menu) Run() {

	config := config.LoadLocalGitData()
	fmt.Println(config)

	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	option := 1
	limit := 7
	menu.MenuPrinter.show()
	// showMenu(option, limit)

	for {
		keyValue, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		if key == keyboard.KeyArrowDown && option < limit-1 {
			option++
		} else if key == keyboard.KeyArrowUp && option > 1 {
			option--
		} else if key == keyboard.KeyCtrlC || key == keyboard.KeyCtrlD || keyValue == 'Q' || keyValue == 'q' {
			break
		}
		menu.MenuPrinter.show()
		// showMenu(option, limit)
	}
}
