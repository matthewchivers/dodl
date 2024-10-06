package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dodl",
	Short: "dodl is a document creation tool",
	Long:  `dodl is a command-line tool for creating structured documents using templates.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to dodl! Run 'dodl help' to get started.")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
