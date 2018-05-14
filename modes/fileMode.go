package modes

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/AielloChan/echoo/libs"
)

// fileModeHandler 输出为文件的模式
func fileModeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(logfileTplFile.FilePath())
	libs.ErrorHandler(err)

	libs.PrepareFile("./logs/log.txt")
	f, err := os.OpenFile("./logs/log.txt", os.O_WRONLY|os.O_APPEND, 0666)
	libs.ErrorHandler(err)
	defer f.Close()

	err = t.Execute(f, makeData(r))
	libs.ErrorHandler(err)
	log.Println(r.URL)
}
