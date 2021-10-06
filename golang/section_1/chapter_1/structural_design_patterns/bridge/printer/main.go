package main

func main() {
	macComputer := &mac{}
	windowsComputer := &windows{}
	hpPrinter := &hp{}
	epsonPrinter := &epson{}

	macComputer.setPrinter(hpPrinter)
	macComputer.print()

	macComputer.setPrinter(epsonPrinter)
	macComputer.print()

	windowsComputer.setPrinter(hpPrinter)
	windowsComputer.print()

	windowsComputer.setPrinter(epsonPrinter)
	windowsComputer.print()
}
