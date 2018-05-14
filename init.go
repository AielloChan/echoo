package main

import (
	"flag"
	"log"
	"os"
)

func init() {
	const (
		hostInfo = "" +
			"Host      : Set your host name\n" +
			"0.0.0.0 means that all client are allowed\n" +
			"\n" +
			"0.0.0.0 表示允许任何客户端访问" +
			"\n"
		portInfo = "" +
			"Port      : Service work port\n" +
			"\n" +
			"设置该服务工作的端口号" +
			"\n"
		modeInfo = "" +
			"Interactive mode\n" +
			"echo      : Return request infomation with html\n" +
			"terminal  : Output request infomation to terminal\n" +
			"file      : Write request infomation in log file\n" +
			"ws        : Real-time display info on webpage with websocket\n" +
			"\n" +
			"echo      : 通过网页直接返回请求信息\n" +
			"terminal  : 将请求信息输出到终端\n" +
			"file      : 将请求信息写入到日志文件\n" +
			"ws        : 基于 Websocket 技术实时展示请求数据" +
			"\n"
		fileInfo = "" +
			"File path : Set the log file storage location\n" +
			"\n" +
			"设置日志文件存放的位置" +
			"\n"
		versionInfo = "" +
			"Print version\n" +
			"\n" +
			"输出版本信息" +
			"\n"
		version = "" +
			"	Name     : " + NAME + "\n" +
			"	Version  : " + VERSION + "\n" +
			"	Auther   : " + AUTHER + "\n" +
			"	Licence  : " + LICENCE + "\n" +
			"	Repo     : " + REPO
	)

	flag.StringVar(&host, "h", "0.0.0.0", hostInfo)
	flag.IntVar(&port, "p", 8888, portInfo)
	flag.StringVar(&mode, "m", "echo", modeInfo)
	flag.StringVar(&file, "f", "./log.txt", fileInfo)
	flag.String("v", "", versionInfo)

	// 判断用户只否只是获取版本号
	if len(os.Args) > 1 && (os.Args[1] == "-v" || os.Args[1] == "-version") {
		// 直接输出版本号并退出
		log.Print(version)
		os.Exit(0)
	} else {
		flag.Parse()
	}
}
