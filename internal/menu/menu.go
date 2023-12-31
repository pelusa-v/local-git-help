package menu

import (
	"log"

	"github.com/eiannone/keyboard"
	"github.com/pelusa-v/local-git-help.git/internal/config"
)

type Menu struct {
	MenuPrinter  *MenuPrinter
	ErrorPrinter *ErrorPrinter
}

func (menu *Menu) ValidateConfiguration() bool {
	if !config.FileExists(config.GetLocalConfigPath()) {
		menu.ErrorPrinter.Content = "Local configuration file doesn't exist"
		menu.ErrorPrinter.show()
		return false
	}

	return true
}

func (menu *Menu) LoadDefaultMenuOptions() {
	menu.MenuPrinter.Options = make([]string, 0)
	menu.MenuPrinter.SelectedOption = 0
	config := config.LoadLocalGitData()
	configValues := (*config)
	for i := 0; i < len(configValues); i++ {
		menu.MenuPrinter.Options = append(menu.MenuPrinter.Options, configValues[i].GetSummary())
	}
}

func (menu *Menu) Run() {

	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	menu.MenuPrinter.show()

	for {
		keyValue, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		}

		if key == keyboard.KeyArrowDown && menu.MenuPrinter.SelectedOption < len(menu.MenuPrinter.Options)-1 {
			menu.MenuPrinter.SelectedOption++
		} else if key == keyboard.KeyArrowUp && menu.MenuPrinter.SelectedOption > 0 {
			menu.MenuPrinter.SelectedOption--
		} else if key == keyboard.KeyCtrlC || key == keyboard.KeyCtrlD || keyValue == 'Q' || keyValue == 'q' {
			break
		}

		menu.MenuPrinter.show()
	}
}
