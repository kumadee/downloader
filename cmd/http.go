package cmd

import (
	"fmt"
	http "github.com/kumadee/downloader/pkg/http"

	"github.com/spf13/cobra"
)

// httpCmd represents the http command
var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Download file over http(s)",
	Long:  `Download file over http(s).`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Download file:", url)
		err := http.DownloadFile(destinationPath, url)
		if err != nil {
			fmt.Println(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(httpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// httpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// httpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
