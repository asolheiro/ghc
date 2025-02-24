package cmd

import (
	"fmt"
	"log"
	"os"
	"time"


	"github.com/asolheiro/gita-healthcheck/internal/auth"
	"github.com/asolheiro/gita-healthcheck/internal/count"
	"github.com/asolheiro/gita-healthcheck/internal/md"
	"github.com/spf13/cobra"
)

var generateMdCmd = &cobra.Command{
	Use:   "generate-md",
	Short: "Generate a markdown file with a simple report of Gita's plataform",
	Run: func(cmd *cobra.Command, args []string) {
		t1 := time.Now()
		authResponse, err := auth.Authentication()
		if err != nil {
			log.Fatal(err)
		}

		count, _ := count.GetUserCount(authResponse.AccessToken)

		for _, msgCount := range count.Msg {
			fmt.Printf("\nGenerating report for organization: %s\n", msgCount.Organization.Name)
			fmt.Printf("Ω Total clusters found: %d\n", len(msgCount.Clusters))

			for i, cluster := range msgCount.Clusters {
				fmt.Printf("  Δ Cluster %d: %s (ID: %s)\n", i+1, cluster.Name, cluster.ClusterID)
			}

			path := "reports"
			_ = os.Mkdir(path, os.ModePerm)
			mainFile := fmt.Sprintf(path+"/%s.md", msgCount.Organization.Name)

			f1, err := os.OpenFile(mainFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
			if err != nil {
				log.Fatalf("Error creating main file: %v", err)
			}
			defer f1.Close()

			orgHeader := fmt.Sprintf("# %s - %s\n\n", msgCount.Organization.Name, timeNowToFormattedDate())
			if _, err := f1.WriteString(orgHeader); err != nil {
				log.Fatalf("Error writing org header: %v", err)
			}

			for i, cluster := range msgCount.Clusters {
				fileVars := md.FindInfo(authResponse, msgCount, i, cluster)

				tmpFile := fmt.Sprintf("%d-gita-report-%s.md", i+1, cluster.Name)
				md.GenerateFile(fileVars)

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
		}
		t2 := time.Now()
		delta := t2.Sub(t1)
		fmt.Println(delta)
	},
}

func timeNowToFormattedDate() string {
	monthMap := map[time.Month]string{
		time.January:   "Janeiro",
		time.February:  "Fevereiro",
		time.March:     "Março",
		time.April:     "Abril",
		time.May:       "Maio",
		time.June:      "Junho",
		time.July:      "Julho",
		time.August:    "Agosto",
		time.September: "Setembro",
		time.October:   "Outubro",
		time.November:  "Novembro",
		time.December:  "Dezembro",
	}

	now := time.Now()
	return fmt.Sprintf("%d de %s de %d", now.Day(), monthMap[now.Month()], now.Year())
}
