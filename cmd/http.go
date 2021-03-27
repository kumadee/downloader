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
}
