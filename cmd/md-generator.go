package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/asolheiro/gita-healthcheck/internal/auth"
	"github.com/asolheiro/gita-healthcheck/internal/count"
	"github.com/asolheiro/gita-healthcheck/internal/maps"
	"github.com/asolheiro/gita-healthcheck/internal/md"
	"github.com/spf13/cobra"
)


	func init() {
		generateMdCmd.Flags().StringVarP(&orgFilter, "org", "o", "", "Filter report to specific organization name")
	}

var generateMdCmd = &cobra.Command{
	Use:   "generate-md",
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
					generateOrgReport(org, authResponse)
				}
			}
		} else {
			generateAllReports(count, authResponse)
		}	
	},
}



func generateOrgReport(org count.Msg, auth *auth.AuthResponse) string {
	fmt.Printf("\nGenerating report for organization: %s\n", org.Organization.Name)
	fmt.Printf("Ω Total clusters found: %d\n", len(org.Clusters))

	path := "reports"
	_ = os.Mkdir(path, os.ModePerm)

	safeOrgName := strings.ReplaceAll(org.Organization.Name, " ", "_")
	timeStr := time.Now().Format("02-01-2006") 
	mainFile := fmt.Sprintf(path+"/%s_%s.md", safeOrgName, timeStr)

	f1, err := os.OpenFile(mainFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("Error creating main file: %v", err)
	}
	defer f1.Close()

	orgHeader := fmt.Sprintf("# %s - %s\n\n", org.Organization.Name, maps.TimeNowToFormattedDate())
	if _, err := f1.WriteString(orgHeader); err != nil {
		log.Fatalf("Error writing org header: %v", err)
	}

	for i, cluster := range org.Clusters {
		
		var fileVars md.FileVars
		done := make(chan struct{})
		go func() {
			fileVars = md.FindInfo(auth, org, i, cluster)
			close(done)
		}()

		<-done

		tmpFile := fmt.Sprintf("%d-gita-report-%s.md", i+1, cluster.Name)
		
		done = make(chan struct{})
		go func() {
			md.GenerateFile(fileVars)
			close(done)
		}()
		
		<-done

		tmpContent, err := os.ReadFile(tmpFile)
		if err != nil {
			log.Fatalf("Error reading temp file: %v", err)
		}

		if _, err := f1.Seek(0, 2); err != nil {
			log.Fatalf("Error seeking file: %v", err)
		}

		if _, err := f1.Write(tmpContent); err != nil {
			log.Fatalf("Error writing cluster content: %v", err)
		}

		if err := os.Remove(tmpFile); err != nil {
			log.Printf("Warning: couldn't remove temp file %s: %v", tmpFile, err)
		}

		fmt.Printf("    Σ Finished processing cluster %d: %s\n", i+1, cluster.Name)
	}
	return mainFile
}

func generateAllReports(count count.CountResponse, authResponse *auth.AuthResponse) {
	for _, msgCount := range count.Msg {
		fmt.Printf("\nGenerating report for organization: %s\n", msgCount.Organization.Name)
		fmt.Printf("Ω Total clusters found: %d\n", len(msgCount.Clusters))
		
		_ = generateOrgReport(msgCount, authResponse)
	}
}