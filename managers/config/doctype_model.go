package config

// DocType is the document type struct.
// It provides a way to define the structure of a document type in the configuration file
type DocType struct {
	// ID is the identifier for the document type (e.g. meeting, note, etc.)
	ID string `yaml:"id"`

	// DirectoryPattern is the pattern to use when creating a subdirectory
	DirectoryPattern string `yaml:"directoryPattern,omitempty"`

	// FileNamePattern is the pattern to use when creating a file name
	FileNamePattern string `yaml:"fileNamePattern,omitempty"`

	// Topic is a name to be used for templating (e.g. a meeting about a certain topic)
	// Expect this to be primarily set using cli params, but can be set in the config file
	Topics []string `yaml:"topics,omitempty"`

	// Editor is the editor to use when opening files
	Editor string `yaml:"editor,omitempty"`
}

// Apply "applies" the given DocType to the current DocType.
// Essentially, it merges the given DocType into the current DocType.
func (d *DocType) Apply(docType *DocType) {
	if docType == nil {
		return
	}
	if docType.DirectoryPattern != "" {
		d.DirectoryPattern = docType.DirectoryPattern
	}
	if docType.FileNamePattern != "" {
		d.FileNamePattern = docType.FileNamePattern
	}
	if len(docType.Topics) > 0 {
		for _, topic := range docType.Topics {
			d.AppendTopicNoDuplicates(topic)
		}
	}
	if docType.Editor != "" {
		d.Editor = docType.Editor
	}
}

// AppendTopicNoDuplicates appends the given topic to the current DocType's topics if it doesn't already exist.
func (d *DocType) AppendTopicNoDuplicates(topic string) {
	for _, t := range d.Topics {
		if t == topic {
			return
		}
	}
	d.Topics = append(d.Topics, topic)
}
