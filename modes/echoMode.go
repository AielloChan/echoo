package modes

import (
	"html/template"
	"log"
	"net/http"

	"github.com/AielloChan/echoo/libs"
)

// echoModeHandler 负责处理 echo 模式的业务
func echoModeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(echoTplFile.FilePath())
	libs.ErrorHandler(err)

	err = t.Execute(w, makeData(r))
	libs.ErrorHandler(err)
	log.Println(r.URL)
}
