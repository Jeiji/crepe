package crepes

import (
	"fmt"
	"regexp"

	"crepe/util"
	"strconv"

	"github.com/gocolly/colly"
)

type DockerParser struct {
	scraper *colly.Collector
	config  *DockerParserConfig
}

type DockerParserConfig struct {
	Tech string
	URL  string
}

func (p *DockerParser) CheckDate(s *string) bool {

	yearRegex := regexp.MustCompile(`\d{4,}`)
	monthRegex := regexp.MustCompile(`\d{4,4}-(\d{2,2})`)
	dayRegex := regexp.MustCompile(`\d{2,}$`)

	postYear, _ := strconv.Atoi(yearRegex.FindString(*s))
	postMonth, _ := strconv.Atoi(monthRegex.FindStringSubmatch(*s)[1])
	postDay, _ := strconv.Atoi(dayRegex.FindString(*s))

	return util.IsToday(postYear, postMonth, postDay, p.config.Tech)

	// return false
}

func (p *DockerParser) Scrape() {

	var somethingNew bool

	p.scraper.OnHTML("h2 ~ p", func(e *colly.HTMLElement) {
		// e.Request.Visit(e.Attr("href"))
		// fmt.Println("THE SELECTION ", e.Text)

		dateRegexp := regexp.MustCompile(`^\d{4,}-\d{2,}-\d{2,}$`)

		if dateRegexp.MatchString(e.Text) {
			if p.CheckDate(&e.Text) {
				somethingNew = true
				// Send to Slack
				util.SendNewSlackWebhook(p.config.Tech, p.config.URL, "")
				fmt.Println("[ NEW ] (Docker)")

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

func NewDockerDesktopWindowsParser() *DockerParser {
	return &DockerParser{
		scraper: colly.NewCollector(),
		config: &DockerParserConfig{
			URL:  "https://docs.docker.com/docker-for-windows/release-notes/",
			Tech: "Docker Desktop (Windows)",
		},
	}
}

func NewDockerDesktopMacParser() *DockerParser {
	return &DockerParser{
		scraper: colly.NewCollector(),
		config: &DockerParserConfig{
			URL:  "https://docs.docker.com/docker-for-mac/release-notes/",
			Tech: "Docker Desktop (Mac)",
		},
	}
}
