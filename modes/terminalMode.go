package modes

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/Sirupsen/logrus"
)

// terminalModeHandler 控制台输出模式
func terminalModeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(terminalTplFile.FilePath())
	if err != nil {
		logrus.Fatal("Parse template file "+terminalTplFile.FilePath()+" err: ", err)
	}

	err = t.Execute(os.Stdout, makeData(r, map[string]interface{}{
		"colorStart": fmt.Sprintf("%c[%d;%d;%dm", 0x1B, 1, 0, 32),
		"colorEnd":   fmt.Sprintf("%c[0m", 0x1B),
	}))
	if err != nil {
		logrus.Fatal("Excute template file "+terminalTplFile.FilePath()+" err: ", err)
	}
}
