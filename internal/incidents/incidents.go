package incidents

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
	Msg    []Incident `json:"msg"`
	Total  int        `json:"total"`
}

type Incident struct {
	ID               ID           `json:"_id"`
	Type             string       `json:"type"`
	CheckID          string       `json:"check_id"`
	Msg              AlertMessage `json:"msg"`
	Kind             string       `json:"kind"`
	ConfigID         string       `json:"config_id"`
	Severity         string       `json:"severity"`
	AlertID          string       `json:"id"`
	Name             string       `json:"name"`
	Namespace        string       `json:"namespace"`
	CreationDate     CustomTime   `json:"creation_date"` 
	DataHora         CustomTime   `json:"data_hora"`     
	Cliente          Cliente      `json:"cliente"`
	VirtuallyDeleted string       `json:"virtually_deleted"`
	Deleted          string       `json:"deleted"`
	Ack              Ack          `json:"ack"`
}

type CustomTime struct {
	time.Time
}

type ID struct {
	OID string `json:"$oid"`
}

type AlertMessage struct {
	State  string `json:"state"`
	Reason string `json:"reason"`
}

type Cliente struct {
	ClusterID string `json:"cluster_id"`
}

type Ack struct {
	State    string        `json:"state"`
	Timeline []AckTimeline `json:"timeline"`
}

type AckTimeline struct {
	ID      string     `json:"_id"`
	User    string     `json:"user"`
	Email   string     `json:"email"`
	Message string     `json:"message"`
	Date    CustomTime `json:"date"`
	State   string     `json:"state"`
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	var wrapper struct {
		Date string `json:"$date"`
	}

	if err := json.Unmarshal(b, &wrapper); err != nil {
		return err
	}

	// Parse the date string
	parsedTime, err := time.Parse(time.RFC3339, wrapper.Date)
	if err != nil {
		return err
	}

	ct.Time = parsedTime
	return nil
}

func GetIncidents(token, clusterId string) (Response, error) {
	url := fmt.Sprintf(
		"https://api-principal-geral.api.gita.cloud/incident/%s?page=1&limit=100",
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
		return Response{}, fmt.Errorf("error decoding JSON response, err: %w", err)
	}
	return response, nil
}

type rawAlertMessage struct {
    State  string `json:"state"`
    Reason string `json:"reason"`
}

// UnmarshalJSON implements custom unmarshaling for AlertMessage
func (am *AlertMessage) UnmarshalJSON(data []byte) error {
    // First, try to unmarshal as a string
    var stringMsg string
    if err := json.Unmarshal(data, &stringMsg); err == nil {
        // If successful, treat the string as the message
        am.State = "unknown"
        am.Reason = stringMsg
        return nil
    }

    // If that fails, try to unmarshal as an object
    var objMsg rawAlertMessage
    if err := json.Unmarshal(data, &objMsg); err != nil {
        return fmt.Errorf("failed to unmarshal AlertMessage: %v", err)
    }

    // Copy the values from the raw message
    am.State = objMsg.State
    am.Reason = objMsg.Reason
    return nil
}