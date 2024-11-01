package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatePathPart(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		wantErr  bool
		errorMsg string
	}{
		// Valid directory cases
		{"Valid simple name", "notes", false, ""},
		{"Valid name with numbers", "notes123", false, ""},
		{"Valid name with dash", "my-notes", false, ""},
		{"Valid name with underscore", "my_notes", false, ""},
		{"Valid name with spaces", "my notes", false, ""},

		// Invalid directory substring cases
		{"Invalid relative path", "../secrets", true, `invalid substring ".." found in path pattern`},

		// Invalid directory character cases
		{"Invalid character slash", "my/notes", true, `invalid character "/" found in path pattern`},
		{"Invalid character backslash", "my\\notes", true, `invalid character "\\" found in path pattern`},
		{"Invalid character colon", "my:notes", true, `invalid character ":" found in path pattern`},
		{"Invalid character asterisk", "my*notes", true, `invalid character "*" found in path pattern`},
		{"Invalid character question mark", "my?notes", true, `invalid character "?" found in path pattern`},
		{"Invalid character double quotes", `my"notes`, true, `invalid character "\"" found in path pattern`},
		{"Invalid character less than", "my<notes", true, `invalid character "<" found in path pattern`},
		{"Invalid character greater than", "my>notes", true, `invalid character ">" found in path pattern`},
		{"Invalid character pipe", "my|notes", true, `invalid character "|" found in path pattern`},

		// Control character case
		{"Control character null", "my\000notes", true, "control character found in path pattern"},

		// Edge cases
		{"Empty string", "", true, "empty string (or string with only whitespace) found in path"},
		{"Whitespace only", "   ", true, "empty string (or string with only whitespace) found in path"},

		// Valid File name cases
		{"Valid file name", "my-notes.md", false, ""},

		// Invalid file name cases
		{"Invalid file name", "my/notes.md", true, `invalid character "/" found in path pattern`},
		{"Invalid file name", "my\\notes.md", true, `invalid character "\\" found in path pattern`},
		{"Invalid file name", "my:notes.md", true, `invalid character ":" found in path pattern`},
		{"Invalid file name", "my*notes.md", true, `invalid character "*" found in path pattern`},
		{"Invalid file name", "my?notes.md", true, `invalid character "?" found in path pattern`},
		{"Invalid file name", `my"notes.md`, true, `invalid character "\"" found in path pattern`},
		{"Invalid file name", "my<notes.md", true, `invalid character "<" found in path pattern`},
		{"Invalid file name", "my>notes.md", true, `invalid character ">" found in path pattern`},
		{"Invalid file name", "my|notes.md", true, `invalid character "|" found in path pattern`},
		{"Invalid file name", "my\000notes.md", true, "control character found in path pattern"},
		{"Invalid file name", "my..notes.md", true, `invalid substring ".." found in path pattern`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePathPart(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
