package config

// Config is a struct that holds configuration.
type Config struct {
	// DocTypes is a list of document types
	DocTypes []DocType `yaml:"docTypes"`
}

func (c *Config) Apply(cfg *Config) {
	if cfg == nil {
		return
	}
	if len(cfg.DocTypes) > 0 {
		for _, docToApply := range cfg.DocTypes {
			c.ApplyDoc(docToApply)
		}
	}
}

func (c *Config) ApplyDoc(ovrDoc DocType) {
	found := false
	for i, doc := range c.DocTypes {
		if doc.ID == ovrDoc.ID {
			// Modify the actual element in the slice
			c.DocTypes[i].Apply(&ovrDoc)
			found = true
			break // Exit loop once the document is found and applied
		}
	}
	if !found {
		c.DocTypes = append(c.DocTypes, ovrDoc)
	}
}
