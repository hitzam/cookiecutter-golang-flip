package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "Run {{ cookiecutter.app_name }} worker",
	Long:  `Run {{ cookiecutter.app_name }} worker`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running your worker logic here....")
	},
}

func init() {
	rootCmd.AddCommand(workerCmd)
}
