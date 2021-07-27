package crepes

import (
	"crepe/storage"
	"crepe/util"
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

type VueParser struct {
	scraper *colly.Collector
	config  *VueParserConfig
}

type VueParserConfig struct {
	tech string
	URL  string
}

func (p *VueParser) Scrape() {

	p.scraper.OnHTML("div.release-entry:first-of-type > div > div > div > div > div > a", func(e *colly.HTMLElement) {

		redisVueVersion, err := storage.Get("VueVersion")
		if err == nil {
			fmt.Println("This is in the Vue bank: ", redisVueVersion)
		}

		rawText := e.Text
		justTheVersion := strings.Replace(rawText, "v", "", 1)

		if redisVueVersion == justTheVersion {
			fmt.Println("[ DONE ] Vue already up to date")
		} else {

			storage.Set("VueVersion", justTheVersion)
			fmt.Println("What went in: ", justTheVersion)

			util.SendNewSlackWebhook(p.config.tech, p.config.URL, justTheVersion)

			// slack.PostWebhook(os.Getenv("SLACK_HOOK_URL"), &slack.WebhookMessage{
			// 	Username: "Crépe",
			// 	Text:     fmt.Sprintf("This is new %s info. Title: %v.", p.config.tech, *e),
			// })
		}

	})

	p.scraper.Visit(p.config.URL)

}

func NewVueParser() *VueParser {
	return &VueParser{
		scraper: colly.NewCollector(),
		config: &VueParserConfig{
			URL:  "https://github.com/vuejs/vue/releases",
			tech: "Vue",
		},
	}
}
