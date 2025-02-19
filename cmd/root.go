package cmd

import (
	"fmt"

	"log"
	"os"

	"github.com/asolheiro/gita-healthcheck/internal/auth"
	"github.com/asolheiro/gita-healthcheck/internal/count"
	"github.com/asolheiro/gita-healthcheck/internal/metrics"
	"github.com/spf13/cobra"
)
const ovhCluster = "e2d718b6919aab2b2cb459520e48f278"

var rootCmd = &cobra.Command{
	Use:   "gita-healthcheck",
	Short: "Gita HealthCheck is a simple report generator of Gita's plataform",
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
			fmt.Printf("  %v. %s (%s)\n", i + 1, msg.Organization.Name, msg.Organization.ID)
			for j, cluster := range msg.Clusters {
				fmt.Printf("   %d.%d. %s (%s)\n", i+1, j+1, cluster.Name, cluster.ClusterID)
					cm, _ := metrics.GetMetrics(authResponse.AccessToken, cluster.ClusterID)
					fmt.Printf("     Usage - CPU: %.2f%%, Mem: %.2f%%, Pods: %d/%d\n", cm.CPUUsePercentage, cm.MemoryUsePercentage, cm.TotalPods,cm.TotalPodCapacity)
			}
			fmt.Println()
		}
		metr, _ := metrics.GetMetrics(authResponse.AccessToken, ovhCluster)

		fmt.Println(
			// cluster.GetNodes(authResponse.AccessToken, ovhCluster),
			// internal.TestEndpoint(authResponse.AccessToken, ovhCluster),
			metr,
		)
	},
  }
  
  func Execute() {
	if err := rootCmd.Execute(); err != nil {
	  fmt.Println(err)
	  os.Exit(1)
	}
  }