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

func (h *InitCommandHandler) Handle(cmd *cobra.Command, args []string) {
	wsp := workspace.NewWorkspaceManager()
	err := wsp.InitialiseWorkspace()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
