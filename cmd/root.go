/*
Copyright © 2024 Matheus Kemuel kemuel.g7363@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "gocurl",
		Short: "A golang version of curl",
		// Run: func(cmd *cobra.Command, args []string) {},
	}

	headersString string
	Headers       = make(map[string]string, 0)
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

	if headersString != "" {
		ParseHeaders()
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().
		StringVarP(&headersString, "headers", "H", "", "Pass custom headers to server")
}

func ParseHeaders() []map[string]string {
	xs := strings.Split(headersString, ";")
	for _, v := range xs {
		headerAndValue := strings.Split(v, ":")
		if len(headerAndValue) < 2 {
			fmt.Println("error: headers in wrong format, use key:value")
			os.Exit(1)
		}

		header := strings.Trim(headerAndValue[0], " ")
		value := strings.Trim(headerAndValue[1], " ")
		Headers[header] = value
	}

	fmt.Println(Headers)
	return nil

}
