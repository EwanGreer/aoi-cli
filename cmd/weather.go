/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
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

// weatherCmd represents the weather command
var weatherCmd = &cobra.Command{
	Use:   "weather",
	Short: "Gets the weather for belfast",
	Long:  `Gets the weather for belfast from wttr.in using format 3`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get("https://wttr.in/belfast?A") // /A return ANSI format
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		fmt.Print(string(b))
		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(weatherCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// weatherCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// weatherCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
