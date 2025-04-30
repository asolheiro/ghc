package cmd

import (
	"fmt"

	"github.com/asolheiro/gita-healthcheck/internal/webhooks"
	"github.com/spf13/cobra"
)

const WEBHOOK_URL = "https://chat.jacexperts.io/hooks/rt1ctb8ox7y1dekjmax1xekyqw"


func init() {
	whMessageCmd.Flags().StringVarP(&orgFilter, "org", "o", "", "Filter report to specific organization name")
}

var whMessageCmd = &cobra.Command{
	Use:   "wh-msg",
	Short: "test",
	Run: func(cmd *cobra.Command, args []string) {
		
			message := "Here's a PDF document for your review!"
			pdfPath := "../reports/CHESF_14-04-2025.pdf" // Replace with the path to your PDF file
			
			// For Mattermost testing, use this method
			fmt.Println("Sending PDF to Mattermost...")
			err := webhooks.SendMessageWithPDF(WEBHOOK_URL, message, pdfPath)
			if err != nil {
				fmt.Printf("Error sending to Mattermost: %v\n", err)
			} else {
				fmt.Println("Message with PDF sent successfully to Mattermost!")
			}
			
			// For Microsoft Teams (when you're ready to switch)
			// Note: You may need to adjust this based on Teams' specific requirements
			/*
			fmt.Println("Sending PDF to Microsoft Teams...")
			err = sendPDFMultipart(webhookURL, message, pdfPath)
			if err != nil {
				fmt.Printf("Error sending to Microsoft Teams: %v\n", err)
			} else {
				fmt.Println("Message with PDF sent successfully to Microsoft Teams!")
			}
			*/
	},
	
}
