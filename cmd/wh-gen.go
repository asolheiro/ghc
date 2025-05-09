package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/asolheiro/gita-healthcheck/internal/api-calls/auth"
	"github.com/asolheiro/gita-healthcheck/internal/api-calls/count"
	"github.com/asolheiro/gita-healthcheck/internal/webhooks"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func init() {
	whMessageCmd.Flags().StringVarP(&orgFilter, "org", "o", "", "Filter report to specific organization name")
}

var whMessageCmd = &cobra.Command{
	Use:   "wh-msg",
	Short: "test",
	Run: func(cmd *cobra.Command, args []string) {

		authResponse, err := auth.Authentication()
		if err != nil {
			log.Fatalf("failed to authenticate on Gita, err: %s", err)
		}

		userCount, err := count.GetUserCount(authResponse.AccessToken)
		if err != nil {
			log.Fatalf("failed to get orgs info on Gita, err: %s", err)
		}

		if orgFilter == "" {
			fmt.Printf("Please, enter a '-o/--org' value")
		} else {
			if err := godotenv.Load(); err != nil {
				log.Fatalf("failed to load .env, err: %s ", err)
			}

			webhookURL := os.Getenv("WEBHOOK_URL")
			if webhookURL == "" {
				log.Fatal("WEBHOOK_URL environment variable is not set")
			}

			var organizationCount count.Msg
			for _, orgCount := range userCount.Msg {
				if orgCount.Organization.Name == orgFilter {
					organizationCount = orgCount
					break
				}
			}

			for _, cluster := range organizationCount.Clusters {
				k8sReportData := webhooks.FindClustersInfo(authResponse, organizationCount, cluster)

				if err != nil {
					log.Fatalf("failed to generate string from markdown, err: %s", err)
				}

				if err := webhooks.SendKubernetesReportCard(webhookURL, k8sReportData); err != nil {
					fmt.Println("Erro:", err)
				} else {
					fmt.Println("Enviado com sucesso!")
				}
			}
		}

	},
}
