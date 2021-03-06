package crepes

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"crepe/util"

	"github.com/gocolly/colly"
)

type PHPParser struct {
	scraper *colly.Collector
	config  *PHPParserConfig
}

type PHPParserConfig struct {
	Tech string
	URL  string
}

func (p *PHPParser) CheckDate(s *string) bool {
	justTheDate := strings.Replace(*s, "Released: ", "", 1)

	parsedDate, _ := time.Parse("02 Jan 2006", justTheDate)

	postYear, postMonth, postDay := parsedDate.Date()

	if util.IsToday(postYear, int(postMonth), postDay, p.config.Tech) {
		return true
	}

	return false
}

func (p *PHPParser) Scrape() {
	var somethingNew bool

	p.scraper.OnHTML("li", func(e *colly.HTMLElement) {
		// e.Request.Visit(e.Attr("href"))
		// fmt.Println("THE SELECTION ", e.Text)

		dateRegexp := regexp.MustCompile(`^Released: \d{2,2} \w{3,3} \d{4,4}$`)

		if dateRegexp.MatchString(e.Text) {
			if p.CheckDate(&e.Text) {
				somethingNew = true
				// Send to Slack
				util.SendNewSlackWebhook(p.config.Tech, p.config.URL, "")

				// slack.PostWebhook(os.Getenv("SLACK_HOOK_URL"), &slack.WebhookMessage{
				// 	Username: "Crépe",
				// 	Text:     fmt.Sprintf("This is new %s info. Title: %v.", p.config.Tech, *e),
				// })

			}
		}

	})

	p.scraper.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	p.scraper.Visit(p.config.URL)

	if !somethingNew {
		fmt.Printf("[ DONE ] %s has had no new updates.\n", p.config.Tech)
	}

}

func NewPHPParser() *PHPParser {
	return &PHPParser{
		scraper: colly.NewCollector(),
		config: &PHPParserConfig{
			URL:  "https://www.php.net/releases/index.php",
			Tech: "PHP",
		},
	}
}
