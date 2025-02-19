package organizations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const apiUrl = "https://api-cluster.api.gita.cloud/organization"

type responseStruct struct {
	Status string `json:"status"`
	Organizations []Organization `json:"msg"`
}

type Organization struct {
	Id string `json:"id"`
	Name string `json:"organization_name"`
	PermissionType string `json:"permission_type"`
	StripeCustomerId string `json:"stripe_customer_id"`
	Disabled int `json:"disabled"`
	Ticket int `json:"ticket"`
	OwnerId int `json:"owner_id"`
}

func ListOrgs(token string) ([]Organization, error) {
	request, err := http.NewRequest(
		"GET", 
		apiUrl, 
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	
	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Set("Accept", "application/json")

	
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer response.Body.Close()

	var responseStruct responseStruct
	if err := json.NewDecoder(response.Body).Decode(&responseStruct); err != nil {
		return nil, fmt.Errorf("error decoding JSON response: err %w", err)
	}

	return responseStruct.Organizations, nil
}