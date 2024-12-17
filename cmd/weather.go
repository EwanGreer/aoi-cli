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
	Long:  `Gets the weather for belfast from wttr.in in ANSI format`,
	Run: func(cmd *cobra.Command, args []string) {
		location, err := cmd.Flags().GetString("location")
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		resp, err := http.Get(fmt.Sprintf("https://wttr.in/%s?A", location))
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

	weatherCmd.Flags().StringP("location", "l", "belfast", "The location to be queried for weather data")
}
