package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:           "gita-healthcheck",
		Short:         "Gita HealthCheck is a simple report generator of Gita's plataform",
		Long: `
Gita HealthCheck is a simple report generator of Gita's plataform that automatize
the manual work that analists do every morning.


It gives the user a complete markdown report of what he can see in the plataform
in the cost of a few lines in CLI
	`,
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	rootCmd.AddCommand(
		printHcCmd,
	)

	return rootCmd
}

func Execute() {
	if err := NewRootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}