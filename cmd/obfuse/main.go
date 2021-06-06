// shamefully taken from
// https://github.com/GH0st3rs/obfus/blob/master/obfus.go
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
	"os"
	"unsafe"
)

var (
	s = flag.String("s", "", "String to obfuscate.")
	v = flag.String("v", "str", "Variable name.")
	r = flag.Bool("r", false, "Print summarized code.")
)

const (
	EAX = uint8(unsafe.Sizeof(true))
	ONE = "EAX"
)

func getNumber(n byte) (buf string) {
	var arr []byte
	for n > EAX {
		if n%2 == EAX {
			arr = append(arr, EAX)
		} else {
			arr = append(arr, 0)
		}
		n = n >> EAX
	}

	buf = ONE
	rand.Seed(time.Now().Unix())

	for i := len(arr) - 1; i >= 0; i-- {
		buf = fmt.Sprintf("%s<<%s", buf, ONE)
		if arr[i] == EAX {
			if rand.Intn(2) == 0 {
				buf = fmt.Sprintf("(%s^%s)", buf, ONE)
			} else {
				buf = fmt.Sprintf("(%s|%s)", buf, ONE)
			}
		}
	}
	return buf
}

func TextToCode(txt string) string {
	b := []byte(txt)
	tmp := "var " + *v + " []byte\n"
	for _, item := range b {
		tmp = fmt.Sprintf("%s\n" + *v + " = append(" + *v +", %s)", tmp, getNumber(item))
	}
	tmp += "\nfmt.Println(string(" + *v + "))"
	return tmp
}

func main() {
    flag.Parse()

        if (len(os.Args) < 2) || *s == "" {
	        flag.PrintDefaults()
	        os.Exit(1)
        }

	if *r {
		fmt.Println(TextToCode(*s))
		fmt.Println("}\n")
	} else {
		fmt.Println("package main\n")

		fmt.Println("import (")
		fmt.Println("	\"fmt\"")
		fmt.Println("	\"unsafe\"")
		fmt.Println(")\n")

		fmt.Println("const (")
		fmt.Println("	EAX = uint8(unsafe.Sizeof(true))")
		fmt.Println(")\n")

		fmt.Println("func main() {")
		fmt.Println(TextToCode(*s))
		fmt.Println("}\n")
	}
}
