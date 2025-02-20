package cmd

import (
	"fmt"
	"log"
	"os"


	"github.com/asolheiro/gita-healthcheck/internal/alerts"
	"github.com/asolheiro/gita-healthcheck/internal/auth"
	"github.com/asolheiro/gita-healthcheck/internal/count"
	"github.com/asolheiro/gita-healthcheck/internal/incidents"
	"github.com/asolheiro/gita-healthcheck/internal/md"
	"github.com/asolheiro/gita-healthcheck/internal/metrics"
	"github.com/asolheiro/gita-healthcheck/internal/problem"
	"github.com/asolheiro/gita-healthcheck/internal/security"
	"github.com/spf13/cobra"
)

var generateMdCmd = &cobra.Command{
    Use:   "generate-md",
    Short: "Generate a markdown file with a simple report of Gita's plataform",
    Run: func(cmd *cobra.Command, args []string) {
        authResponse, err := auth.Authentication()
        if err != nil {
            log.Fatal(err)
        }

        count, _ := count.GetUserCount(authResponse.AccessToken)
        
        for _, msgCount := range count.Msg {
            fmt.Printf("\nGenerating report for organization: %s\n", msgCount.Organization.Name)
            fmt.Printf("Total clusters found: %d\n", len(msgCount.Clusters))
            
            // Print all clusters first to verify the data
            for i, cluster := range msgCount.Clusters {
                fmt.Printf("Cluster %d: %s (ID: %s)\n", i+1, cluster.Name, cluster.ClusterID)
            }
            
            mainFile := fmt.Sprintf("healthCheck-%s.md", msgCount.Organization.Name)
            
            // Create the file with write permissions
            f1, err := os.OpenFile(mainFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
            if err != nil {
                log.Fatalf("Error creating main file: %v", err)
            }
            defer f1.Close()

            // Write the organization header first
            orgHeader := fmt.Sprintf("# %s\n\n", msgCount.Organization.Name)
            if _, err := f1.WriteString(orgHeader); err != nil {
                log.Fatalf("Error writing org header: %v", err)
            }
            
            // Process each cluster
            for i, cluster := range msgCount.Clusters {
                fmt.Printf("Processing cluster %d/%d: %s\n", i+1, len(msgCount.Clusters), cluster.Name)
                
                incidents, err := incidents.GetIncidents(authResponse.AccessToken, cluster.ClusterID)
                if err != nil {
                    log.Fatal(err)
                }
                problem, err := problem.GetProblems(authResponse.AccessToken, cluster.ClusterID)
                if err != nil {
                    log.Fatal(err)
                }
                security, err := security.GetSecurity(authResponse.AccessToken, cluster.ClusterID)
                if err != nil {
                    log.Fatal(err)
                }
                clusterMetrics, err := metrics.GetClusterMetrics(authResponse.AccessToken, cluster.ClusterID)
                if err != nil {
                    log.Fatal(err)
                }
                nodeMetrics, err := metrics.GetNodeMetrics(authResponse.AccessToken, cluster.ClusterID)
                if err != nil {
                    log.Fatal(err)
                }
                alertsList, err := alerts.GetAlerts(authResponse.AccessToken, cluster.ClusterID)
                if err != nil {
                    log.Fatal(err)
                }
                
                fileVars := md.FileVars{
                    Index:          i + 1,  // Make sure index starts at 1
                    Auth:           *authResponse,
                    OrgName:        msgCount.Organization.Name,
                    Cluster:        cluster,
                    Incidents:      incidents,
                    Problem:        problem,
                    Security:       security,
                    ClusterMetrics: clusterMetrics,
                    NodeMetrics:    nodeMetrics,
                    AlertsList:     alertsList,
                }

                // Generate content for this cluster
                tmpFile := fmt.Sprintf("%d-gita-report-%s.md", i+1, cluster.Name)
                md.GenerateFile(fileVars)
                
                // Read the generated content
                tmpContent, err := os.ReadFile(tmpFile)
                if err != nil {
                    log.Fatalf("Error reading temp file: %v", err)
                }
                
                // Seek to end of file before writing new content
                if _, err := f1.Seek(0, 2); err != nil {
                    log.Fatalf("Error seeking file: %v", err)
                }
                
                // Write the cluster content
                if _, err := f1.Write(tmpContent); err != nil {
                    log.Fatalf("Error writing cluster content: %v", err)
                }
                
                // Remove temp file
                if err := os.Remove(tmpFile); err != nil {
                    log.Printf("Warning: couldn't remove temp file %s: %v", tmpFile, err)
                }
                
                fmt.Printf("Finished processing cluster %d: %s\n", i+1, cluster.Name)
            }
        }
    },
}