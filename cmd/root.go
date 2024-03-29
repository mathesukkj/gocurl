/*
Copyright © 2024 Matheus Kemuel kemuel.g7363@gmail.com
*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gocurl",
	Short: "A golang version of curl",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("error: no url provided")
		}

		urlStr := args[0]
		res, err := http.Get(urlStr)
		if err != nil {
			log.Fatal(err)
		}

		r, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(r))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
