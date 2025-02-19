package problem

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	httputils "github.com/asolheiro/gita-healthcheck/internal/http_utils"
)

type Response struct {
	Status string     `json:"status"`
	Msg    []Security `json:"msg"`
	Total  int        `json:"total"`
}

type Security struct {
	ID           ObjectID  `json:"_id"`
	Type         string    `json:"type"`
	CheckID      string    `json:"check_id"`
	Msg          MsgDetail `json:"msg"`
	Severity     string    `json:"severity"`
	ConfigID     string    `json:"config_id"`
	Kind         string    `json:"kind"`
	UniqueID     string    `json:"id"`
	Name         string    `json:"name"`
	Namespace    string    `json:"namespace"`
	CreationDate TimeStamp `json:"creation_date"`
	DataHora     TimeStamp `json:"data_hora"`
	Cliente      Cliente   `json:"cliente"`
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

type MsgDetail struct {
	ContainerName string `json:"container_name"`
	Reason        string `json:"reason"`
}

func GetProblems(token, clusterId string) (Response, error) {
	url := fmt.Sprintf(
		"https://api-principal-geral.api.gita.cloud/problem/%s?page=1&limit=100",
		clusterId,
	)
	body, err := httputils.HttpRequest(
		nil,
		http.MethodGet,
		url,
		token,
	)
	if err != nil {
		return Response{}, fmt.Errorf("error requesting data, err: %w", err)
	}

	var response Response
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(&response); err != nil {
		return Response{}, fmt.Errorf("error decoding JSON, err: %w", err)
	}
	return response, nil
}