package validation

import (
	"fmt"
	"strings"
	"unicode"
)

var disallowedSubstrings = []string{".."}
var disallowedChars = []rune{'/', '\\', ':', '*', '?', '"', '<', '>', '|'}

func ValidatePathPart(part string) error {
	// Check for disallowed substrings
	if strings.TrimSpace(part) == "" {
		return fmt.Errorf("empty string (or string with only whitespace) found in path")
	}
	for _, substr := range disallowedSubstrings {
		if strings.Contains(part, substr) {
			return fmt.Errorf("invalid substring %q found in path pattern: %q", substr, part)
		}
	}

	// Check for disallowed characters
	for _, char := range part {
		if unicode.IsControl(char) {
			return fmt.Errorf("control character found in path pattern: %q", part)
		}
		for _, disallowed := range disallowedChars {
			if char == disallowed {
				return fmt.Errorf("invalid character %q found in path pattern: %q", string(char), part)
			}
		}
	}

	return nil
}
