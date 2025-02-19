package cmd

import (
	"log"

	"github.com/asolheiro/gita-healthcheck/internal/auth"
	"github.com/spf13/cobra"
)

type fileVars struct {
	auth auth.AuthResponse

}
var generateMdCmd = &cobra.Command{
	Use:   "generate-md",
	Short: "Generate a markdown file with a simple report of Gita's plataform",
	Long: `
Gita HealthCheck is a simple report generator of Gita's plataform that automatize
the manual work that analists do every morning.


It gives the user a complete markdown report of what he can see in the plataform
in the cost of a few lines in CLI
	`,
	Run: func(cmd *cobra.Command, args []string) {
		authResponse, err := auth.Authentication()
		if err != nil {
			log.Fatal(err)
		}
		
		fileVar 
	},
}