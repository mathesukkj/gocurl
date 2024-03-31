/*
Copyright © 2024 Matheus Kemuel kemuel.g7363@gmail.com
*/
package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Does get requests, and output the response to stdout",
	Long:  `Uses the golang http package to do get requests, and the fmt package to output the response as text to stdout.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("error: no url provided")
			os.Exit(1)
		}

		getRequest(args[0])
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}

func getRequest(url string) {
	r, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	AddFlagsToRequest(r)

	response, err := http.DefaultClient.Do(r)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Print(string(data))
}
