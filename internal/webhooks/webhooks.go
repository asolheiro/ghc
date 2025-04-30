package webhooks

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// SimpleMessage represents a basic text message
type SimpleMessage struct {
	Text string `json:"text"`
}

// FileAttachment represents a file to be attached to a message
type FileAttachment struct {
	Name     string `json:"name"`
	Content  string `json:"content"` // Base64 encoded content
	FileType string `json:"file_type"`
}

// MessageWithAttachment represents a message with a file attachment
type MessageWithAttachment struct {
	Text        string           `json:"text"`
	Attachments []FileAttachment `json:"attachments,omitempty"`
}

// SendTextMessage sends only a text message to the webhook
func SendTextMessage(webhookURL, message string) error {
	simpleMessage := SimpleMessage{
		Text: message,
	}
	
	jsonData, err := json.Marshal(simpleMessage)
	if err != nil {
		return fmt.Errorf("error marshalling JSON: %v", err)
	}
	
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error sending webhook: %v", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("webhook request failed with status code: %d", resp.StatusCode)
	}
	
	return nil
}

// sendMessageWithPDF sends a text message with a PDF attachment
// For Mattermost: This uses the JSON payload method
func SendMessageWithPDF(webhookURL, message, pdfPath string) error {
	fileData, err := os.ReadFile(pdfPath)
	if err != nil {
		return fmt.Errorf("error reading PDF file: %v", err)
	}
	
	encodedContent := base64.StdEncoding.EncodeToString(fileData)
	
	messageWithAttachment := MessageWithAttachment{
		Text: message,
		Attachments: []FileAttachment{
			{
				Name:     filepath.Base(pdfPath),
				Content:  encodedContent,
				FileType: "application/pdf",
			},
		},
	}
	
	jsonData, err := json.Marshal(messageWithAttachment)
	if err != nil {
		return fmt.Errorf("error marshalling JSON: %v", err)
	}
	
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error sending webhook: %v", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("webhook request failed with status code: %d, response: %s", resp.StatusCode, string(respBody))
	}
	
	return nil
}

// sendPDFMultipart sends a PDF file using multipart/form-data
// This is the method likely needed for Microsoft Teams
func SendPDFMultipart(webhookURL, message, pdfPath string) error {
	// Create a new multipart writer
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	
	// Add text part
	if err := writer.WriteField("text", message); err != nil {
		return fmt.Errorf("error writing text field: %v", err)
	}
	
	// Add file part
	file, err := os.Open(pdfPath)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()
	
	fileContents, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}
	
	// Reset file pointer to beginning
	file.Seek(0, 0)
	
	part, err := writer.CreateFormFile("file", filepath.Base(pdfPath))
	if err != nil {
		return fmt.Errorf("error creating form file: %v", err)
	}
	
	if _, err = part.Write(fileContents); err != nil {
		return fmt.Errorf("error writing file contents: %v", err)
	}
	
	// Close the writer
	if err = writer.Close(); err != nil {
		return fmt.Errorf("error closing writer: %v", err)
	}
	
	// Create request
	req, err := http.NewRequest("POST", webhookURL, body)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}
	
	// Set content type
	req.Header.Set("Content-Type", writer.FormDataContentType())
	
	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()
	
	// Check response status
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("webhook request failed with status code: %d, response: %s", resp.StatusCode, string(respBody))
	}
	
	return nil
}