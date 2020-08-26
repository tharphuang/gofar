package text

import "strings"

var (
	// defaultTrimChars are the characters which are stripped by Trim* functions in default.
	defaultTrimChars = string([]byte{
		'\t', // Tab.
		'\v', // Vertical tab.
		'\n', // New line (line feed).
		'\r', // Carriage return.
		'\f', // New page.
		' ',  // Ordinary space.
		0x00, // NUL-byte.
		0x85, // Delete.
		0xA0, // Non-breaking space.
	})
)

// TrimLeft strips whitespace (or other characters) from the beginning of a string.
func TrimLeft(str string, characterMask ...string) string {
	if len(characterMask) == 0 {
		return strings.TrimLeft(str, defaultTrimChars)
	} else {
		return strings.TrimLeft(str, defaultTrimChars+characterMask[0])
	}
}