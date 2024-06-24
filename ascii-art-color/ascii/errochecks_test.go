package ascii

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

// Several test functions to test the functions in errochecks.go
// For each test, we initialize testCases to store our cases and run subtests on each case
// We capture the result by calling the faction with our tc.input
// Then we compare the result to the expectedRslt
// If they're not equal, we throw an error.
func TestIsPrintableAscii(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		expectedErr string
	}{
		{
			name:        "Chinese Characters",
			input:       "学中文",
			expectedErr: "学中文: Not within the printable ascii range",
		},
		{
			name:        "Printable Characters",
			input:       "123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			expectedErr: "nil",
		},
		{
			name:        "ü Characters",
			input:       "goürd",
			expectedErr: "ü: Not within the printable ascii range",
		},
		{
			name:        "Escape Character \\f",
			input:       "go\\fhome",
			expectedErr: "\\f: Not within the printable ascii range",
		},
		{
			name:        "Escape Character \\r",
			input:       "go\\rhome",
			expectedErr: "\\r: Not within the printable ascii range",
		},
		{
			name:        "Escape Character \\v",
			input:       "go\\vhome",
			expectedErr: "\\v: Not within the printable ascii range",
		},
		{
			name:        "Escape Character \\a",
			input:       "go\\ahome",
			expectedErr: "\\a: Not within the printable ascii range",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := IsPrintableAscii(tc.input)
			if tc.expectedErr == "nil" && err != nil {
				t.Errorf("Test%s Failed.\n IsPrintableAscii(\"%s\")\n"+
					"Expected:\nError: %v\n\n"+
					"Found: \nError: %v\n",
					tc.name, tc.input, tc.expectedErr, err)
			} else if err != nil && (err.Error() != tc.expectedErr) {
				t.Errorf("Test%s Failed.\n IsPrintableAscii(\"%s\")\n"+
					"Expected:\nError: %v\n\n"+
					"Found: \nError: %v\n",
					tc.name, tc.input, tc.expectedErr, err)
			}
		})
	}
}

func TestCheckFileTamper(t *testing.T) {
	testCases := []struct {
		name     string
		fileName string
	}{
		{
			name:     "Thinkertoy",
			fileName: "thinkertoy.txt",
		},
		{
			name:     "Standard",
			fileName: "standard.txt",
		},
		{
			name:     "Shadow",
			fileName: "shadow.txt",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			basePath := filepath.Join("..", "banner")
			filePath := os.DirFS(basePath)
			expectedErr := "nil"
			content, _ := fs.ReadFile(filePath, tc.fileName)
			err := CheckFileTamper(tc.fileName, content)
			if err != nil {
				t.Errorf("Test%s Failed.\n CheckFileValidity(\"%s\")\n"+
					"Expected Error: %v\n\n"+
					"Found Error: %v\n",
					tc.name, tc.fileName, expectedErr, err.Error())
			}
		})
	}
}
