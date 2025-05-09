package core

import (
	"fmt"
	"path/filepath"

	"github.com/matthewchivers/dodl/pkg/config"
	"github.com/matthewchivers/dodl/pkg/dateutils"
	"github.com/matthewchivers/dodl/pkg/filesystem"
	"github.com/matthewchivers/dodl/pkg/templating"
	"github.com/matthewchivers/dodl/pkg/validation"
	"github.com/matthewchivers/dodl/pkg/workspace"
)

// CreateCommand holds all the data/context required to create a new document
type CreateCommand struct {
	DocName      string
	DocType      config.DocumentType
	CustomFields map[string]interface{}
	Topic        string
	AppCtx       *AppContext
	Workspace    *workspace.Workspace
}

// Execute creates a new document based on the data/context in the CreateCommand
func (c *CreateCommand) Execute() error {
	refTime := c.AppCtx.ReferenceTime

	data := map[string]interface{}{
		"Today":     refTime,
		"Topic":     c.Topic,
		"WeekStart": dateutils.GetDefaultWeekStartDate(refTime),
	}

	data["DocName"] = c.DocName

	for k, v := range c.CustomFields {
		data[k] = v
	}

	for k, v := range c.DocType.CustomFields {
		data[k] = v
	}

	filename, err := templating.RenderTemplate(c.DocType.FileNamePattern, data, refTime)
	if err != nil {
		return err
	}
	if err := validation.ValidatePathPart(filename); err != nil {
		return err
	}

	dirParts := []string{}
	for _, part := range c.DocType.DirectoryPattern {
		dirPart, err := templating.RenderTemplate(part, data, refTime)
		if err != nil {
			return err
		}
		if err := validation.ValidatePathPart(dirPart); err != nil {
			return err
		}
		dirParts = append(dirParts, dirPart)
	}
	dirname := filepath.Join(dirParts...)

	templateData, err := c.Workspace.LoadTemplate(c.DocType.TemplateFile)
	if err != nil {
		return err
	}
	content, err := templating.RenderTemplate(string(templateData), data, refTime)
	if err != nil {
		return err
	}

	filepath := filepath.Join(c.Workspace.RootPath(), dirname, filename)

	msg, err := filesystem.EnsureFileExists(filepath, []byte(content))
	if err != nil {
		return err
	}

	if msg != "" {
		fmt.Println(msg)
	}

	return nil
}
