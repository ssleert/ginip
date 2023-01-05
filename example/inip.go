package main

import (
	"fmt"
	"os"

	"github.com/ssleert/ginip"
)

func main() {
	ini, _ := ginip.Load(os.Args[1])
	s, _ := ini.GetValueString(os.Args[2], os.Args[3])
	fmt.Println(s)
}
