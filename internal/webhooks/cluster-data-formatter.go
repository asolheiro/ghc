package webhooks

import (
	"log"
	"sync"
	"time"

	"github.com/asolheiro/gita-healthcheck/internal/api-calls/alerts"
	"github.com/asolheiro/gita-healthcheck/internal/api-calls/auth"
	"github.com/asolheiro/gita-healthcheck/internal/api-calls/count"
	"github.com/asolheiro/gita-healthcheck/internal/api-calls/incidents"
	"github.com/asolheiro/gita-healthcheck/internal/api-calls/metrics"
	"github.com/asolheiro/gita-healthcheck/internal/api-calls/problem"
	"github.com/asolheiro/gita-healthcheck/internal/api-calls/security"
	"github.com/asolheiro/gita-healthcheck/internal/api-calls/version"
	"github.com/asolheiro/gita-healthcheck/internal/maps"
)

// FindClustersInfo finds information about clusters
func FindClustersInfo(auth *auth.AuthResponse, msgCount count.Msg, cluster count.Cluster) KubernetesReportData {
	var wg sync.WaitGroup
	wg.Add(7)

	incidentsCh := make(chan incidents.Response, 1)
	incidentsErrCh := make(chan error, 1)
	go func() {
		defer wg.Done()
		incidents, err := incidents.GetIncidents(auth.AccessToken, cluster.ClusterID)
		if err != nil {
			incidentsErrCh <- err
			return
		}
		incidentsCh <- incidents
	}()

	problemCh := make(chan problem.Response, 1)
	problemErrCh := make(chan error, 1)
	go func() {
		defer wg.Done()
		problems, err := problem.GetProblems(auth.AccessToken, cluster.ClusterID)
		if err != nil {
			problemErrCh <- err
			return
		}
		problemCh <- problems
	}()

	securityCh := make(chan security.Response, 1)
	securityErrCh := make(chan error, 1)
	go func() {
		defer wg.Done()
		sec, err := security.GetSecurity(auth.AccessToken, cluster.ClusterID)
		if err != nil {
			securityErrCh <- err
			return
		}
		securityCh <- sec
	}()

	clusterMetricsCh := make(chan metrics.TotalMetrics, 1)
	clusterMetricsErrCh := make(chan error, 1)
	go func() {
		defer wg.Done()
		metrics, err := metrics.GetClusterMetrics(auth.AccessToken, cluster.ClusterID)
		if err != nil {
			clusterMetricsErrCh <- err
			return
		}
		clusterMetricsCh <- metrics
	}()

	nodeMetricsCh := make(chan []metrics.TotalNodeMetrics, 1)
	nodeMetricsErrCh := make(chan error, 1)
	go func() {
		defer wg.Done()
		metrics, err := metrics.GetNodeMetrics(auth.AccessToken, cluster.ClusterID)
		if err != nil {
			nodeMetricsErrCh <- err
			return
		}
		nodeMetricsCh <- metrics
	}()

	alertsCh := make(chan []alerts.Msg, 1)
	alertsErrCh := make(chan error, 1)
	go func() {
		defer wg.Done()
		alertsList, err := alerts.GetAlerts(auth.AccessToken, cluster.ClusterID)
		if err != nil {
			alertsErrCh <- err
			return
		}
		alertsCh <- alertsList
	}()

	k8sInfoCh := make(chan version.KubernetesVersionResponse, 1)
	k8sInfoErrCh := make(chan error, 1)

	go func() {
		defer wg.Done()
		k8sInfo, err := version.GatherKubernetesInfo(cluster.KubernetesVersion)
		if err != nil {
			k8sInfoErrCh <- err
			return
		}
		k8sInfoCh <- k8sInfo
	}()

	wg.Wait()

	var incidentsData incidents.Response
	var problemData problem.Response
	var securityData security.Response
	var clusterMetricsData metrics.TotalMetrics
	var nodeMetricsData []metrics.TotalNodeMetrics
	var alertsDataRaw []alerts.Msg
	var k8sInfoData version.KubernetesVersionResponse

	select {
	case err := <-incidentsErrCh:
		log.Fatal(err)
	default:
		incidentsData = <-incidentsCh
	}

	select {
	case err := <-problemErrCh:
		log.Fatal(err)
	default:
		problemData = <-problemCh
	}

	select {
	case err := <-securityErrCh:
		log.Fatal(err)
	default:
		securityData = <-securityCh
	}

	select {
	case err := <-clusterMetricsErrCh:
		log.Fatal(err)
	default:
		clusterMetricsData = <-clusterMetricsCh
	}

	select {
	case err := <-nodeMetricsErrCh:
		log.Fatal(err)
	default:
		nodeMetricsData = <-nodeMetricsCh
	}

	select {
	case err := <-alertsErrCh:
		log.Fatal(err)
	default:
		alertsDataRaw = <-alertsCh
	}

	select {
	case err := <-k8sInfoErrCh:
		log.Fatal(err)
	default:
		k8sInfoData = <-k8sInfoCh
	}

	var gt65, gt80, lt65 int
	for _, nodeMetric := range nodeMetricsData {
		switch {
		case nodeMetric.MemoryUsePercentage >= 80.00:
			gt80++
		case (nodeMetric.MemoryUsePercentage >= 65.00) && (nodeMetric.MemoryUsePercentage < 80.00):
			gt65++
		case nodeMetric.MemoryUsePercentage < 65.00:
			lt65++
		}
	}

	var (
		loki, monitoring, cattleSystem, gita, kubeSystem, certManager, others int
	)
	if incidentsData.Total > 0 {
		for _, incident := range incidentsData.Msg {
			switch incident.Namespace {
			case "loki":
				loki += 1
			case "monitoring":
				monitoring += 1
			case "cattle-monitoring-system":
				monitoring += 1
			case "cattle-system":
				cattleSystem += 1
			case "gita":
				gita += 1
			case "kube-system":
				kubeSystem += 1
			case "cert-manager":
				certManager += 1
			default:
				others += 1
			}
		}
	}

	var alertsData []Alert
	for _, alert := range alertsDataRaw {
		alertsData = append(alertsData,
			Alert{
				AlertName:        alert.Name,
				AlertResource:    alert.Kind,
				AlertNamespace:   alert.Namespace,
				AlertDescription: alert.Msg,
				AlertDate:        alert.DataHora.Date.Format("02/01/06"),
			},
		)
	}

	return KubernetesReportData{
		Title:          cluster.Name,
		CreatedUtc:     time.Now(),
		NodesCount:     cluster.Node,
		IncidentsCount: incidentsData.Total,
		ProblemsCount:  problemData.Total,
		SecurityCount:  securityData.Total,

		// Kubernetes Version
		KubernetesVersion:         k8sInfoData.Version,
		EndOfSupport:              k8sInfoData.EndOfSupport,
		LatestRelease:             k8sInfoData.LastRelease,
		K8sVersionSupportColorMap: k8sInfoData.SupportStatus,

		// Métrics
		TotalCPU:             clusterMetricsData.TotalCPU,
		ColorMapCPU:          maps.ColorRuleResources(clusterMetricsData, "CPU"),
		TotalMemory:          clusterMetricsData.TotalMemory,
		ColorMapMemory:       maps.ColorRuleResources(clusterMetricsData, "MEM"),
		TotalPods:            clusterMetricsData.TotalPods,
		ColorMapPods:         maps.ColorRuleResources(clusterMetricsData, "POD"),
		CounterLessThan65:    lt65,
		CounterGreaterThan65: gt65,
		CounterGreaterThen80: gt80,

		// Alerts
		Alerts: alertsData,

		// Incidents
		CertManagerIncidentsColorMap: maps.ColorRuleIncident(certManager),
		GitaIncidentsColorMap:        maps.ColorRuleIncident(gita),
		LokiIncidentsColorMap:        maps.ColorRuleIncident(loki),
		MonitoringIncidentsColorMap:  maps.ColorRuleIncident(monitoring),
		RancherIncidentsColorMap:     maps.ColorRuleIncident(cattleSystem),
		OtherIncidentsColorMap:       maps.ColorRuleIncident(others),
	}
}
