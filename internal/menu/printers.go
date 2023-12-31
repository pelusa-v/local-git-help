package menu

import "fmt"

type Printer interface {
	// Run() error
	show() (string, error)
}

type MenuPrinter struct {
}

func (menuPrinter *MenuPrinter) show(option int, limit int) {
	clearScreen(limit + 1)
	fmt.Println("Select local github user to use (↑ or ↓)")
	for i := 1; i < limit; i++ {
		if i == option {
			fmt.Printf("%d) Opción %d   <--------\n", i, i)
		} else {
			fmt.Printf("%d) Opción %d\n", i, i)
		}
	}
}

type ErrorPrinter struct {
}

func (errorPrinter *ErrorPrinter) show(option int, limit int) {
	clearScreen(limit + 1)
	fmt.Println("Select local github user to use (↑ or ↓)")
	for i := 1; i < limit; i++ {
		if i == option {
			fmt.Printf("%d) Opción %d   <--------\n", i, i)
		} else {
			fmt.Printf("%d) Opción %d\n", i, i)
		}
	}
}
