package md

import (
	"fmt"
	"log"
	"sync"

	"github.com/asolheiro/gita-healthcheck/internal/alerts"
	"github.com/asolheiro/gita-healthcheck/internal/auth"
	"github.com/asolheiro/gita-healthcheck/internal/count"
	"github.com/asolheiro/gita-healthcheck/internal/incidents"
	"github.com/asolheiro/gita-healthcheck/internal/metrics"
	"github.com/asolheiro/gita-healthcheck/internal/problem"
	"github.com/asolheiro/gita-healthcheck/internal/security"
	"github.com/asolheiro/gita-healthcheck/internal/version"
)

func FindInfo(auth *auth.AuthResponse, msgCount count.Msg, i int, cluster count.Cluster) FileVars {
	fmt.Printf("    σ Processing cluster %d/%d: %s\n", i+1, len(msgCount.Clusters), cluster.Name)

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
	var alertsData []alerts.Msg
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
		alertsData = <-alertsCh
	}

	select {
	case err := <-k8sInfoErrCh:
		log.Fatal(err)
	default:
		k8sInfoData = <-k8sInfoCh
	}

	return FileVars{
		Index:          i + 1,
		Auth:           *auth,
		OrgName:        msgCount.Organization.Name,
		Cluster:        cluster,
		Incidents:      incidentsData,
		Problem:        problemData,
		Security:       securityData,
		ClusterMetrics: clusterMetricsData,
		NodeMetrics:    nodeMetricsData,
		AlertsList:     alertsData,
		KubernetesInfo: k8sInfoData,
	}
}