package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	// Setup chrome options for non-headless mode (for debugging)
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false), // Disable headless mode
		chromedp.Flag("disable-gpu", false),
		chromedp.Flag("start-maximized", true), // Start with maximized window
	)

	// Create a context with the above options
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// Create browser context
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// Set timeout for the flow
	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	// Run the Salesforce login and flow test
	err := chromedp.Run(ctx, runSalesforceFlowTest(ctx))
	if err != nil {
		log.Fatalf("Failed to run the flow test: %v", err)
	}
}

func runSalesforceFlowTest(ctx context.Context) chromedp.Tasks {
	return chromedp.Tasks{
		// Navigate to Salesforce login page
		chromedp.Navigate("https://login.salesforce.com"),

		// Wait for the login page to load
		chromedp.WaitVisible(`#username`, chromedp.ByID),

		// Input username and password (use your credentials)
		chromedp.SendKeys(`#username`, "your_salesforce_username"),
		chromedp.SendKeys(`#password`, "your_salesforce_password"),

		// Click the login button
		chromedp.Click(`#Login`, chromedp.ByID),

		// Wait for the Salesforce dashboard to load (increase this wait if needed)
		chromedp.Sleep(5 * time.Second),

		// Navigate to the specific flow URL (if you have a direct link)
		chromedp.Navigate("https://your_salesforce_instance_url/flowScreenPage"),

		// Wait for a flow element to appear (replace with actual flow element)
		chromedp.WaitVisible(`#flowElementID`, chromedp.ByID),

		// Example of interacting with a button on the flow screen
		chromedp.Click(`#flowButtonID`, chromedp.ByID),
	}
}
