package initialise

import (
	"fmt"
	"os"

	"github.com/matthewchivers/dodl/managers/workspace"
	"github.com/spf13/cobra"
)

type InitCommandHandler struct {
}

func NewInitCommandHandler() *InitCommandHandler {
	return &InitCommandHandler{}
}

func (h *InitCommandHandler) Handle(_ *cobra.Command, _ []string) {
	entryPath, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	wsp, err := workspace.GetManager(entryPath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = wsp.InitialiseWorkspace()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
