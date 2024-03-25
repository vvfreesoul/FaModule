package main

import (
	"code.hashdata.xyz/cloudberry/cloudberryui/server/resources"
	"fmt"
	"github.com/vvfreesoul/adder"
)

func main() {
	sum := adder.Add(1, 2)
	fmt.Println(sum)
	fmt.Println(resources.Password)

}
