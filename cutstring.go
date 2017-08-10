// ysoftman
// title : cutstring
// desc : http request 등의 파라미터 키,값을 분리하여 보기 편하게 보여준다.
// dependency
// go get github.com/fatih/color

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func getNextColorString(i int, str string) string {
	n := i % 6
	switch n {
	case 0:
		red := color.New(color.FgRed).SprintFunc()
		return red(str)
	case 1:
		green := color.New(color.FgGreen).SprintFunc()
		return green(str)
	case 2:
		yellow := color.New(color.FgYellow).SprintFunc()
		return yellow(str)
	case 3:
		blue := color.New(color.FgBlue).SprintFunc()
		return blue(str)
	case 4:
		magenta := color.New(color.FgMagenta).SprintFunc()
		return magenta(str)
	case 5:
		cyan := color.New(color.FgCyan).SprintFunc()
		return cyan(str)
	default:
		white := color.New(color.FgWhite).SprintFunc()
		return white(str)
	}

}

func main() {
	instr := ""
	if len(os.Args) == 2 {
		instr = os.Args[1]
	} else {
		return
	}

	// splitHttp := strings.Split(instr, "&")
	multisep := func(c rune) bool {
		return c == '&' || c == '?'
	}
	splitHttp := strings.FieldsFunc(instr, multisep)

	fmt.Println("[Result]")
	for i := 0; i < len(splitHttp); i++ {
		fmt.Fprintf(os.Stdout, "%v\n", getNextColorString(i, splitHttp[i]))
	}
}
