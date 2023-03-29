package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/atomotic/archiviiify/internetarchive"
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/cobra"
)

var Item string

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "download an item from Internet Archive and generate an IIIF manifest",
	Long:  `download an item from Internet Archive and generate an IIIF manifest`,
	Run: func(cmd *cobra.Command, args []string) {
		if Item == "" {
			cmd.Help()
			os.Exit(0)
		}

		Logo()

		item, err := internetarchive.New(Item)
		if err != nil {
			fmt.Println("metadata error")
		}

		fmt.Printf("· downloading  %s\n", item.Metadata.Title)
		fmt.Printf("· from         %s\n", item.JP2Zip)

		if item.Downloaded() {
			fmt.Println("already downloaded")
		} else {
			err = item.Download()
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
		}

		time.Sleep(2 * time.Second)

		fmt.Println("· generating IIIF manifest")
		err = item.Manifest()
		if err != nil {
			fmt.Print(err)
			os.Exit(0)
		}
		fmt.Printf("view http://localhost:9000/?manifest=%s\n", item.Metadata.Identifier)

	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringVarP(&Item, "identifier", "i", "", "IA item identifier")
}
