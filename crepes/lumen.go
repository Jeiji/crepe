package crepes

import (
	"crepe/storage"
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type LumenParser struct {
	scraper *colly.Collector
	config  *LumenParserConfig
}

type LumenParserConfig struct {
	tech string
	URL  string
}

func (p *LumenParser) Scrape() {

	p.scraper.OnHTML("p", func(e *colly.HTMLElement) {

		fmt.Println("HOLLAAAA")

		rawText := e.Text
		justTheVersion := strings.Replace(rawText, "Lumen ", "", 1)

		storage.Set("lumenVersion", justTheVersion)
		fmt.Println(storage.Get("lumenVersion"))
	})

}

func NewLumenParser() *LumenParser {
	return &LumenParser{
		scraper: colly.NewCollector(),
		config: &LumenParserConfig{
			URL:  "https://lumen.laravel.com/docs/5.3/releases",
			tech: "Lumen",
		},
	}
}
