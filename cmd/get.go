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

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "download an item from Internet Archive and generate an IIIF manifest",
	Long:  `download an item from Internet Archive and generate an IIIF manifest`,
	Run: func(cmd *cobra.Command, args []string) {
		if Item == "" {
			cmd.Help()
			os.Exit(0)
		}

		fmt.Printf("archiviiify\n\n")

		item, err := internetarchive.New(Item)
		if err != nil {
			fmt.Println("metadata error")
		}

		fmt.Printf("[1/3] Downloading %s:\n", item.Metadata.Title)
		fmt.Printf(" Source: %s\n", item.JP2Zip)

		if item.Downloaded() {
			fmt.Println(" Item already downloaded")
		} else {
			err = item.Download()
			if err != nil {
				fmt.Println(err)
				os.Exit(0)
			}
		}

		time.Sleep(2 * time.Second)

		fmt.Println("[2/3] Generating IIIF manifest")
		err = item.Manifest()
		if err != nil {
			fmt.Print(err)
			os.Exit(0)
		}
		fmt.Printf("[3/3] View the item:\n %s/?manifest=%s\n", os.Getenv("HOSTNAME"), item.Metadata.Identifier)

	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().StringVarP(&Item, "identifier", "i", "", "IA item identifier")
}
