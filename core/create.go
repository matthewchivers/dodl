package core

import (
	"fmt"
	"path/filepath"

	"github.com/matthewchivers/dodl/config"
	"github.com/matthewchivers/dodl/filesystem"
	"github.com/matthewchivers/dodl/templating"
	"github.com/matthewchivers/dodl/workspace"
)

type CreateCommand struct {
	DocName      string
	DocType      config.DocumentType
	CustomFields map[string]interface{}
	Topic        string
	AppCtx       *AppContext
}

func (c *CreateCommand) Execute() error {
	data := map[string]interface{}{
		"Today": c.AppCtx.StartTime,
		"Topic": c.Topic,
	}

	data["DocName"] = c.DocName

	for k, v := range c.CustomFields {
		data[k] = v
	}

	for k, v := range c.DocType.CustomFields {
		data[k] = v
	}

	filename, err := templating.RenderTemplate(c.DocType.FileNamePattern, data)
	if err != nil {
		return err
	}

	dirname, err := templating.RenderTemplate(c.DocType.DirectoryPattern, data)
	if err != nil {
		return err
	}

	templateData, err := workspace.LoadTemplate(c.AppCtx.WorkingDir, c.DocType.TemplateFile)
	if err != nil {
		return err
	}
	content, err := templating.RenderTemplate(string(templateData), data)
	if err != nil {
		return err
	}

	filepath := filepath.Join(c.AppCtx.WorkspaceRoot, dirname, filename)

	err = filesystem.EnsureFileExists(filepath, []byte(content))
	if err != nil {
		return err
	}

	fmt.Printf("Document created at %q\n", filepath)

	return nil
}
