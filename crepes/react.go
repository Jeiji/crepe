package crepes

import (
	"fmt"
	"time"

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

		today := time.Now()
		// today := time.Date(2021, time.June, 8, 0, 0, 0, 0, time.Local)

		t := item.PublishedParsed

		if item.UpdatedParsed != nil {
			t = item.UpdatedParsed
		}

		todayYear, todayMonth, todayDay := today.Date()
		itemYear, itemMonth, itemDay := t.Date()

		if todayYear != itemYear ||
			todayMonth != itemMonth ||
			todayDay != itemDay {
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
