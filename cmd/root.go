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

	headersString  string
	bodyString     string
	outputFilename string
	cookiesString  string
	isVerbose      bool
	includeHeaders bool

	Headers = make(map[string]string, 0)
	Body    = make(map[string]string, 0)
	Cookies = make(map[string]string, 0)
)

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

	if headersString != "" {
		ParseToVariable(headersString, Headers)
	}

	if bodyString != "" {
		ParseToVariable(bodyString, Body)
	}

	if cookiesString != "" {
		ParseToVariable(cookiesString, Cookies)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().
		StringVarP(&headersString, "headers", "H", "", "Pass custom headers to server")

	rootCmd.PersistentFlags().
		StringVarP(&bodyString, "data", "d", "", "Pass data via request body to server")

	rootCmd.PersistentFlags().
		StringVarP(&cookiesString, "cookies", "c", "", "Send cookies with the request")

		// TODO
	rootCmd.PersistentFlags().
		StringVarP(&outputFilename, "output", "o", "", "Output the response to a file")

	rootCmd.PersistentFlags().
		BoolVarP(&includeHeaders, "include", "i", false, "Include headers on output")

	rootCmd.PersistentFlags().
		BoolVarP(&isVerbose, "verbose", "v", false, "Outputs detailed data about the request and response")
}

func ParseToVariable(str string, variable map[string]string) {
	xs := strings.Split(str, ";")
	for _, v := range xs {
		keyAndValue := strings.Split(v, ":")

		if len(keyAndValue) < 2 {
			fmt.Printf("error: wrong format, use key%svalue\n", ":")
			os.Exit(1)
		}

		key := strings.Trim(keyAndValue[0], " ")
		value := strings.Trim(keyAndValue[1], " ")
		variable[key] = value
	}

	fmt.Println(variable)
}
