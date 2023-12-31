package menu

import (
	"fmt"

	"github.com/pelusa-v/local-git-help.git/internal/config"
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
	Options        []config.LocalGitUser
}

func (menuPrinter *MenuPrinter) show() {
	clearScreen(len(menuPrinter.Options) + 1)
	fmt.Println("Select local github user to use (↑ or ↓), 'q' or 'Q' for exit")
	for i := 0; i < len(menuPrinter.Options); i++ {
		if i == menuPrinter.SelectedOption {
			fmt.Printf("%d) %s   <--------\n", i+1, menuPrinter.Options[i].GetSummary())
		} else {
			fmt.Printf("%d) %s\n", i+1, menuPrinter.Options[i].GetSummary())
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
