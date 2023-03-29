package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "archiviiify",
	Short: "download JP2 from Internet Archive",
	Long:  `Archiviiify - download JP2 processed images from Internet Archive and rehost on IIPSRV IIIF.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}
