package modes

import (
	"html/template"
	"net/http"

	"github.com/Sirupsen/logrus"
)

// echoModeHandler 负责处理 echo 模式的业务
func echoModeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(echoTplFile.FilePath())
	if err != nil {
		logrus.Fatal("Parse template file "+echoTplFile.FilePath()+" err: ", err)
	}

	err = t.Execute(w, makeData(r))
	if err != nil {
		logrus.Fatal("Excute template file "+echoTplFile.FilePath()+" err: ", err)
	}
	logrus.Info(r.URL)
}
