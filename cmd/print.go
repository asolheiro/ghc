package cmd

import (
	"fmt"
	"sync"

	"log"

	"github.com/asolheiro/gita-healthcheck/internal/alerts"
	"github.com/asolheiro/gita-healthcheck/internal/auth"
	"github.com/asolheiro/gita-healthcheck/internal/count"
	"github.com/asolheiro/gita-healthcheck/internal/incidents"
	"github.com/asolheiro/gita-healthcheck/internal/metrics"
	"github.com/asolheiro/gita-healthcheck/internal/problem"
	"github.com/asolheiro/gita-healthcheck/internal/security"
	terminalutils "github.com/asolheiro/gita-healthcheck/internal/terminal-utils"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

//go:generate go run main.go

func init() {
	printHcCmd.Flags().StringVarP(
		&orgFilter, "org", "o", "", "Filter report to specific organization name",
	)
	printHcCmd.Flags().IntVarP(
		&concurrencyLimit, "concurrency", "c", 2, "Maximum number of concurrent cluster operations",
	)
}

var printHcCmd = &cobra.Command{
	Use:   "print-hc",
	Short: "Print on terminal is a simple report of Gita's plataform",
	Long: `
Gita HealthCheck is a simple report generator of Gita's plataform that automatize
the manual work that analists do every morning.


It gives the user a complete markdown report of what he can see in the plataform
in the cost of a few lines in CLI
	`,
	Run: func(cmd *cobra.Command, args []string) {
		authResponse, err := auth.Authentication()
		if err != nil {
			log.Fatalln(err)
		}
		userCount, err := count.GetUserCount(authResponse.AccessToken)
		if err != nil {
			log.Fatal(err)
		}
		
		if orgFilter == "" {
			processAllOrgs(userCount.Msg, *authResponse)
		} else {
			org := getOrg(userCount.Msg)
			if org.Organization.Name == "" {
				log.Fatalf("Organization '%s' not found", orgFilter)
			}
			processOrg(org, *authResponse)
		}
	},
}

func getOrg(orgs []count.Msg) count.Msg {
	for _, org := range orgs {
		if orgFilter == org.Organization.Name {
			return org
		}
	}
	return count.Msg{}
}

type hcVars struct {
	usageCPU      string
	usageRAM      string
	usagePod      string
	alertsNum     string
	incidentsNum  string
	problemsNum   string
	securitiesNum string
}

func getInfoWithGoRoutine(auth auth.AuthResponse, clusterId string) hcVars {
	var res hcVars
	var wg sync.WaitGroup
	var mutex sync.Mutex

	wg.Add(5)

	go func() {
		defer wg.Done()
		cm, err := metrics.GetClusterMetrics(auth.AccessToken, clusterId)
		if err != nil {
			log.Printf("Error getting cluster metrics: %v", err)
			return
		}
		
		mutex.Lock()
		res.usageCPU = fmt.Sprintf("%.2f%%", cm.CPUUsePercentage)
		res.usageRAM = fmt.Sprintf("%.2f%%", cm.MemoryUsePercentage)
		res.usagePod = fmt.Sprintf("%d/%d", cm.TotalPods, cm.TotalPodCapacity)
		mutex.Unlock()
	}()
	
	go func() {
		defer wg.Done()
		alerts, err := alerts.GetAlerts(auth.AccessToken, clusterId)
		if err != nil {
			log.Printf("Error getting alerts: %v", err)
			return
		}
		
		mutex.Lock()
		res.alertsNum = fmt.Sprintf("%d", len(alerts))
		mutex.Unlock()
	}()
	
	go func() {
		defer wg.Done()
		incidents, err := incidents.GetIncidents(auth.AccessToken, clusterId)
		if err != nil {
			log.Printf("Error getting incidents: %v", err)
			return
		}
		
		mutex.Lock()
		res.incidentsNum = fmt.Sprintf("%d", incidents.Total)
		mutex.Unlock()
	}()
	
	go func() {
		defer wg.Done()
		securities, err := security.GetSecurity(auth.AccessToken, clusterId)
		if err != nil {
			log.Printf("Error getting securities: %v", err)
			return
		}
		
		mutex.Lock()
		res.securitiesNum = fmt.Sprintf("%d", securities.Total)
		mutex.Unlock()
	}()
	
	go func() {
		defer wg.Done()
		problems, err := problem.GetProblems(auth.AccessToken, clusterId)
		if err != nil {
			log.Printf("Error getting problems: %v", err)
			return
		}
		
		mutex.Lock()
		res.problemsNum = fmt.Sprintf("%d", problems.Total)
		mutex.Unlock()
	}()
	
	wg.Wait()
	return res
}

func processOrg(org count.Msg, authResponse auth.AuthResponse) {
	fmt.Printf("\n %s (%s)\n", org.Organization.Name, org.Organization.ID)
	
	semaphore := make(chan struct{}, concurrencyLimit)
	
	var wg sync.WaitGroup
	var mutex sync.Mutex
	rows := make([]table.Row, 0, len(org.Clusters))
	
	for _, cluster := range org.Clusters {
		wg.Add(1)
		go func(c count.Cluster) {
			defer wg.Done()
			
			// Acquire a slot in the semaphore
			semaphore <- struct{}{}
			defer func() { <-semaphore }()
			
			fmt.Printf("Processing cluster: %s\n", c.Name)
			res := getInfoWithGoRoutine(authResponse, c.ClusterID)
			
			row := []any{
				c.Name,
				c.ClusterID,
				res.usageCPU,
				res.usageRAM,
				res.usagePod,
				res.alertsNum,
				res.incidentsNum,
				res.problemsNum,
				res.securitiesNum,
			}
			
			mutex.Lock()
			rows = append(rows, row)
			mutex.Unlock()
		}(cluster)
	}
	
	wg.Wait()
	terminalutils.PrettyTables(rows)
}

func processAllOrgs(orgs []count.Msg, authResponse auth.AuthResponse) {
	for _, org := range orgs {
		processOrg(org, authResponse)
	}
}