package internal

import "fmt"

type Printer interface {
	Print(text string)
}

type SimplePrinter struct{}

func (p *SimplePrinter) Print(text string) {
	fmt.Print(text)
}

func PrintAPIDetails(target API, printer Printer) {
	result := target.Method + " " + target.Path
	result += fmt.Sprintf("\nOperationID: %s\n", target.OperationID)
	result += target.Description
	result += target.PrintParameters()
	printer.Print(result)
}
