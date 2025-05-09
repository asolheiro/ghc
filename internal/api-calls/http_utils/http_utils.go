package httputils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func HttpRequest(pl any, method, url, token string) ([]byte, error) {
	payload, err := json.Marshal(pl)
	if err != nil {
		return nil, fmt.Errorf("error creating JSON Payload, err: %v", err)
	}

	request, err := http.NewRequest(
		method,
		url,
		bytes.NewBuffer(payload),
	)
	if err != nil {
		return nil, fmt.Errorf("error creating request, err : %w", err)
	}

	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error sending request, err: %v", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code: %d. err: %s", response.StatusCode, string(body))
	}

	return body, nil
}
