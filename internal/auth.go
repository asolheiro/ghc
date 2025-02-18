package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const authUrl = "https://api-principal-geral.api.gita.cloud/auth/login"

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
}

func Login(email, password string) (*AuthResponse, error){
	payload, err := json.Marshal(AuthRequest{
		Email: email,
		Password: password,
	})
	if err != nil {
		return nil, fmt.Errorf("error creating JSON Payload, err: %v", err)
	}

	request, err := http.NewRequest(
		"POST",
		authUrl,
		bytes.NewBuffer(payload),
	)
	 if err != nil {
		return nil, fmt.Errorf("error creating request, err : %w", err)
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error sending request, err: %v", err)
	}
	defer response.Body.Close()

	fmt.Println(request.Body)
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code: %d", response.StatusCode)
	}

	var authResponse AuthResponse
	if err := json.NewDecoder(response.Body).Decode(&authResponse); err != nil {
		return nil, fmt.Errorf("error decoding JSON response, err: %w", err)
	}
	return &authResponse, nil
}
	
