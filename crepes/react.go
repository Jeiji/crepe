package crepes

import (
	"crepe/util"
	"fmt"

	"github.com/mmcdole/gofeed"
	"github.com/slack-go/slack"
)

type ReactParser struct {
	scraper *gofeed.Parser
	config  *ReactParserConfig
}

type ReactParserConfig struct {
	URL string
}

func (p *ReactParser) Scrape() {
	feed, _ := p.scraper.ParseURL(p.config.URL)
	// fmt.Printf("+%v", feed)
	for _, item := range feed.Items {

		t := item.PublishedParsed

		if item.UpdatedParsed != nil {
			t = item.UpdatedParsed
		}

		itemYear, itemMonth, itemDay := t.Date()

		if !util.IsToday(itemYear, int(itemMonth), itemDay) {
			break
		}

		// fmt.Println("Sent to Srack")
		slack.PostWebhook("https://hooks.slack.com/services/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", &slack.WebhookMessage{
			Username: "Crépe",
			Text:     item.Title,
		})
		fmt.Println(item.Title)

	}
}

func NewReactParser() *ReactParser {
	return &ReactParser{
		scraper: gofeed.NewParser(),
		config: &ReactParserConfig{
			URL: "https://reactjs.org/feed.xml",
		},
	}
}
