package main

import (
	"crepe/crepes"
	"sync"

	"github.com/robfig/cron"
)

type Crepe interface {
	Scrape()
}

func CrepeItUp() {

	crepes := []Crepe{
		crepes.NewReactParser(),
		crepes.NewDockerDesktopWindowsParser(),
		crepes.NewDockerDesktopMacParser(),
		crepes.NewPHPParser(),
		crepes.NewLumenParser(),
		crepes.NewLaravelParser(),
		crepes.NewVueParser(),
		crepes.NewNextjsParser(),
		crepes.NewNuxtjsParser(),
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

	c := cron.New()
	c.AddFunc("@hourly", func() {
		CrepeItUp()
	})

}
