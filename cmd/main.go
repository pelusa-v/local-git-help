package main

import (
	"github.com/pelusa-v/local-git-help.git/internal/menu"
)

func main() {

	errPrinter := menu.ErrorPrinter{}
	menuPrinter := menu.MenuPrinter{}

	menu := menu.Menu{
		MenuPrinter:  &menuPrinter,
		ErrorPrinter: &errPrinter,
	}

	if menu.ValidateConfiguration() {
		menu.LoadDefaultMenuOptions()
		menu.Run()
	}
}
