package crepes

import (
	"crepe/storage"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
	"github.com/slack-go/slack"
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

	p.scraper.OnHTML("div.release-entry:first-of-type > div > div > div > div > h4.commit-title:first-of-type > a", func(e *colly.HTMLElement) {

		redislumenVersion, _ := storage.Get("lumenVersion")
		fmt.Println("This is in the Lumen bank: ", redislumenVersion)
		rawText := e.Text
		justTheVersion := strings.Replace(rawText, "v", "", 1)

		if redislumenVersion == justTheVersion {
			fmt.Println("[ DONE ] Lumen already up to date")
		} else {

			lumenSet := storage.Set("lumenVersion", justTheVersion)
			fmt.Println("lumen set error: ", lumenSet)
			fmt.Println("What went into lumenVersion: ", justTheVersion)

			slack.PostWebhook(os.Getenv("SLACK_HOOK_URL"), &slack.WebhookMessage{
				Username: "Crépe",
				Text:     fmt.Sprintf("This is new %s info. Title: %v.", p.config.tech, *e),
			})
		}

	})

	p.scraper.Visit(p.config.URL)

}

func NewLumenParser() *LumenParser {
	return &LumenParser{
		scraper: colly.NewCollector(),
		config: &LumenParserConfig{
			URL:  "https://github.com/laravel/lumen/releases",
			tech: "Lumen",
		},
	}
}
