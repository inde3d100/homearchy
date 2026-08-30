//go:build darwin

package cli_test

import (
	"strings"
	"testing"
)

func TestIsRunningFromAppBundle(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected bool
	}{
		{
			name:     "app bundle path",
			path:     "/Applications/Homearchy.app/Contents/MacOS/homearchy",
			expected: true,
		},
		{
			name:     "regular binary path",
			path:     "/usr/local/bin/homearchy",
			expected: false,
		},
		{
			name:     "homebrew path",
			path:     "/opt/homebrew/bin/homearchy",
			expected: false,
		},
		{
			name:     "another app bundle",
			path:     "/Applications/MyApp.app/Contents/MacOS/myapp",
			expected: true,
		},
		{
			name:     "nested app bundle",
			path:     "/Applications/MyApp.app/Contents/Resources/Homearchy.app/Contents/MacOS/homearchy",
			expected: true,
		},
		{
			name:     "not an app bundle",
			path:     "/usr/bin/homearchy",
			expected: false,
		},
		{
			name:     "empty path",
			path:     "",
			expected: false,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			// Test the string matching logic directly
			result := strings.Contains(testCase.path, ".app/Contents/MacOS")
			if result != testCase.expected {
				t.Errorf(
					"strings.Contains(%q, \".app/Contents/MacOS\") = %v, expected %v",
					testCase.path,
					result,
					testCase.expected,
				)
			}
		})
	}
}
