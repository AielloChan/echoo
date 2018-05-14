package modes

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/AielloChan/echoo/libs"
	"golang.org/x/net/websocket"
)

// 模版地址
var (
	// tplBase template base path
	tplBaseDir      = "view"
	echoTplFile     = libs.TplFile{Base: tplBaseDir, Path: "", FileName: "echo.html"}
	wsTplFile       = libs.TplFile{Base: tplBaseDir, Path: "", FileName: "ws.html"}
	terminalTplFile = libs.TplFile{Base: tplBaseDir, Path: "", FileName: "terminal.tpl"}
	logfileTplFile  = libs.TplFile{Base: tplBaseDir, Path: "", FileName: "logfile.tpl"}
)

var (
	logfilePath  string
	processQueue []func(http.ResponseWriter, *http.Request)
)

// RunWithMode 以指定模式运行
func RunWithMode(mode string, host string, port int, file string) {
	// 配置一些处理器所需要的参数
	logfilePath = file
	hostURL := url.URL{
		Scheme: "http",
		Host:   host + ":" + strconv.Itoa(port),
	}
	openWS := false

	// 根据模式来运行不同的处理器
	switch mode {
	case "echo":
		processQueue = append(processQueue, echoModeHandler)
	case "terminal":
		processQueue = append(processQueue, terminalModeHandler)
	case "file":
		processQueue = append(processQueue, fileModeHandler)
	case "ws":
		processQueue = append(processQueue, wsModeHandler)
		openWS = true
	default:
		log.Println("Try run 'echoX -h'")
		os.Exit(1)
	}

	startAPISer(hostURL, dispatcher, openWS)
}

// 服务分发器,方便以后扩展
func dispatcher(w http.ResponseWriter, r *http.Request) {
	// 处理队列中的任务
	for _, f := range processQueue {
		f(w, r)
	}
}

// 启动 echo api 服务器
func startAPISer(hostURL url.URL, h func(http.ResponseWriter, *http.Request), openWS bool) {

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/favicon.ico")
	})
	if openWS {
		http.Handle("/ws", websocket.Handler(wsHandler))
	}
	http.HandleFunc("/", h)
	// http.Handle("/echo", websocket.Handler(EchoServer))
	err := http.ListenAndServe(hostURL.Host, nil)
	// log.Println(NAME + " are already running at " + curURL.String())
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// 从用户请求中剥离用于展示的数据
func makeData(r *http.Request, more ...map[string]interface{}) map[string]interface{} {
	data := map[string]interface{}{
		"time":          template.HTML(time.Now().String()),
		"method":        r.Method,
		"url":           r.URL.String(),
		"protocol":      r.Proto,
		"host":          r.Host,
		"fullUrl":       libs.GetFullURL(r),
		"cookie":        strings.Join(libs.FmtCookies(r.Cookies()), ","),
		"contentLength": r.ContentLength,
		"remoteAddr":    r.RemoteAddr,
		"headers":       libs.MapReducer(r.Header),
		"urlParams":     libs.MapReducer(r.URL.Query()),
		"body":          template.HTML(libs.ReadAll(r.Body)),
	}

	// 可追加个性化参数
	if len(more) > 0 {
		for key, val := range more[0] {
			data[key] = val
		}
	}

	return data
}
