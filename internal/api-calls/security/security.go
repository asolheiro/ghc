package security

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	httputils "github.com/asolheiro/gita-healthcheck/internal/api-calls/http_utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Response struct {
	Status string          `json:"status" bson:"status"`
	Msg    []SecurityAlert `json:"msg" bson:"msg"`
	Total  int             `json:"total" bson:"total"`
}

type SecurityAlert struct {
	ID struct {
		Oid primitive.ObjectID `json:"$oid" bson:"$oid"`
	} `json:"_id" bson:"_id"`
	Type         string        `json:"type" bson:"type"`
	CheckID      string        `json:"check_id" bson:"check_id"`
	Message      MsgField      `json:"msg" bson:"msg"`
	Severity     SeverityField `json:"severity" bson:"severity"`
	ConfigID     string        `json:"config_id" bson:"config_id"`
	Kind         string        `json:"kind" bson:"kind"`
	UUID         string        `json:"id" bson:"id"`
	Name         string        `json:"name" bson:"name"`
	Namespace    string        `json:"namespace" bson:"namespace"`
	CreationDate struct {
		Date time.Time `json:"$date" bson:"$date"`
	} `json:"creation_date" bson:"creation_date"`
	DataHora struct {
		Date time.Time `json:"$date" bson:"$date"`
	} `json:"data_hora" bson:"data_hora"`
	Cliente          Cliente `json:"cliente" bson:"cliente"`
	VirtuallyDeleted string  `json:"virtually_deleted" bson:"virtually_deleted"`
	Ack              Ack     `json:"ack" bson:"ack"`
}

type Cliente struct {
	ClusterID string `json:"cluster_id" bson:"cluster_id"`
}

type Ack struct {
	State    string     `json:"state" bson:"state"`
	Timeline []Timeline `json:"timeline" bson:"timeline"`
}

type Timeline struct {
	ID      string `json:"_id" bson:"_id"`
	User    string `json:"user" bson:"user"`
	Email   string `json:"email" bson:"email"`
	Message string `json:"message" bson:"message"`
	Date    struct {
		Date time.Time `json:"$date" bson:"$date"`
	} `json:"date" bson:"date"`
	State string `json:"state" bson:"state"`
}

type MsgField struct {
	Text  string                 // Caso seja uma string
	Other map[string]interface{} // Caso seja um objeto
}

type SeverityField []string

func GetSecurity(token, clusterId string) (Response, error) {
	url := fmt.Sprintf("https://api-principal-geral.api.gita.cloud/security/%s?page=1&limit=100", clusterId)
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
		return Response{}, fmt.Errorf("error decoding JSON response, err: %w", err)
	}

	return response, nil
}
