// this package basically has functions which handles errors at different stages
package ascii

import (
	"fmt"
	"strings"
)

// the function takes string and returns an error
// the function checks if the string passed is printable or has any of the escape values,
// and if so the nonasciivalues  are stored in the varriable nonprintable and  the escapes value in foundEscapes respectively
// if both foundescapes and nonprintable varriable is not empty ,the error message is printed  along with the found values
func IsPrintableAscii(str string) error {
	var nonPrintables string
	var foundEscapes string
	errMessage := ": Not within the printable ascii range"
	for index, char := range str {

		escapes := "avrfb"
		var next byte
		if index < len(str)-1 {
			next = str[index+1]
		}
		// Check if the next character is an escape letter
		NextIsAnEscapeLetter := strings.ContainsAny(string(next), escapes)
		// Check if the current character is an escape character
		isAnEscape := (char == '\\' && NextIsAnEscapeLetter)
		// Check if the current character is non-printable
		isNonPrintable := ((char < ' ' || char > '~') && char != '\n')

		if isAnEscape {
			foundEscapes += "\\" + string(next)
		}
		if isNonPrintable {
			nonPrintables += string(char)
		}
	}
	// Construct error message if escape characters or non-printable characters are found
	if foundEscapes != "" {
		return fmt.Errorf("%s%s", foundEscapes, errMessage)
	} else if nonPrintables != "" {
		return fmt.Errorf("%s%s", nonPrintables, errMessage)
	}

	return nil
}

// function takes string  which is filename , content of filename respectively and then returns an error message
// if the name obtained does not match the expected length respectively, then error message is displayed
func CheckFileTamper(fileName string, content []byte) error {
	errMessage := " is tampered"
	lengthContent := len(content)
	// Checks if the length of the content matches the expected length of each file
	if fileName == "standard.txt" && lengthContent != 6623 {
		return fmt.Errorf("%s%s", fileName, errMessage)
	} else if fileName == "thinkertoy.txt" && lengthContent != 5558 {
		return fmt.Errorf("%s%s", fileName, errMessage)
	} else if fileName == "shadow.txt" && lengthContent != 7465 {
		return fmt.Errorf("%s%s", fileName, errMessage)
	}

	return nil
}
