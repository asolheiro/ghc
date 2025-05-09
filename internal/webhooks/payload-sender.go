package webhooks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Helper function to send payload to webhook
func sendWebhookPayload(webhookURL string, payload any) error {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshalling JSON: %v", err)
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error sending webhook: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		// Read response body for more details
		var body []byte
		body, _ = io.ReadAll(resp.Body)
		return fmt.Errorf("webhook request failed with status code: %d, response: %s", resp.StatusCode, string(body))
	}

	return nil
}
