package main

import (
	"github.com/ssleert/ginip"
	"os"
	"fmt"
)

func main() {
	ini, _ := ginip.Load(os.Args[1])
	s, _ := ini.GetValueString(os.Args[2], os.Args[3])
	fmt.Println(s)
}
