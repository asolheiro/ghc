package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/asolheiro/gita-healthcheck/internal/api-calls/auth"
	"github.com/asolheiro/gita-healthcheck/internal/api-calls/count"
	"github.com/asolheiro/gita-healthcheck/internal/graphql"
	"github.com/spf13/cobra"
)


func init() {
	graphQlCmd.Flags().StringVarP(&orgFilter, "org", "o", "", "Filter report to specific organization name")
}

var graphQlCmd = &cobra.Command{
	Use:   "gql-gen",
	Short: "Generate a GraphQL request to post healthcheck directly inside a Wiki.Js",
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
					err := graphql.GraphQLPost(fileName)
					if err != nil {
						fmt.Printf("\nerror sending request, err: %s", err)
					}

					if err := convertReport(fileName); err != nil {
						fmt.Printf("\nerror converting reports, err: %s", err)
					}
					if err := os.Remove(fileName); err != nil {
						fmt.Printf("\nerror removing '%s', err: ", err)
					}
				}
			}
		} else {
			fmt.Printf("Please, enter a '-o/--org' value")
		}	
	},
}