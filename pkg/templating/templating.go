package templating

import (
	"bytes"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

// RenderTemplate renders a template string using the provided data map.
// It returns the rendered string or an error if rendering fails.
func RenderTemplate(templateStr string, data map[string]interface{}) (string, error) {
	funcMap := sprig.TxtFuncMap()
	funcMap = addCustomFuncs(funcMap)

	// Parse the template string
	tmpl, err := template.New("template").Funcs(funcMap).Parse(templateStr)
	if err != nil {
		return "", err
	}

	// Execute the template with the data
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
