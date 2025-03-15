package version

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/asolheiro/gita-healthcheck/internal/maps"
)

type ResponseEOL struct {
	Cycle             string `json:"cycle"`
	ReleaseDate       string `json:"releaseDate"`
	EndOfLife         string `json:"eol"`
	Latest            string `json:"latest"`
	LatestReleaseDate string `json:"latestReleaseDate"`
	LTS               bool   `json:"lts"`
	Support           string `json:"support"`
}

var CycleInfoRKE = ResponseEOL{
	Cycle:             "0.00.00",
	ReleaseDate:       "0000-00-00",
	EndOfLife:         "0000-00-00",
	Latest:            "0.00.00",
	LatestReleaseDate: "0000-00-00",
	LTS:               false,
	Support:           "0000-00-00",
}

func GetCycleInfo(product, cycle string) (ResponseEOL, error) {
	if product == "RKE" {
		return CycleInfoRKE, nil
	}

	url := fmt.Sprintf("https://endoflife.date/api/%s/%s.json", product, cycle)

	req, err := http.NewRequest(
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		return ResponseEOL{}, fmt.Errorf("error creating request, err: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return ResponseEOL{}, fmt.Errorf("error sending request, err: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return ResponseEOL{}, fmt.Errorf("response with status code %d. %s", res.StatusCode, res.Body)
	}

	var response ResponseEOL
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return ResponseEOL{}, fmt.Errorf("error decoding JSON response, err: %w", err)
	}
	return response, nil
}

func DaysUntilEndOfSupport(endDateStr string) (int, error) {
	const layout = "2006-01-02"

	endDate, err := time.Parse(layout, endDateStr)
	if err != nil {
		return 0, fmt.Errorf("error parsing date, err: %v", err)
	}

	today := time.Now()

	daysRemaining := int(endDate.Sub(today).Hours() / 24)

	return daysRemaining, nil
}


func IdentifyProductType(version string) (string, string) {
	re := regexp.MustCompile(`v(\d+\.\d+)`)

	matches := re.FindStringSubmatch(version)
	if len(matches) < 2 {
		return "Unknown", "Unknown"
	}
	cycle := matches[1]

	switch {
	case strings.Contains(version, "-gke"):
		return "gke", cycle
	case strings.Contains(version, "+rke2"):
		return "RKE", cycle
	case strings.Contains(version, "-eks"):
		return "eks", cycle
	default:
		return "kubernetes", cycle
	}
}

type KubernetesVersionResponse struct {
	Version       string
	EndOfSupport string
	LastRelease   string
	SupportStatus string
}

func GatherKubernetesInfo(version string) (KubernetesVersionResponse, error) {
	product, cycle := IdentifyProductType(version)
	cycleInfo, err := GetCycleInfo(product, cycle)
	if err != nil {
		return KubernetesVersionResponse{}, fmt.Errorf("error requesting cycle version info, err: %w", err)
	}

	if cycleInfo == CycleInfoRKE {
		return KubernetesVersionResponse{
			Version:       version,
			EndOfSupport: "---",
			LastRelease: "[RKE2-Releases](https://github.com/rancher/rke2/releases)",
			SupportStatus: "[Support-Matrix](https://www.suse.com/pt-br/suse-rke1/support-matrix/all-supported-versions/rke1-v1-31/)",
		}, nil
	}
	
	daysRemaining, err := DaysUntilEndOfSupport(cycleInfo.EndOfLife)
	if err != nil {
		return KubernetesVersionResponse{}, fmt.Errorf("error retrieving remaing days")
	}
	colorMark := maps.MapDaysToColorMark(daysRemaining)

	endOfSupport, err := time.Parse("2006-01-02", cycleInfo.EndOfLife)
	if err != nil {
		return KubernetesVersionResponse{}, fmt.Errorf("error parsing date, err: %v", err)
	}

	
	return KubernetesVersionResponse{
		Version:       version,
		EndOfSupport: endOfSupport.Format("02/01/2006"),
		LastRelease: cycleInfo.Latest,
		SupportStatus: colorMark,
	}, nil
}
