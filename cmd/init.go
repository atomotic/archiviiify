package cmd

import (
	"fmt"
	"os"

	_ "embed"

	"github.com/spf13/cobra"
)

//go:embed resources/index.html
var index []byte

//go:embed resources/docker-compose.yml
var compose []byte

//go:embed resources/env
var env []byte

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "create directories and docker configurations.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		os.MkdirAll("./data/images", 0755)
		os.MkdirAll("./data/www", 0755)
		os.MkdirAll("./data/manifests", 0755)
		os.WriteFile("./data/www/index.html", index, 0644)
		os.WriteFile("./docker-compose.yml", compose, 0644)
		os.WriteFile("./.env", env, 0644)

		fmt.Println(`# init — the following directories have been created
├── archiviiify
├── .env
├── data
│   ├── images
│   ├── manifests
│   └── www
│       └── index.html
└── docker-compose.yml`)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
