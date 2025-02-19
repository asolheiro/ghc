package internal

import (
	"fmt"
	"net/http"

	httputils "github.com/asolheiro/gita-healthcheck/internal/http_utils"
)
const ovhCluster = "e2d718b6919aab2b2cb459520e48f278"

func TestEndpoint(token, clusterId string) (error) {

	url := fmt.Sprintf("https://api-principal-geral.api.gita.cloud/dashboard/all/node/"+ovhCluster)
	body, err := httputils.HttpRequest(
		nil,
		http.MethodGet,
		url, 
		token,

	)
	if err != nil {
		return fmt.Errorf("error requesting, err: %v", err)
	}

	fmt.Println(string(body))
	return nil
}