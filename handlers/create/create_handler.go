package create

import (
	"fmt"

	"github.com/spf13/cobra"
)

type CommandHandler struct {
}

func NewCreateCommandHandler() *CommandHandler {
	return &CommandHandler{}
}

func (h *CommandHandler) Handle(_ *cobra.Command, args []string) {
	docType := args[0]
	fmt.Println("Creating a new document of type:", docType)
}
