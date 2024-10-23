package config

type Config struct {
	DefaultDocumentType string                  `yaml:"default_document_type"`
	CustomValues        map[string]interface{}  `yaml:"custom_values"`
	DocumentTypes       map[string]DocumentType `yaml:"document_types"`
}

type DocumentType struct {
	TemplateFile     string                 `yaml:"template_file"`
	FileNamePattern  string                 `yaml:"file_name_pattern"`
	DirectoryPattern string                 `yaml:"directory_pattern"`
	CustomValues     map[string]interface{} `yaml:"custom_values"`
}
