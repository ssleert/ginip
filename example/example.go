package main

import (
	"fmt"

	"github.com/ssleert/ginip"
)

func main() {
	f, _ := ginip.Load("./example/example.ini")

	fmt.Println(f.GetValueString("package1", "name"))
	fmt.Println(f.GetValueString("package1", "version"))
	fmt.Println(f.GetValueString("package1", "asd"))
	fmt.Println(f.GetValueString("package2", "name"))
	fmt.Println(f.GetValueString("package2", "version"))
}
