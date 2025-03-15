package alerts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	httputils "github.com/asolheiro/gita-healthcheck/internal/api-calls/http_utils"
)

type AlertResponse struct {
	Status string `json:"status"`
	Msg    []Msg  `json:"msg"`
	Total  int    `json:"total"`
}

type Msg struct {
	ID           ObjectID  `json:"_id"`
	Type         string    `json:"type"`
	CheckID      string    `json:"check_id"`
	Kind         string    `json:"kind"`
	Severity     string    `json:"severity"`
	ConfigID     string    `json:"config_id"`
	Name         string    `json:"name"`
	Msg          string    `json:"msg"`
	Namespace    string    `json:"namespace"`
	Cliente      Cliente   `json:"cliente"`
	CreationDate TimeStamp `json:"creation_date"`
	DataHora     TimeStamp `json:"data_hora"`
	Ack          Ack       `json:"ack"`
}

type ObjectID struct {
	OID string `json:"$oid"`
}

type TimeStamp struct {
	Date time.Time `json:"$date"`
}

type Cliente struct {
	ClusterID string `json:"cluster_id"`
}

type Ack struct {
	State    string     `json:"state"`
	Timeline []Timeline `json:"timeline"`
}

type Timeline struct {
	ID      string    `json:"_id"`
	User    string    `json:"user"`
	Email   string    `json:"email"`
	Message string    `json:"message"`
	Date    TimeStamp `json:"date"`
	State   string    `json:"state"`
}

func GetAlerts(token, clusterId string) ([]Msg, error) {
	url := fmt.Sprintf("https://api-principal-geral.api.gita.cloud/alert/%s?page=1&limit=100", clusterId)
	body, err := httputils.HttpRequest(
		nil,
		http.MethodGet,
		url,
		token,
	)
	if err != nil {
		return nil, fmt.Errorf("error requesting data, err: %w", err)
	}

	var response AlertResponse
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(&response); err != nil {
		return nil, fmt.Errorf("error decoding json, err: %w", err)
	}
	return response.Msg, nil
}
