package config

import (
	"fmt"
)

const (
	// NAME app name
	NAME = "Echoo"
	// VERSION Current version
	VERSION = "1.0"
	// AUTHER aiello chan
	AUTHER = "Aiello Chan <aiello.chan@gmail.com>"
	// LICENCE GPL open source licence
	LICENCE = "GPLv3"
	// REPO github repository addr
	REPO = "https://github.com/AielloChan/echoo"
)

// 运行参数
var (
	Host string
	Port int
	Mode string // echo|terminal|file|ws
	File string
)

const (
	// HostInfo HostInfo
	HostInfo = "" +
		"Host      : Set your host name\n" +
		"0.0.0.0 means that all client are allowed\n" +
		"\n" +
		"0.0.0.0 表示允许任何客户端访问" +
		"\n"
	// PortInfo PortInfo
	PortInfo = "" +
		"Port      : Service work port\n" +
		"\n" +
		"设置该服务工作的端口号" +
		"\n"
	// ModeInfo ModeInfo
	ModeInfo = "" +
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
	// FileInfo FileInfo
	FileInfo = "" +
		"File path : Set the log file storage location\n" +
		"\n" +
		"设置日志文件存放的位置" +
		"\n"
	// VersionInfo VersionInfo
	VersionInfo = "" +
		"Print version\n" +
		"\n" +
		"输出版本信息" +
		"\n"
)

var (
	// LogoChart logo 的彩色图形
	LogoChart = fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, 1, 0, 33,
		""+
			"┌─┐┌─┐┬ ┬┌─┐┌─┐\n"+
			"├┤ │  ├─┤│ ││ │\n"+
			"└─┘└─┘┴ ┴└─┘└─┘\n", 0x1B)
	// Version Version
	Version = "" +
		LogoChart + "\n" +
		"Name     : " + NAME + "\n" +
		"Version  : " + VERSION + "\n" +
		"Auther   : " + AUTHER + "\n" +
		"Licence  : " + LICENCE + "\n" +
		"Repo     : " + REPO
)

// SocketData websocket data type
type SocketData struct {
	UUID string `json:"uuid"`
}

func (sd SocketData) String() string {
	return sd.UUID
}
