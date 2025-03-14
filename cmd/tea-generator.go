package cmd

import (
	"fmt"
	"log"

	"github.com/asolheiro/gita-healthcheck/internal/auth"
	"github.com/asolheiro/gita-healthcheck/internal/count"
	tea "github.com/asolheiro/gita-healthcheck/internal/teaOutput"
	"github.com/spf13/cobra"
)


func init() {
	generateTeaCmd.Flags().StringVarP(&orgFilter, "org", "o", "", "Filter report to specific organization name")
}

var generateTeaCmd = &cobra.Command{
	Use:   "generate-tea",
	Short: "Generate a markdown file with a simple report of Gita's plataform",
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
					tea.PrintTea(fileName)
				}
			}
		} else {
			fmt.Printf("Please, enter a '-o/--org' value")
		}	
	},
}