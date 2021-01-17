package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/slack-go/slack"
)

func handler(event events.CloudWatchEvent) error {

	spend, err := describeSpend()
	findings, err := parse(event.Detail)

	if err != nil {
		return err
	}

	if findings != nil {
		for _, finding := range *findings {
			err := finding.postWebhook()
			if err != nil {
				return err
			}
		}
		return nil
	}

	message := fmt.Sprintf("実績値: %s USD、月末の予測値: %s USD", *spend.ActualSpend.Amount, *spend.ForecastedSpend.Amount)
	webhookMessage := &slack.WebhookMessage{Text: message}
	return slack.PostWebhook(incomingWebhookURL, webhookMessage)
}

func main() {
	lambda.Start(handler)
}
