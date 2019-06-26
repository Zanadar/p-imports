package simple_imports

import (
	foo "fmt"
	. "os"
)

func test() {
	foo.Println(Getpagesize())
}
