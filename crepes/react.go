package crepes

import (
	"crepe/util"
	"fmt"

	"github.com/mmcdole/gofeed"
)

type ReactParser struct {
	scraper *gofeed.Parser
	config  *ReactParserConfig
}

type ReactParserConfig struct {
	URL  string
	Tech string
}

func (p *ReactParser) Scrape() {
	feed, fError := p.scraper.ParseURL(p.config.URL)

	if fError != nil {
		fmt.Println("\n React Crepe ran into a problem...")
		util.SendNewSlackNotification("焦げられたクレープ！", fmt.Sprintf("%s クレープはエラーが発生しちゃった。クレープ機をチェックしてください！", p.config.Tech), p.config.Tech)
	} else {
		for _, item := range feed.Items {

			t := item.PublishedParsed

			if item.UpdatedParsed != nil {
				t = item.UpdatedParsed
			}

			itemYear, itemMonth, itemDay := t.Date()

			if !util.IsToday(itemYear, int(itemMonth), itemDay, p.config.Tech) {
				fmt.Println("[ DONE ] React already up to date")
				break
			}

			// fmt.Println("Sent to Srack")
			util.SendNewSlackWebhook(p.config.Tech, p.config.URL, "")
			// slack.PostWebhook("https://hooks.slack.com/services/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", &slack.WebhookMessage{
			// 	Username: "Crépe",
			// 	Text:     item.Title,
			// })
			fmt.Println(item.Title)

		}
	}
	// fmt.Printf("+%v", feed)

}

func NewReactParser() *ReactParser {
	return &ReactParser{
		scraper: gofeed.NewParser(),
		config: &ReactParserConfig{
			URL:  "https://reactjs.org/feed.xml",
			Tech: "React",
		},
	}
}
