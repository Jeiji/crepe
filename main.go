package main

import (
	"crepe/crepes"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

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

	done := make(chan os.Signal, 1)

	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	c := cron.New()
	c.AddFunc("0 0 */1 * * *", func() {
		CrepeItUp()
	})
	c.Start()

	fmt.Println("\n\nAwaiting Signal...")
	<-done
	fmt.Println("\n\nExiting...")

}
