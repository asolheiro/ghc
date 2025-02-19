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
	"github.com/spf13/cobra"
)
 
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

		fmt.Printf("\n%d organizations found:\n", len(userCount.Msg))
		for i, msg := range userCount.Msg {
			fmt.Printf("\n  %v. %s (%s)\n", i + 1, msg.Organization.Name, msg.Organization.ID)
			for j, cluster := range msg.Clusters {
				fmt.Printf("   %d.%d. %s (%s)\n", i+1, j+1, cluster.Name, cluster.ClusterID)
				
				cm, _ := metrics.GetClusterMetrics(authResponse.AccessToken, cluster.ClusterID)
				fmt.Println("    Usage:")
				fmt.Printf("    - CPU: %.2f%%\n", cm.CPUUsePercentage)
				fmt.Printf("    - Mem: %.2f%%\n", cm.MemoryUsePercentage,)
				fmt.Printf("    - Pods: %d/%d\n", cm.TotalPods,cm.TotalPodCapacity)
				
				alerts, _ := alerts.GetAlerts(authResponse.AccessToken, cluster.ClusterID)
				fmt.Println("\n    Alerts: ", len(alerts))
				if len(alerts) > 0 {
					for _, alert := range alerts {
						fmt.Printf("      - %s.%s - %s\n", alert.Kind, alert.Namespace, alert.Msg)
					}
				}
				incidents, err := incidents.GetIncidents(authResponse.AccessToken, cluster.ClusterID)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("\n    Incidents: ", len(incidents))
				if len(incidents) > 0 {
					var (loki, monitoring, cattleSystem, gita, kubeSystem, certManager, others int)
					for _, incident := range incidents {
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
							gita+=1
						case "kube-system":
							kubeSystem += 1
						case "cert-manager":
							certManager += 1
						default:
							others += 1
						}
					}
					fmt.Println("      - loki: ", loki)
					fmt.Println("      - monitoring: ", monitoring)
					fmt.Println("      - cattle-system: ", cattleSystem)
					fmt.Println("      - gita: ", gita)
					fmt.Println("      - kube-system: ", kubeSystem)
					fmt.Println("      - cert-manager: ", certManager)
					fmt.Println("      - others: ", others)
				}
				
				securities, _ := security.GetSecurity(authResponse.AccessToken, cluster.ClusterID)
				fmt.Println("\n    Security alerts: ", securities.Total)

				problems, _ := problem.GetProblems(authResponse.AccessToken, cluster.ClusterID)
				fmt.Println("\n    Problem alerts: ", problems.Total)
			}
		}
	},
}
