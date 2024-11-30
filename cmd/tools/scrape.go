package tools

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var scrapeCmd = &cobra.Command{
	Use:   "scrape <url>",
	Short: "Scrape a website and print the text",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("url is not specified")
		}

		fullURL := fmt.Sprintf("https://r.jina.ai/%s", args[0])

		ctx, cancel := context.WithTimeout(context.Background(), time.Minute*1)
		defer cancel()

		req, err := http.NewRequestWithContext(ctx, http.MethodGet, fullURL, nil)
		if err != nil {
			cmd.PrintErrf("failed to create request: %v", err)
			return nil
		}

		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			cmd.PrintErrf("call failed: %v", err)
			return nil
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			cmd.PrintErrf("Call failed with status code %v", resp.Status)
			return nil
		}

		bb, err := io.ReadAll(resp.Body)
		if err != nil {
			cmd.PrintErrf("Failed to read response body: %v", err)
		}

		fmt.Println(string(bb))

		return nil
	},
}
