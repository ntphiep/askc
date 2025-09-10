package internal

import (
	"bytes"
	"fmt"
	"net/http"
)

func SendToSlack(webhookURL, message string) error {
	payload := fmt.Sprintf(`{"text": "%s"}`, message)
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("Slack webhook returned status %d", resp.StatusCode)
	}
	return nil
}
