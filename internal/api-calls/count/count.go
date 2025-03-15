package count

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	httputils "github.com/asolheiro/gita-healthcheck/internal/api-calls/http_utils"
)

const apiUrl = "https://api-cluster.api.gita.cloud/count/"

type CountResponse struct {
    Status string `json:"status"`
    Msg    []Msg  `json:"msg"`
}

type Msg struct {
    Organization Organization `json:"organization"`
    Clusters     []Cluster   `json:"clusters"`
}

type Organization struct {
    ID                string `json:"id"`
    Name              string `json:"name"`
    UserPermission    string `json:"userPermission"`
    StripeCustomerID  string `json:"stripeCustomerId,omitempty"`
    Disabled          int    `json:"disabled"`
    Ticket           int    `json:"ticket"`
    OwnerID          int    `json:"owner_id"`
    OwnerCreationTime string `json:"owner_creation_time"`
}

type Cluster struct {
    UserPermission        string `json:"userPermission"`
    ClusterID            string `json:"clusterId"`
    Name                 string `json:"name"`
    KubernetesVersion    string `json:"kubernetes_version"`
    ProxyProtocol        string `json:"proxyProtocol"`
    ProxyUser            string `json:"proxyUser"`
    ProxyPassword        string `json:"proxyPassword"`
    ProxyHost            string `json:"proxyHost"`
    ProxyPort            string `json:"proxyPort"`
    EnableLogs           bool   `json:"enableLogs"`
    EnableDeprecations   bool   `json:"enableDeprecations"`
    EnableExec           bool   `json:"enableExec"`
    EnableOpenshift      bool   `json:"enableOpenshift"`
    InstalledLogs        bool   `json:"installedLogs"`
    InstalledDeprecations bool  `json:"installedDeprecations"`
    InstalledExec        bool   `json:"installedExec"`
    Stopped              bool   `json:"stopped"`
    Running              string `json:"running"`
    Incident             int    `json:"incident"`
    Problem              int    `json:"problem"`
    Security             int    `json:"security"`
    Alert                int    `json:"alert"`
    Node                 int    `json:"node"`
}

func GetUserCount(token string) (CountResponse, error) {
	body, err := httputils.HttpRequest(
		nil,
		http.MethodGet,
		apiUrl,
		token,
	)

	if err != nil {
		return CountResponse{}, fmt.Errorf("error requesting data, err: %w", err)
	}
	var response CountResponse
	if err := json.NewDecoder(bytes.NewReader(body)).Decode(&response); err != nil {
		return CountResponse{}, fmt.Errorf("error decoding JSON response, err: %w", err)
	}
	return response, nil
}