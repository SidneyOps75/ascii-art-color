package main

import (
	"fmt"
	"os/exec"
	"testing"
)

// First, we initialize a testCases [] struct to store all our input and want cases
// We then loop through testCases and run subtests for each case
// We create a cmd using os/exec.Command to construct a command to run the program
// We then capture the output of the command (got) and compare it to want
// If they're not equal, we throw an error.
func TestMain(t *testing.T) {
	testCases := []struct {
		input  string
		banner string
		want   string
	}{
		{
			input:  "",
			banner: "standard",
			want:   "",
		},
		{
			input:  "No File",
			banner: "hello",
			want: `open hello.txt: no such file or directory$
`,
		},
		{
			input:  "\n",
			banner: "thinkertoy",
			want:   "$\n",
		},
		{
			input:  "Hello\n",
			banner: "shadow",
			want: `                                 $
_|    _|          _| _|          $
_|    _|   _|_|   _| _|   _|_|   $
_|_|_|_| _|_|_|_| _| _| _|    _| $
_|    _| _|       _| _| _|    _| $
_|    _|   _|_|_| _| _|   _|_|   $
                                 $
                                 $
$
`,
		},
		{
			input:  "1Hello 2There",
			banner: "standard",
			want: `     _    _          _   _                         _______   _                           $
 _  | |  | |        | | | |                ____   |__   __| | |                          $
/ | | |__| |   ___  | | | |   ___         |___ \     | |    | |__     ___   _ __    ___  $
| | |  __  |  / _ \ | | | |  / _ \          __) |    | |    |  _ \   / _ \ | '__|  / _ \ $
| | | |  | | |  __/ | | | | | (_) |        / __/     | |    | | | | |  __/ | |    |  __/ $
|_| |_|  |_|  \___| |_| |_|  \___/        |_____|    |_|    |_| |_|  \___| |_|     \___| $
                                                                                         $
                                                                                         $
`,
		},
		{
			input:  "{Hello There}",
			banner: "thinkertoy",
			want: `                                                          $
  o-o o  o     o o           o-O-o o                o-o   $
  |   |  |     | |             |   |                  |   $
o-O   O--O o-o | | o-o         |   O--o o-o o-o o-o   O-o $
  |   |  | |-' | | | |         |   |  | |-' |   |-'   |   $
  o-o o  o o-o o o o-o         o   o  o o-o o   o-o o-o   $
                                                          $
                                                          $
`,
		},
		{
			input:  "Hello\n\nThere",
			banner: "shadow",
			want: `                                 $
_|    _|          _| _|          $
_|    _|   _|_|   _| _|   _|_|   $
_|_|_|_| _|_|_|_| _| _| _|    _| $
_|    _| _|       _| _| _|    _| $
_|    _|   _|_|_| _| _|   _|_|   $
                                 $
                                 $
$
                                               $
_|_|_|_|_| _|                                  $
    _|     _|_|_|     _|_|   _|  _|_|   _|_|   $
    _|     _|    _| _|_|_|_| _|_|     _|_|_|_| $
    _|     _|    _| _|       _|       _|       $
    _|     _|    _|   _|_|_| _|         _|_|_| $
                                               $
                                               $
`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			cmd := exec.Command("sh", "-c", fmt.Sprintf("go run . \"%s\"  %s| cat -e", tc.input, tc.banner))
			output, err := cmd.Output()
			if err != nil {
				t.Fatalf("Error running program: %v", err)
			}
			got := string(output)
			if got != tc.want {
				t.Errorf("\ngot:\n%v\nwant:\n%v\n", got, tc.want)
			}
		})
	}
}
