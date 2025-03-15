package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/asolheiro/gita-healthcheck/internal/api-calls/auth"
	"github.com/asolheiro/gita-healthcheck/internal/api-calls/count"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/russross/blackfriday/v2"
	"github.com/spf13/cobra"
)

func init() {
	generatePdfCmd.Flags().StringVarP(&orgFilter, "org", "o", "", "Filter report to specific organization name")
}

var generatePdfCmd = &cobra.Command {
		Use:   "gen-pdf",
		Short: "Generate PDF from markdown files",
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

						mdFile, err := findMarkdownFiles("./reports")
							if err != nil {
								fmt.Printf("Error finding markdown files: %v\n", err)
								return
							}
							pdfFile := strings.TrimSuffix(mdFile[0], ".md") + ".pdf"
						if err := convertMarkdownToPDF(mdFile[0], pdfFile); err != nil {
							fmt.Printf("Error converting %s: %v\n", mdFile, err)
						} else {
							fmt.Printf("Successfully converted %s to %s\n", mdFile[0], pdfFile)
						}
					}
				}
			} else {
				generateAllReports(count, authResponse)
				mdFiles, err := findMarkdownFiles("./reports")
				if err != nil {
					fmt.Printf("Error finding markdown files: %v\n", err)
					return
				}
				for _, mdFile := range mdFiles {
					pdfFile := strings.TrimSuffix(mdFile, ".md") + ".pdf"
					fmt.Printf("Converting %s to %s\n", mdFile, pdfFile)
	
					if err := convertMarkdownToPDF(mdFile, pdfFile); err != nil {
						fmt.Printf("Error converting %s: %v\n", mdFile, err)
					} else {
						fmt.Printf("Successfully converted %s to %s\n", mdFile, pdfFile)
					}
				}
			}
		},
	}


func findMarkdownFiles(dir string) ([]string, error) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil, fmt.Errorf("directory for organization %s does not exist", dir)
	}

	var mdFiles []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			mdFiles = append(mdFiles, path)
		}
		return nil
	})
	return mdFiles, err
}

func convertMarkdownToPDF(mdFile, pdfFile string) error {
    mdContent, err := os.ReadFile(mdFile)
    if err != nil {
        return fmt.Errorf("failed to read markdown file: %w", err)
    }

    html := blackfriday.Run(mdContent)

    htmlContent := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <style>
        body { 
            font-family: Arial, sans-serif; 
            line-height: 1.6; 
            margin: 2em;
            max-width: 800px;
        }
        pre { 
            background-color: #f4f4f4; 
            padding: 1em; 
            border-radius: 4px; 
            overflow: auto;
        }
        code { font-family: monospace; }
        h1, h2, h3 { color: #333; }
        a { color: #0366d6; }
        table { border-collapse: collapse; width: 100%%; }
        th, td { border: 1px solid #ddd; padding: 8px; }
        th { background-color: #f4f4f4; }
    </style>
</head>
<body>
    %s
</body>
</html>`, string(html))

    tempFile, err := os.CreateTemp("", "markdown-*.html")
    if err != nil {
        return fmt.Errorf("failed to create temp file: %w", err)
    }
    tempFilePath := tempFile.Name()
    defer os.Remove(tempFilePath)

    if _, err := tempFile.WriteString(htmlContent); err != nil {
        tempFile.Close()
        return fmt.Errorf("failed to write HTML to temp file: %w", err)
    }
    tempFile.Close()

    opts := append(chromedp.DefaultExecAllocatorOptions[:],
        chromedp.Flag("headless", true),
        chromedp.Flag("disable-gpu", true),
        chromedp.Flag("no-sandbox", true),
        chromedp.Flag("disable-software-rasterizer", true),
        chromedp.Flag("disable-dev-shm-usage", true),
    )

    allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
    defer cancel()

    ctx, cancel := chromedp.NewContext(allocCtx)
    defer cancel()

    ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
    defer cancel()

    fileURL := fmt.Sprintf("file://%s", tempFilePath)

    var buf []byte
    if err := chromedp.Run(ctx, chromedp.Tasks{
        chromedp.Navigate(fileURL),
        chromedp.Sleep(1 * time.Second),
        chromedp.ActionFunc(func(ctx context.Context) error {
            var err error
            buf, _, err = page.PrintToPDF().
                WithPrintBackground(true).
                WithMarginTop(0.4).
                WithMarginBottom(0.4).
                WithMarginLeft(0.4).
                WithMarginRight(0.4).
                WithPaperWidth(8.27). 
                WithPaperHeight(11.69).
                Do(ctx)
            return err
        }),
    }); err != nil {
        return fmt.Errorf("failed to generate PDF: %w", err)
    }

    if len(buf) == 0 {
        return fmt.Errorf("generated PDF is empty")
    }

    if err := os.WriteFile(pdfFile, buf, 0644); err != nil {
        return fmt.Errorf("failed to write PDF file: %w", err)
    }

    return nil
}