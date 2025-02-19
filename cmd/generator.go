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
    markdown "github.com/nao1215/markdown"
)

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

        count, _ := count.GetUserCount(authResponse.AccessToken)
        
        for _, msgCount := range count.Msg {
            fmt.Printf("\ngenerating report to %s\n", msgCount.Organization.Name)
            
            mainFile := fmt.Sprintf("healthCheck-%s.md", msgCount.Organization.Name)
            
            f1, err := os.Create(mainFile)
            if err != nil {
                fmt.Printf("Error creating main file %s, err: %v\n", mainFile, err)
                return
            }
            
            m := markdown.NewMarkdown(f1)
            m.H1(msgCount.Organization.Name)
            m.Build()
            
            f1.Close()
            
            for i, cluster := range msgCount.Clusters {
                fmt.Println("- cluster: ", cluster.Name)
                
                incidents, _ := incidents.GetIncidents(authResponse.AccessToken, cluster.ClusterID)
                problem, _ := problem.GetProblems(authResponse.AccessToken, cluster.ClusterID)
                security, _ := security.GetSecurity(authResponse.AccessToken, cluster.ClusterID)
                clusterMetrics, _ := metrics.GetClusterMetrics(authResponse.AccessToken, cluster.ClusterID)
                nodeMetrics, _ := metrics.GetNodeMetrics(authResponse.AccessToken, cluster.ClusterID)
                alertsList, _ := alerts.GetAlerts(authResponse.AccessToken, cluster.ClusterID)
                
                tmpFile := fmt.Sprintf("%d-gita-report-%s.md", i+1, cluster.Name)
                
                md.GenerateFile(md.FileVars{
                    Index:          i + 1, // Ensure we're using 1-based indexing
                    Auth:           *authResponse,
                    OrgName:        msgCount.Organization.Name,
                    Cluster:        cluster,
                    Incidents:      incidents,
                    Problem:        problem,
                    Security:       security,
                    ClusterMetrics: clusterMetrics,
                    NodeMetrics:    nodeMetrics,
                    AlertsList:     alertsList,
                })
                
                tmpContent, err := os.ReadFile(tmpFile)
                if err != nil {
                    fmt.Printf("Error reading temporary file %s, err: %v\n", tmpFile, err)
                    return
                }
                
                f1, err = os.OpenFile(mainFile, os.O_APPEND|os.O_WRONLY, 0644)
                if err != nil {
                    fmt.Printf("Error opening main file %s for appending, err: %v\n", mainFile, err)
                    return
                }
                
                f1.WriteString("\n")
                
                if _, err := f1.Write(tmpContent); err != nil {
                    fmt.Printf("Error appending content to main file, err: %v\n", err)
                    f1.Close()
                    return
                }
                
                f1.Close()
                
                if err := os.Remove(tmpFile); err != nil {
                    fmt.Printf("Error removing temporary file %s, err: %v\n", tmpFile, err)
                }
            }
        }
    },
}