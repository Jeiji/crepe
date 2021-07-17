package crepes

import (
	"fmt"
	"regexp"
	"time"
	"crepe/util"
	"strconv"

	"github.com/gocolly/colly"
	// "github.com/slack-go/slack"
	
)

type DockerParser struct {
	scraper *colly.Collector
	config  *DockerParserConfig
}

type DockerParserConfig struct {
	tech string
	URL string
}

func (p *DockerParser) CreateDate(s *string) {

	yearRegex :=  regexp.MustCompile(`\d{4,}`)
	monthRegex :=  regexp.MustCompile(`\d{4,4}-(\d{2,2})`)
	dayRegex :=  regexp.MustCompile(`\d{2,}$`)

	postYear, _ := strconv.Atoi(yearRegex.FindString(*s))
	postMonth, _ := strconv.Atoi(monthRegex.FindStringSubmatch(*s)[1])
	postDay, _ := strconv.Atoi(dayRegex.FindString(*s))


	fmt.Println(util.TodayDay)

	
	if util.TodayYear == postYear ||
			util.TodayMonth == time.Month(postMonth) ||
			util.TodayDay == postDay {
		}

	// fmt.Println(time.Parse("0000-00-00", *s))
}


func (p *DockerParser) Scrape() {

	p.scraper.OnHTML("h2 ~ p", func(e *colly.HTMLElement) {
		// e.Request.Visit(e.Attr("href"))
		fmt.Println("THE SELECTION ", e.Text)

		dateRegexp :=  regexp.MustCompile(`^\d{4,}-\d{2,}-\d{2,}$`)

		if dateRegexp.MatchString(e.Text) {
			p.CreateDate(&e.Text)
		}


	})

	p.scraper.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	p.scraper.Visit(p.config.URL)

	// Send to Slack

}

func NewDockerDesktopWindowsParser() *DockerParser {
	return &DockerParser{
		scraper: colly.NewCollector(),
		config: &DockerParserConfig{
			URL: "https://docs.docker.com/docker-for-windows/release-notes/",
			tech: "Docker Desktop (Windows)",
		},
	}
}

func NewDockerDesktopMacParser() *DockerParser {
	return &DockerParser{
		scraper: colly.NewCollector(),
		config: &DockerParserConfig{
			URL: "https://docs.docker.com/docker-for-mac/release-notes/",
			tech: "Docker Desktop (Mac)",
		},
	} 
}



