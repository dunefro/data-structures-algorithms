package main

func main() {
	cl := &client{}
	mac := &mac{}

	cl.insertLightningConnectorIntoComputer(mac)

	windowMachine := &windows{}
	windowsMachineAdapter := &windowsAdapter{
		windowMachine: windowMachine,
	}
	cl.insertLightningConnectorIntoComputer(windowsMachineAdapter)
}
