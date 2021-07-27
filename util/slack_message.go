package util

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

type SlackCrepeField struct {
	Type  string `json:"type,omitempty"`
	Text  string `json:"text,omitempty"`
	Emoji bool   `json:"emoji,omitempty"`
}

type SlackCrepeActions struct {
	Type     string       `json:"type,omitempty"`
	Elements []SlackCrepe `json:"elements"`
}

type SlackCrepe struct {
	Type  string          `json:"type,omitempty"`
	Style string          `json:"style,omitempty"`
	URL   string          `json:"url,omitempty"`
	Text  SlackCrepeField `json:"text,omitempty"`
}

func (c SlackCrepe) BlockType() slack.MessageBlockType {
	return slack.MessageBlockType(c.Type)
}

func (c SlackCrepeActions) BlockType() slack.MessageBlockType {
	return slack.MessageBlockType(c.Type)
}

func SendNewSlackWebhook(pName string, pURL string, newVersion string) {

	crepeHeader := SlackCrepe{
		Type: "header",
		Text: SlackCrepeField{
			Type: "plain_text",
			Text: fmt.Sprintf("焼き立てのクレープ: %s %s！", pName, newVersion),
		},
	}

	crepeBody := SlackCrepe{
		Type: "section",
		Text: SlackCrepeField{
			Type: "mrkdwn",
			Text: "お待たせしました。以下になるボタンをお押しになりごゆっくりどうぞ！",
		},
	}

	crepeButton := SlackCrepeActions{
		Type: "actions",
		Elements: []SlackCrepe{
			{
				Type: "button",
				Text: SlackCrepeField{
					Type:  "plain_text",
					Emoji: true,
					Text:  "いただきます！",
				},
				Style: "primary",
				URL:   pURL,
			},
		},
	}

	slack.PostWebhook(os.Getenv("SLACK_HOOK_URL"), &slack.WebhookMessage{
		Username: "Crépe",
		Blocks: &slack.Blocks{
			BlockSet: []slack.Block{
				crepeHeader,
				crepeBody,
				crepeButton,
			},
		},
	},
	)
}
