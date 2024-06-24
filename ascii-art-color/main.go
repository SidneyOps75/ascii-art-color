package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"strings"

	"ascii-art-color/ascii"
)

var usage = fmt.Errorf(`Usage: go run . [OPTION] [STRING]

EX: go run . --color=<color> "something"`)

func main() {
	var str string
	var substr string
	fileName := "standard.txt"

	c := flag.String("color", "", "--color=<your color>")
	if err := ValidateFlag(); err != nil {
		fmt.Println(usage)
		return
	}
	flag.Parse()

	if flag.NFlag() == 0 {
		switch flag.NArg() {
		case 1:
			str = flag.Arg(0)
		case 2:
			str = flag.Arg(0)
			fileName = flag.Arg(1) + ".txt"
		default:
			fmt.Println(usage)

		}
	} else if flag.NFlag() == 1 {
		switch flag.NArg() {
		case 1:
			str = flag.Arg(0)
		case 2:
			if CheckFile(flag.Arg(1)) {
				fileName = flag.Arg(1) + ".txt"
				str = flag.Arg(0)
			} else {
				substr = flag.Arg(0)
				str = flag.Arg(1)
			}
		case 3:
			substr = flag.Arg(0)
			str = flag.Arg(1)
			fileName = flag.Arg(2) + ".txt"
		default:
			fmt.Println(usage)

		}
	} else {
		fmt.Print(usage)
	}
	str = strings.ReplaceAll(str, "\\t", "    ")
	str = strings.ReplaceAll(str, "\n", "\\n")
	err := ascii.IsPrintableAscii(str)
	if err != nil {
		fmt.Println(err)
		return
	}
	// Set the second argument for the banner file name. the default has been set to standard.txt
	filePath := os.DirFS("./banner")
	contentByte, err := fs.ReadFile(filePath, fileName)
	if err != nil {
		fmt.Print(err)
		return
	}
	if len(contentByte) == 0 {
		fmt.Println("Banner file is empty")
		return
	}
	// check if the banner file has been tampered with
	er := ascii.CheckFileTamper(fileName, contentByte)
	if er != nil {
		fmt.Println(er)
		return
	}
	contentString := string(contentByte[1:])
	if fileName == "thinkertoy.txt" {
		// convert all carriage returns to newlines
		contentString = strings.ReplaceAll(string(contentByte[2:]), "\r\n", "\n")
	}
	contentSlice := strings.Split(contentString, "\n\n")
	// Split the input string by "\\n" to get individual words
	words := strings.Split(str, "\\n")
	count := 0
	for _, str := range words {
		if str == "" {
			count++
			if count < len(words) {
				fmt.Println()
			}
		} else {
			// Print the ASCII representation of the word
			ascii.PrintAscii(str, substr, *c, contentSlice, 0)
		}
	}
}

func CheckFile(s string) bool {
	files := []string{"standard", "shadow", "thinkertoy"}
	for _, file := range files {
		if file == s {
			return true
		}
	}
	return false
}

func ValidateFlag() error {
	seenFlags := make(map[string]bool)
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "-") {
			if strings.HasPrefix(arg, "-color") {
				return usage
			} else if !strings.Contains(arg, "=") && strings.Contains(arg, "color") {
				return usage
			}
			flagName := strings.SplitN(arg[2:], "=", 2)[0]
			if seenFlags[flagName] {
				return usage
			}
			seenFlags[flagName] = true
		}
	}
	return nil
}
