package cmd

import (
	"github.com/matthewchivers/dodl/handlers/initialise"
	"github.com/spf13/cobra"
)

var (
	initHandler = initialise.NewInitCommandHandler()
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new dodl register",
	Run:   initHandler.Handle,
}

func init() {
	rootCmd.AddCommand(initCmd)
}
