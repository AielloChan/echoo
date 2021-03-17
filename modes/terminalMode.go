package modes

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

// terminalModeHandler 控制台输出模式
func terminalModeHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(terminalTplFile.FilePath())
	if err != nil {
		logrus.Fatal("Parse template file "+terminalTplFile.FilePath()+" err: ", err)
	}
	data := map[string]interface{}{
		"colorStart": fmt.Sprintf("%c[%d;%d;%dm", 0x1B, 1, 0, 32),
		"colorEnd":   fmt.Sprintf("%c[0m", 0x1B),
	}

	err = t.Execute(os.Stdout, makeData(r, data))

	if err != nil {
		logrus.Fatal("Excute template file "+terminalTplFile.FilePath()+" err: ", err)
	}
}
