package main

import (
	"fmt"
	"log"
	"os"

	"askc/internal"
)

func main() {
	webhookURL := os.Getenv("SLACK_WEBHOOK_URL")
	if webhookURL == "" {
		log.Fatal("SLACK_WEBHOOK_URL environment variable not set")
	}
	costMsg, err := internal.GetDailyCost()
	if err != nil {
		log.Fatalf("Error getting AWS cost: %v", err)
	}
	if err := internal.SendToSlack(webhookURL, costMsg); err != nil {
		log.Fatalf("Error sending to Slack: %v", err)
	}
	fmt.Println("Report sent to Slack.")
}
