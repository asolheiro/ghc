package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/asolheiro/gita-healthcheck/internal"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

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
	  fmt.Println("Enter yout credentials:")
	  var email string
	  fmt.Print("\ne-mail: ")
	  fmt.Scan(&email)

	  fmt.Print("password: ")
	  bytePassword, err := term.ReadPassword(int(os.Stdin.Fd()))
	  if err != nil {
		log.Fatalf("failed to read password: %v", err)
	  }

	  password := string(bytePassword)
	  authResponse, err := internal.Login(email, password)
	  fmt.Println(authResponse, err)

	},
  }
  
  func Execute() {
	if err := rootCmd.Execute(); err != nil {
	  fmt.Println(err)
	  os.Exit(1)
	}
  }