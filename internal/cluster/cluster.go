package cluster

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	httputils "github.com/asolheiro/gita-healthcheck/internal/http_utils"
)

const apiUrl = "https://api-principal-geral.api.gita.cloud/node/filter/"
const requestBody = `{
	"page": 1,
	"limit": 1,
	"search": "string",
	"order": "asc"
  }`
type Node struct {}

func GetNodes(token, clusterId string) ([]Node, error) {
	payload, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("error creating JSON Payload, err: %w", err)
	}
	
	request, err := http.NewRequest(
		"POST", 
		apiUrl+clusterId, 
		bytes.NewBuffer(payload),
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

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	fmt.Println(string(body))
	// var responseStruct responseStruct
	// if err := json.NewDecoder(response.Body).Decode(&responseStruct); err != nil {
	// 	return nil, fmt.Errorf("error decoding JSON response: err %w", err)
	// }

	return nil, nil
}

type VersionResponse struct {
	Status string `json:"status"`
	Msg    Msg    `json:"msg"`
}

type Msg struct {
	ID         OID      `json:"_id"`
	Cliente    Cliente  `json:"cliente"`
	DataHora   DateTime `json:"data_hora"`
	Event      string   `json:"event"`
	Kind       string   `json:"kind"`
	ModifiedAt DateTime `json:"modified_at"`
	Version    string   `json:"version"`
}

type OID struct {
	Oid string `json:"$oid"`
}

type Cliente struct {
	ClusterID string `json:"cluster_id"`
}

type DateTime struct {
	Date time.Time `json:"$date"`
}

func GetKubernetesVersion(token, clusterId string) (string, error) {
	body, err := httputils.HttpRequest(
		nil, 
		"GET", 
		"https://api-principal-geral.api.gita.cloud/cluster-infosversion/" + clusterId,
		token,
	)

	if err != nil {
		return "", fmt.Errorf("error requesting data, err: %w", err)
	}

	var response VersionResponse 
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(&response); err != nil {
		return "", fmt.Errorf("error decoding json, err: %w", err)
	}	
	return response.Msg.Version, nil
}