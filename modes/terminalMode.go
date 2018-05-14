package modes

import (
	"html/template"
	"net/http"
	"os"

	"github.com/AielloChan/echoo/libs"
)

// terminalModeHandler 控制台输出模式
func terminalModeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(terminalTplFile.FilePath())
	libs.ErrorHandler(err)

	err = t.Execute(os.Stdout, makeData(r))
	libs.ErrorHandler(err)
}
