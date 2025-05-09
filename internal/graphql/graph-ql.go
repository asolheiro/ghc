package graphql

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
)

// GraphQLRequest represents a GraphQL request structure
type GraphQLRequest struct {
	Query     string    `json:"query"`
	Variables Variables `json:"variables,omitempty"`
}

type Variables struct {
	Content     string   `json:"content"`
	Description string   `json:"description"`
	Editor      string   `json:"editor"`
	IsPublished bool     `json:"isPublished"`
	IsPrivate   bool     `json:"isPrivate"`
	Locale      string   `json:"locale"`
	Path        string   `json:"path"`
	Title       string   `json:"title"`
	Tags        []string `json:"tags"`
}

// GraphQLResponse represents the response from the GraphQL API
type GraphQLResponse struct {
	Data   map[string]any   `json:"data"`
	Errors []map[string]any `json:"errors"`
}

var mutation = `
mutation CreatePage($content: String!, $description: String!, $editor: String!, $isPublished: Boolean!, $isPrivate: Boolean!, $locale: String!, $path: String!, $title: String!, $tags: [String]!) {
	pages {
		create(
			content: $content
			description: $description
			editor: $editor
			isPublished: $isPublished
			isPrivate: $isPrivate
			locale: $locale
			path: $path
			title: $title
			tags: $tags
		) {
			responseResult {
				succeeded
				slug
				message
			}
		}
	}
}`

// GrapQLPost send a HTTP POST request to create a page with the file in markdownFilePath
func GraphQLPost(markdownFilePath string) error {
	if err := godotenv.Load(); err != nil {
		return errors.New("error loading .env file")
	}

	wikiJsBaseURL := os.Getenv("WIKIJS_URL")
	apiToken := os.Getenv("WIKIJS_API_TOKEN")
	if wikiJsBaseURL == "" || apiToken == "" {
		return errors.New("WIKIJS_URL and WIKIJS_API_TOKEN must be set in .env file")
	}

	content, err := os.ReadFile(markdownFilePath)
	if err != nil {
		return fmt.Errorf("error reading markdown file: %v", err)
	}

	fileName := filepath.Base(markdownFilePath)
	parts := strings.Split(fileName, "_")
	if len(parts) < 2 {
		return fmt.Errorf("invalid markdown file format: %s", fileName)
	}

	organization := parts[0]
	dateStr := strings.TrimSuffix(parts[1], filepath.Ext(parts[1]))
	dateParts := strings.Split(dateStr, "-")
	if len(dateParts) != 3 {
		return fmt.Errorf("invalid date format: %s", dateStr)
	}

	day, month, year := dateParts[0], dateParts[1], dateParts[2]
	pageTitle := fmt.Sprintf("%s %s-%s-%s", organization, day, month, year)

	pagePath := fmt.Sprintf("docs/experts/projetos/%s/healthcheck/hc%s-%s-%s", organization, day, month, year)

	pageContent := string(content)
	pageDescription := fmt.Sprintf("Content from %s", markdownFilePath)
	pageTags := []string{"documentation", "healthcheck", "gita", organization}
	isPrivate := false

	variables := Variables{
		Content:     pageContent,
		Description: pageDescription,
		Editor:      "markdown",
		IsPublished: true,
		IsPrivate:   isPrivate,
		Locale:      "pt-br",
		Path:        pagePath,
		Title:       pageTitle,
		Tags:        pageTags,
	}

	reqBody := GraphQLRequest{
		Query:     mutation,
		Variables: variables,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %v", err)
	}

	req, err := http.NewRequest("POST", wikiJsBaseURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %v", err)
	}

	var result GraphQLResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return fmt.Errorf("error parsing response: %v", err)
	}

	if err := handleGraphQLResponse(result); err != nil {
		return fmt.Errorf("error at GraphQL response, err: %w", err)
	}

	fmt.Println("Healthcheck report available in: https://wiki.jackexperts.com/" + pagePath)
	return nil
}

// handleGraphQLResponse check if there's an error in the Data field
func handleGraphQLResponse(result GraphQLResponse) error {
	if result.Errors != nil {
		var errMessages []string
		for _, err := range result.Errors {
			if msg, ok := err["message"].(string); ok {
				errMessages = append(errMessages, msg)
			}
		}
		return fmt.Errorf("GraphQL errors: %v", errMessages)
	}

	if responseResult, ok := result.Data["pages"].(map[string]any)["create"].(map[string]any)["responseResult"].(map[string]any); ok {
		if msg, ok := responseResult["message"].(string); ok && msg == "Cannot create this page because an entry already exists at the same path." {
			return fmt.Errorf("page creation failed: %s", msg)
		}
	}

	return nil
}
