package crepes

import (
	"crepe/storage"
	"crepe/util"
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type NextjsParser struct {
	scraper *colly.Collector
	config  *NextjsParserConfig
}

type NextjsParserConfig struct {
	tech string
	URL  string
}

func (p *NextjsParser) Scrape() {

	p.scraper.OnHTML("div.release-entry:first-of-type > div > div > div > div > div > a", func(e *colly.HTMLElement) {

		redisNextjsVersion, err := storage.Get("NextjsVersion")
		if err == nil {
			fmt.Println("This is in the Nextjs bank: ", redisNextjsVersion)
		}

		rawText := e.Text
		justTheVersion := strings.Replace(rawText, "v", "", 1)

		if redisNextjsVersion == justTheVersion {
			fmt.Println("[ DONE ] Nextjs already up to date")
		} else {

			storage.Set("NextjsVersion", justTheVersion)
			fmt.Println("[ NEW ] (NextJS) What went in: ", justTheVersion)

			util.SendNewSlackWebhook(p.config.tech, p.config.URL, justTheVersion)

			// slack.PostWebhook(os.Getenv("SLACK_HOOK_URL"), &slack.WebhookMessage{
			// 	Username: "Crépe",
			// 	Text:     fmt.Sprintf("This is new %s info. Title: %v.", p.config.tech, *e),
			// })
		}

	})

	p.scraper.Visit(p.config.URL)

}

func NewNextjsParser() *NextjsParser {
	return &NextjsParser{
		scraper: colly.NewCollector(),
		config: &NextjsParserConfig{
			URL:  "https://github.com/vercel/next.js/releases",
			tech: "Nextjs",
		},
	}
}
