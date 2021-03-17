package config

import (
	"fmt"
)

const (
	// NAME app name
	NAME = "Echoo"
	// VERSION Current version
	VERSION = "1.5"
	// AUTHOR aiello chan
	AUTHOR = "Aiello Chan <aiello.chan@gmail.com>"
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
		"Set your host name | 设置 Host 名\n" +
		"0.0.0.0   : means that all client are allowed\n" +
		"            表示允许任何客户端访问" +
		"\n"
	// PortInfo PortInfo
	PortInfo = "" +
		"Set service work port | 设置该服务工作的端口号\n" +
		"8888      : default is 8888" +
		"          : 默认为 8888" +
		"\n"
	// ModeInfo ModeInfo
	ModeInfo = "" +
		"Choose working mode | 选择程序的工作模式\n" +
		"echo      : Return request infomation with html\n" +
		"            通过网页直接返回请求信息\n" +
		"terminal  : Output request infomation to terminal\n" +
		"            将请求信息输出到终端\n" +
		"file      : Write request infomation in log file\n" +
		"            将请求信息写入到日志文件\n" +
		"ws        : Real-time display info on webpage with websocket\n" +
		"            基于 Websocket 技术实时展示请求数据" +
		"\n"
	// FileInfo FileInfo
	FileInfo = "" +
		"File path : Set the log file storage location\n" +
		"            设置日志文件存放的位置" +
		"\n"
	// VersionInfo VersionInfo
	VersionInfo = "" +
		"Print version info\n" +
		"输出版本信息" +
		"\n"
)

var (
	// LogoChart logo 的彩色图形
	LogoChart = fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, 1, 0, 33,
		`
┌─┐┌─┐┬ ┬┌─┐┌─┐
├┤ │  ├─┤│ ││ │
└─┘└─┘┴ ┴└─┘└─┘`, 0x1B)
	// Version Version
	Version = "" +
		LogoChart + "\n" +
		"Name     : " + NAME + "\n" +
		"Version  : " + VERSION + "\n" +
		"Author   : " + AUTHOR + "\n" +
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
