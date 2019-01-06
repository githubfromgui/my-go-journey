package string_concat_optimized

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	StringConcatWithJoin()
}

func StringConcatWithJoin() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
