package main

import (
	"fmt"
	"strconv"

	"github.com/AielloChan/echoo/modes"
)

//
const (
	// NAME app name
	NAME = "EchoX"
	// VERSION Current version
	VERSION = "0.1"
	// AUTHER aiello chan
	AUTHER = "Aiello Chan <aiello.chan@gmail.com>"
	// LICENCE GPL open source licence
	LICENCE = "GPL"
	// REPO github repository addr
	REPO = ""
)

// 运行参数
var (
	host string
	port int
	mode string // echo|terminal|file|ws
	file string
)

// 主函数
func main() {
	fmt.Println("Echoo serving at " + "http://" + host + ":" + strconv.Itoa(port))

	modes.RunWithMode(mode, host, port, file)

}
