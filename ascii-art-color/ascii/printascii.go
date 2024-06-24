package ascii

import (
	"fmt"
	"strings"
)

/*
function takes str which is string passed at argument one ,contentslice which is characters from filaname,
and index which tracks the number of lines per character.
it recursively print the provided string up to last line
*/
var ansiCodes = map[string]string{
	"black":          "30",
	"red":            "31",
	"green":          "32",
	"yellow":         "33",
	"blue":           "34",
	"magenta":        "35",
	"cyan":           "36",
	"white":          "37",
	"gray":           "90",
	"bright red":     "91",
	"bright green":   "92",
	"bright yellow":  "93",
	"bright blue":    "94",
	"bright magenta": "95",
	"bright cyan":    "96",
	"bright white":   "97",
}

/*
PrintAscii takes str , substr and color which are strings passed at argument one ,contentslice which is characters from filaname,
and index which tracks the number of lines per character.
it recursively print the provided string up to last line
*/
func PrintAscii(str, substr, color string, contentSlice []string, index int) {
	if index == 8 {
		return
	}
	indices := GetIndices(str, substr)
	track := 0
	count := 0
	// loop through each character in a str and prints it line by line.
	for i, char := range str {
		character := contentSlice[int(char)-32]                 // obtain char from contentslice
		character = strings.ReplaceAll(character, "\r\n", "\n") // thinkertoy
		lines := strings.Split(character, "\n")
		if color != "" && substr == "" {
			fmt.Printf("\033[%sm%s\033[0m", ansiCodes[color], lines[index])
		} else if substr != "" && len(indices) > 0 {
			if i >= indices[track] && i < indices[track]+len(substr) {
				fmt.Printf("\033[%sm%s\033[0m", ansiCodes[color], lines[index])
				count++
				if count == len(substr) && track < len(indices)-1 {
					track++
					count = 0
				}
			} else {
				fmt.Print(lines[index])
			}
		} else {
			fmt.Print(lines[index])
		}
	}
	fmt.Println()
	PrintAscii(str, substr, color, contentSlice, index+1)
}

func GetIndices(str, substr string) (indices []int) {
	if substr == "" {
		return
	}
	start := 0
	for {
		index := strings.Index(str[start:], substr)
		if index == -1 {
			break
		}
		indices = append(indices, start+index)
		start += index + len(substr)
	}
	return
}
