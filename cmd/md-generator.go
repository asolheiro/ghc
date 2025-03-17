package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/asolheiro/gita-healthcheck/internal/api-calls/auth"
	"github.com/asolheiro/gita-healthcheck/internal/api-calls/count"
	"github.com/asolheiro/gita-healthcheck/internal/maps"
	"github.com/asolheiro/gita-healthcheck/internal/md"
	"github.com/spf13/cobra"
)

func init() {
	generateMdCmd.Flags().StringVarP(&orgFilter, "org", "o", "", "Filter report to specific organization name")
}

var generateMdCmd = &cobra.Command{
	Use:   "gen-md",
	Short: "Generate a markdown file with a simple report of Gita's platform",
	Run: func(cmd *cobra.Command, args []string) {
		authResponse, err := auth.Authentication()
		if err != nil {
			log.Fatal(err)
		}

		count, err := count.GetUserCount(authResponse.AccessToken)
		if err != nil {
			log.Fatal("Error getting user count:", err)
		}

		if orgFilter != "" {
			for _, org := range count.Msg {
				if org.Organization.Name == orgFilter {
					generateOrgReport(org, authResponse)
					break
				}
			}
		} else {
			generateAllReports(count, authResponse)
		}
	},
}

func generateOrgReport(org count.Msg, auth *auth.AuthResponse) string {
	fmt.Printf("\nGenerating report for organization: %s\n", org.Organization.Name)
	fmt.Printf("╰─ Total clusters found: %d\n", len(org.Clusters))

	path := "reports"
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		log.Fatalf("Error creating reports directory: %v", err)
	}

	safeOrgName := strings.ReplaceAll(org.Organization.Name, " ", "_")
	timeStr := time.Now().Format("02-01-2006")
	mainFile := filepath.Join(path, fmt.Sprintf("%s_%s.md", safeOrgName, timeStr))

	f1, err := os.Create(mainFile)
	if err != nil {
		log.Fatalf("Error creating main file: %v", err)
	}
	defer f1.Close()

	orgHeader := fmt.Sprintf("# %s - %s\n\n", org.Organization.Name, maps.TimeNowToFormattedDate())
	if _, err := f1.WriteString(orgHeader); err != nil {
		log.Fatalf("Error writing org header: %v", err)
	}

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 5)
	
	fileMutex := sync.Mutex{}
	for i, cluster := range org.Clusters {
		wg.Add(1)
		semaphore <- struct{}{}

	var allFileVars []md.FileVars
	var allFileVarsMutex sync.Mutex

	go func(index int, clusterInfo count.Cluster) {
		defer wg.Done()
		defer func() { <-semaphore }()
		
		fileVars := md.FindInfo(auth, org, index, clusterInfo)
		
		allFileVarsMutex.Lock()
		allFileVars = append(allFileVars, fileVars)
		allFileVarsMutex.Unlock()
	}(i, cluster)
	wg.Wait()

	sort.Slice(allFileVars, func(i, j int) bool {
		return allFileVars[i].Index < allFileVars[j].Index
	})

	for _, fileVars := range allFileVars {
		md.GenerateFile(fileVars)
		tmpFile := fmt.Sprintf("%d-gita-report-%s.md", fileVars.Index, fileVars.Cluster.Name)
		tmpContent, err := os.ReadFile(tmpFile)
		if err != nil {
			log.Printf("Error reading temp file %s: %v", tmpFile, err)
			continue
		}
		
		fileMutex.Lock()
		if _, err := f1.Write(tmpContent); err != nil {
			log.Printf("Error writing cluster content: %v", err)
		}
		fileMutex.Unlock()
		
		if err := os.Remove(tmpFile); err != nil {
			log.Printf("Warning: couldn't remove temp file %s: %v", tmpFile, err)
		}
	}
	}

	wg.Wait()
	return mainFile
}

func generateAllReports(countRes count.CountResponse, authResponse *auth.AuthResponse) {
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 3)

	for _, msgCount := range countRes.Msg {
		wg.Add(1)
		semaphore <- struct{}{}
		go func(org count.Msg) {
			defer wg.Done()
			defer func() { <-semaphore }()

			_ = generateOrgReport(org, authResponse)
		}(msgCount)
	}

	wg.Wait()
}