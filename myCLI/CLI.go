package main

import (
	"fmt"
	"myCLI/dcmd"
	"os"
)
//client主函数
func main() {
	//判断参数是不是为一个
	if len(os.Args) == 1 {
		defaultCmd()
	} else {
		//value := Commands[os.Args[1]]
		value := dcmd.Commands[os.Args[1]]
		if value != nil {
			value()
		} else {
			defaultCmd()
		}
	}
}
func defaultCmd() {
	fmt.Println()
	fmt.Println("Usage:	myCLI COMMAND")
	fmt.Println()
	fmt.Println("A docker client use by myself")
}

//docker ls

