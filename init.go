package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/AielloChan/echoo/config"
)

func init() {
	flag.StringVar(&config.Host, "h", "0.0.0.0", config.HostInfo)
	flag.IntVar(&config.Port, "p", 8888, config.PortInfo)
	flag.StringVar(&config.Mode, "m", "echo", config.ModeInfo)
	flag.StringVar(&config.File, "f", "logs/log.txt", config.FileInfo)
	flag.String("v", "", config.VersionInfo)

	// 判断用户只否只是获取版本号
	if len(os.Args) > 1 && (os.Args[1] == "-v" || os.Args[1] == "-version") {
		// 直接输出版本号并退出
		fmt.Printf(config.Version)
		os.Exit(0)
	} else {
		flag.Parse()
	}
}
