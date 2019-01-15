package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kyokomi/emoji"
	"github.com/weAutomateEverything/serverless-alerting/alert/lambda/client"
	client2 "github.com/weAutomateEverything/serverless-alerting/alert/text/client"
	"strings"
)

func main() {
	lambda.Start(Handle)
}

func Handle(event events.CloudWatchEvent) error {
	var detail Detail
	err := json.Unmarshal(event.Detail, &detail)
	if err != nil {
		client.LogLambdaError(err)
		return err
	}

	err = client2.SendError(emoji.Sprintf(":interrobang: *New Macie issue detected*\n"+
		"Description: %v\n"+
		"Risk: %v\n"+
		"Type:%v\n"+
		"URL:%v\n", strings.Replace(detail.Name,"_"," ",-1), detail.RiskScore, detail.NotificationType, detail.URL))

	if err != nil {
		client.LogLambdaError(err)
		return err
	}

	return nil
}

type Detail struct {
		NotificationType string `json:"notification-type"`
		Name             string `json:"name"`
		URL              string `json:"url"`
		RiskScore        int    `json:"risk-score"`
	}
