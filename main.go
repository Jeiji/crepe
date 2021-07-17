package main

import "crepe/crepes"


func Crepe() {
	rp := crepes.NewReactParser()
	rp.Scrape()

	ddwp := crepes.NewDockerDesktopWindowsParser()
	ddwp.Scrape()

	dddmp := crepes.NewDockerDesktopMacParser()
	dddmp.Scrape()
}

func main() {

	Crepe()
	
}
