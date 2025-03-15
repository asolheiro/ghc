package cmd

import (
	"fmt"

	"log"

	"github.com/asolheiro/gita-healthcheck/internal/alerts"
	"github.com/asolheiro/gita-healthcheck/internal/auth"
	"github.com/asolheiro/gita-healthcheck/internal/count"
	"github.com/asolheiro/gita-healthcheck/internal/incidents"
	"github.com/asolheiro/gita-healthcheck/internal/metrics"
	"github.com/asolheiro/gita-healthcheck/internal/problem"
	"github.com/asolheiro/gita-healthcheck/internal/security"
	terminalutils "github.com/asolheiro/gita-healthcheck/internal/terminal-utils"
	"github.com/charmbracelet/bubbles/table"
	"github.com/spf13/cobra"
)

//go:generate go run main.go

func init() {
	printHcCmd.Flags().StringVarP(&orgFilter, "org", "o", "", "Filter report to specific organization name")
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
		if orgFilter == "" {
			log.Fatalln("Please, insert a '-o/--org'")
		}
		authResponse, err := auth.Authentication()
		if err != nil {
			log.Fatalln(err)
		}
		userCount, err := count.GetUserCount(authResponse.AccessToken)
		if err != nil {
			log.Fatal(err)
		}
		
		_, org := getOrg(userCount.Msg)
		
		var rows []table.Row
		for _, cluster := range org.Clusters {
			res := getInfoWithGoRoutine(*authResponse, cluster.ClusterID)

			rows = append(rows, []string{
				cluster.Name,
				cluster.ClusterID,
				res.usageCPU,
				res.usageRAM,
				res.usagePod,
				res.alertsNum,
				res.incidentsNum,
				res.problemsNum,
				res.securitiesNum,
			})
		}

		fmt.Printf("\n %s (%s)\n", org.Organization.Name, org.Organization.ID)
		terminalutils.RunTables(rows)
	},
}

func getOrg(orgs []count.Msg) (int, count.Msg) {
	for i, org := range orgs {
		if orgFilter == org.Organization.Name {
			return i, org
		}
	}
	return 0, count.Msg{}
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
	done := make(chan struct{})
	go func() {
		cm, _ := metrics.GetClusterMetrics(auth.AccessToken, clusterId)
		res.usageCPU = fmt.Sprintf("%.2f%%", cm.CPUUsePercentage)
		res.usageRAM = fmt.Sprintf("%.2f%%", cm.MemoryUsePercentage)
		res.usagePod = fmt.Sprintf("%d/%d", cm.TotalPods, cm.TotalPodCapacity)
		close(done)
	}()
	<-done

	done = make(chan struct{})
	go func() {
		alerts, _ := alerts.GetAlerts(auth.AccessToken, clusterId)
		res.alertsNum = fmt.Sprintf("%d", len(alerts))
		close(done)
	}()
	<-done

	done = make(chan struct{})
	go func() {
		incidents, _ := incidents.GetIncidents(auth.AccessToken, clusterId)
		res.incidentsNum = fmt.Sprintf("%d", incidents.Total)
		close(done)
	}()
	<-done

	done = make(chan struct{})
	go func() {
		securities, _ := security.GetSecurity(auth.AccessToken, clusterId)
		res.securitiesNum = fmt.Sprintf("%d", securities.Total)
		close(done)
	}()
	<-done

	done = make(chan struct{})
	go func() {
		problems, _ := problem.GetProblems(auth.AccessToken, clusterId)
		res.problemsNum = fmt.Sprintf("%d", problems.Total)
		close(done)
	}()
	<-done

	return res
}
