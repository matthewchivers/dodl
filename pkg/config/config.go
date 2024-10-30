package config

// Config holds all the configuration data for dodl.
type Config struct {
	// DefaultDocumentType is the default document type to use if none is specified.
	DefaultDocumentType string                  `yaml:"default_document_type"`

	// CustomFields are the custom fields that can be used in document templates.
	// At this level, these are the custom fields that are available to all document types.
	CustomFields        map[string]interface{}  `yaml:"custom_fields"`

	// DocumentTypes are the available document types.
	DocumentTypes       map[string]DocumentType `yaml:"document_types"`
}

// DocumentType holds the configuration data for a document type.
type DocumentType struct {
	// TemplateFile is the path to the template file for this document type.
	// The path is relative to the .dodl/templates directory.
	TemplateFile     string                 `yaml:"template_file"`

	// FileNamePattern is the pattern to use when generating the filename for a document.
	// The pattern is a Go template string.
	FileNamePattern  string                 `yaml:"file_name_pattern"`

	// DirectoryPattern is the pattern to use when generating the directory for a document.
	// The pattern is a Go template string, and the directory is relative to the workspace root.
	DirectoryPattern string                 `yaml:"directory_pattern"`

	// CustomFields are the custom fields that can be used in document templates.
	// At this level, these are the custom fields that are available to only this document type.
	CustomFields     map[string]interface{} `yaml:"custom_fields"`
}
