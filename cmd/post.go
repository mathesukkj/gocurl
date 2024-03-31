/*
Copyright © 2024 Matheus Kemuel kemuel.g7363@gmail.com
*/
package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// postCmd represents the post command
var postCmd = &cobra.Command{
	Use:   "post",
	Short: "Does post requests, and output the response to stdout",
	Long:  `Uses the golang http package to do post requests, and the fmt package to output the response as text to stdout.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("error: no url provided")
			os.Exit(1)
		}

		postRequest(args[0])
	},
}

func init() {
	rootCmd.AddCommand(postCmd)
}

func postRequest(url string) {
	for k, v := range Body {
		FormValues.Add(k, v)
	}

	r, err := http.NewRequest(http.MethodPost, url, strings.NewReader(FormValues.Encode()))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

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
