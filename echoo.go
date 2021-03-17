package main

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/AielloChan/echoo/asset"
	"github.com/AielloChan/echoo/config"
	"github.com/AielloChan/echoo/modes"

	"github.com/sirupsen/logrus"
)

// 主函数
func main() {
	releaseFiles()

	port := strconv.Itoa(config.Port)
	url := "http://" + config.Host + ":" + port
	localUrl := "http://127.0.0.1:" + port

	logrus.Info("Echoo is serving...")
	logrus.Info("You can access at " + url + " or " + localUrl)

	modes.RunWithMode(config.Mode, config.Host, config.Port, config.File)

}

// 释放资源文件
func releaseFiles() {
	// 释放文件
	isSuccess := true
	dirs := []string{"static", "view"} // 设置需要释放的目录

	for _, dir := range dirs {
		// 解压dir目录到当前目录
		if err := asset.RestoreAssets("./", dir); err != nil {
			isSuccess = false
			logrus.Error("Extract asset failed", err)
			break
		}
	}
	if !isSuccess {
		for _, dir := range dirs {
			os.RemoveAll(filepath.Join("./", dir))
		}
	}
}
