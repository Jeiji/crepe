package main

import (
	"crepe/crepes"
	"sync"
)

type Crepe interface {
	Scrape()
}

func CrepeItUp() {

	crepes := []Crepe{
		// crepes.NewReactParser(),
		// crepes.NewDockerDesktopWindowsParser(),
		// crepes.NewDockerDesktopMacParser(),
		// crepes.NewPHPParser(),
		crepes.NewLumenParser(),
	}

	var wg sync.WaitGroup
	for _, crepe := range crepes {
		wg.Add(1)
		go func(c Crepe) {
			c.Scrape()
			wg.Done()
		}(crepe)
	}
	wg.Wait()

}

func main() {

	CrepeItUp()

}
