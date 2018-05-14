package modes

import (
	"html/template"
	"net/http"
	"os"

	"github.com/AielloChan/echoo/config"
	"github.com/AielloChan/echoo/libs"
	"github.com/Sirupsen/logrus"
)

// fileModeHandler 输出为文件的模式
func fileModeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(logfileTplFile.FilePath())
	if err != nil {
		logrus.Fatal("Parse template file "+logfileTplFile.FilePath()+" err: ", err)
	}

	libs.PrepareFile(config.File)
	f, err := os.OpenFile(config.File, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatal("Open log file to write failed: ", err)
	}
	defer f.Close()

	err = t.Execute(f, makeData(r))
	if err != nil {
		logrus.Fatal("Excute template file "+logfileTplFile.FilePath()+" err: ", err)
	}
	logrus.Info(r.URL)
}
