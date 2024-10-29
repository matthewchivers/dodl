package config

type Config struct {
	DefaultDocumentType string                  `yaml:"default_document_type"`
	CustomFields        map[string]interface{}  `yaml:"custom_fields"`
	DocumentTypes       map[string]DocumentType `yaml:"document_types"`
}

type DocumentType struct {
	TemplateFile     string                 `yaml:"template_file"`
	FileNamePattern  string                 `yaml:"file_name_pattern"`
	DirectoryPattern string                 `yaml:"directory_pattern"`
	CustomFields     map[string]interface{} `yaml:"custom_fields"`
}
