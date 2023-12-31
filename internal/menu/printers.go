package menu

import (
	"fmt"
)

type Printer interface {
	// Run() error
	show() error
}

type MenuOption struct {
	Title       string
	Description string
}

type MenuPrinter struct {
	SelectedOption int
	Options        []string
}

func (menuPrinter *MenuPrinter) show() {
	clearScreen(len(menuPrinter.Options) + 1)
	fmt.Println("Select local github user to use (↑ or ↓), 'q' or 'Q' for exit")
	for i := 0; i < len(menuPrinter.Options); i++ {
		if i == menuPrinter.SelectedOption {
			fmt.Printf("%d) %s   <--------\n", i+1, menuPrinter.Options[i])
		} else {
			fmt.Printf("%d) %s\n", i+1, menuPrinter.Options[i])
		}
	}
}

type ErrorPrinter struct {
	Content string
	// LinesToClear int
}

func (errorPrinter *ErrorPrinter) show() {
	// clearScreen(errorPrinter.LinesToClear)
	fmt.Printf("%s", errorPrinter.Content)
}
