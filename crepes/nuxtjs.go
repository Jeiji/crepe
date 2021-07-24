package crepes

import (
	"crepe/storage"
	"fmt"
	"os"
	"strings"

	"github.com/gocolly/colly"
	"github.com/slack-go/slack"
)

type NuxtjsParser struct {
	scraper *colly.Collector
	config  *NuxtjsParserConfig
}

type NuxtjsParserConfig struct {
	tech string
	URL  string
}

func (p *NuxtjsParser) Scrape() {

	p.scraper.OnHTML("div.release-entry:first-of-type > div > div > div > div > div > a", func(e *colly.HTMLElement) {

		redisNuxtjsVersion, err := storage.Get("NuxtjsVersion")
		if err == nil {
			fmt.Println("This is in the Nuxtjs bank: ", redisNuxtjsVersion)
		}

		rawText := e.Text
		justTheVersion := strings.Replace(rawText, "v", "", 1)

		if redisNuxtjsVersion == justTheVersion {
			fmt.Println("[ DONE ] Nuxtjs already up to date")
		} else {

			storage.Set("NuxtjsVersion", justTheVersion)
			fmt.Println("[ NEW ] (Nuxtjs) What went in: ", justTheVersion)

			slack.PostWebhook(os.Getenv("SLACK_HOOK_URL"), &slack.WebhookMessage{
				Username: "Crépe",
				Text:     fmt.Sprintf("This is new %s info. Title: %v.", p.config.tech, *e),
			})
		}

	})

	p.scraper.Visit(p.config.URL)

}

func NewNuxtjsParser() *NuxtjsParser {
	return &NuxtjsParser{
		scraper: colly.NewCollector(),
		config: &NuxtjsParserConfig{
			URL:  "https://github.com/nuxt/nuxt.js/releases",
			tech: "Nuxtjs",
		},
	}
}
