package cmd

import (
	"fmt"
	"log"

	"github.com/asolheiro/gita-healthcheck/internal/api-calls/auth"
	"github.com/asolheiro/gita-healthcheck/internal/api-calls/count"
	tu "github.com/asolheiro/gita-healthcheck/internal/terminal-utils"
	"github.com/spf13/cobra"
)


func init() {
	generatePageCmd.Flags().StringVarP(&orgFilter, "org", "o", "", "Filter report to specific organization name")
}

var generatePageCmd = &cobra.Command{
	Use:   "pg-gen",
	Short: "Generate a TUI with the proger mardown file with a simple report of Gita's plataform",
	Run: func(cmd *cobra.Command, args []string) {
		authResponse, err := auth.Authentication()
		if err != nil {
			log.Fatal(err)
		}

		count, _ := count.GetUserCount(authResponse.AccessToken)

		if orgFilter != "" {
			for _, org := range count.Msg {
				if org.Organization.Name == orgFilter {
					fileName := generateOrgReport(org, authResponse)
					tu.PrintTea(fileName)
				}
			}
		} else {
			fmt.Printf("Please, enter a '-o/--org' value")
		}	
	},
}