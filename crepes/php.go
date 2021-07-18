package crepes

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"crepe/util"

	"github.com/gocolly/colly"
	"github.com/slack-go/slack"
)

type PHPParser struct {
	scraper *colly.Collector
	config  *PHPParserConfig
}

type PHPParserConfig struct {
	tech string
	URL  string
}

func (p *PHPParser) CheckDate(s *string) bool {
	justTheDate := strings.Replace(*s, "Released: ", "", 1)

	parsedDate, _ := time.Parse("02 Jan 2006", justTheDate)

	postYear, postMonth, postDay := parsedDate.Date()

	if util.IsToday(postYear, int(postMonth), postDay) {
		fmt.Println("OKAY")
		return true
	}

	return false
}

func (p *PHPParser) Scrape() {

	p.scraper.OnHTML("li", func(e *colly.HTMLElement) {
		// e.Request.Visit(e.Attr("href"))
		// fmt.Println("THE SELECTION ", e.Text)

		dateRegexp := regexp.MustCompile(`^Released: \d{2,2} \w{3,3} \d{4,4}$`)

		if dateRegexp.MatchString(e.Text) {
			if p.CheckDate(&e.Text) {
				// Send to Slack
				slack.PostWebhook(os.Getenv("SLACK_HOOK_URL"), &slack.WebhookMessage{
					Username: "Crépe",
					Text:     fmt.Sprintf("This is new %s info. Title: %v.", p.config.tech, *e),
				})

			}
		}

	})

	p.scraper.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	p.scraper.Visit(p.config.URL)

}

func NewPHPParser() *PHPParser {
	return &PHPParser{
		scraper: colly.NewCollector(),
		config: &PHPParserConfig{
			URL:  "https://www.php.net/releases/index.php",
			tech: "PHP",
		},
	}
}
