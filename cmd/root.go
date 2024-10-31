package cmd

import (
	"fmt"
	"os"

	"github.com/matthewchivers/dodl/cmd/create"
	"github.com/matthewchivers/dodl/cmd/initialise"
	"github.com/matthewchivers/dodl/cmd/status"
	wd "github.com/matthewchivers/dodl/pkg/workingdir"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dodl",
	Short: "dodl is a document creation tool",
	Long:  `dodl is a command-line tool for creating structured documents using templates.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to dodl! Run 'dodl help' to get started.")
	},
}

// init registers the subcommands with the root command
func init() {
	createCmd := create.NewCreateCmd(&wd.DefaultWorkingDirProvider{})
	rootCmd.AddCommand(createCmd)
	initCmd := initialise.NewInitCmd(&wd.DefaultWorkingDirProvider{})
	rootCmd.AddCommand(initCmd)
	statusCmd := status.NewStatusCmd(&wd.DefaultWorkingDirProvider{})
	rootCmd.AddCommand(statusCmd)
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
