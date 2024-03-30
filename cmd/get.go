/*
Copyright © 2024 Matheus Kemuel kemuel.g7363@gmail.com
*/
package cmd

import (
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Does get requests, and output the response to stdout",
	Long:  `Uses the golang http package to do get requests, and the fmt package to output the response as text to stdout.`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

}

func getRequest(url string) {
	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	http.DefaultClient.Do(r)
}
