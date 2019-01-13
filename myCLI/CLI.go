package main

import (
	"fmt"
	"myCLI/dcmd"
	"os"
)

//client主函数
func main() {
	//判断参数是不是只有一个
	if len(os.Args) == 1 {
		//如果只有一个,说明只输入了CLI,输出使用说明
		defaultCmd()
	} else {
		value := dcmd.Commands[os.Args[1]]
		if value != nil {
			value(os.Args[1:])
		} else {
			defaultCmd()
		}
	}
}
func defaultCmd() {
	fmt.Println()
	fmt.Println("Usage:	myCLI COMMAND")
	fmt.Println()
	fmt.Println("A docker client written by myself")
	fmt.Println("Commands:")
	fmt.Println()
	fmt.Printf("%-10s%s\n", "images", "List images")
}

//docker ls
