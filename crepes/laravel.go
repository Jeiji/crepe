package crepes

import (
	"crepe/storage"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
	"github.com/slack-go/slack"
)

type LaravelParser struct {
	scraper *colly.Collector
	config  *LaravelParserConfig
}

type LaravelParserConfig struct {
	tech string
	URL  string
}

func (p *LaravelParser) Scrape() {

	p.scraper.OnHTML("div.release-entry:first-of-type > div > div > div > div > div > a", func(e *colly.HTMLElement) {

		redislaravelVersion, err := storage.Get("laravelVersion")
		if err == nil {
			fmt.Println("This is in the Laravel bank: ", redislaravelVersion)
		}

		rawText := e.Text
		justTheVersion := strings.Replace(rawText, "v", "", 1)

		if redislaravelVersion == justTheVersion {
			fmt.Println("[ DONE ] Laravel already up to date")
		} else {

			storage.Set("laravelVersion", justTheVersion)
			fmt.Println("What went in: ", justTheVersion)

			slack.PostWebhook(os.Getenv("SLACK_HOOK_URL"), &slack.WebhookMessage{
				Username: "Crépe",
				Text:     fmt.Sprintf("This is new %s info. Title: %v.", p.config.tech, *e),
			})
		}

	})

	p.scraper.Visit(p.config.URL)

}

func NewLaravelParser() *LaravelParser {
	return &LaravelParser{
		scraper: colly.NewCollector(),
		config: &LaravelParserConfig{
			URL:  "https://github.com/laravel/laravel/releases",
			tech: "Laravel",
		},
	}
}
