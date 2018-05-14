package modes

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/AielloChan/echoo/config"
	"github.com/AielloChan/echoo/libs"
	"github.com/Sirupsen/logrus"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/net/websocket"
)

const (
	// cTimeout 会话超时
	cTimeout = 600 // 秒
	// cUUIDLength uuid 的长度，为固定值
	cUUIDLength = 36
	// cServerURLPostfix 服务器访问时 url 的后缀
	cServerURLPostfix = "server"
	// cClientURLPostfix 客户端访问时 url 的后缀
	cClientURLPostfix = "client"
)

var (
	connPool = make(map[string]map[*websocket.Conn]bool)
)

// EchoServer Echo the data received on the WebSocket.
func wsHandler(ws *websocket.Conn) {
	logrus.Info("New client connected")
	msg := &config.SocketData{}
	for {
		err := websocket.JSON.Receive(ws, &msg)
		if err != nil {
			logrus.Info("Client disconnected ", err)
			break
		}
		// 从 ws 消息中获取目标 uuid
		tarUUID := msg.UUID

		if len(tarUUID) != cUUIDLength {
			// 错误输入
			logrus.Error("Wrong ws message", msg)
			websocket.JSON.Send(ws, map[string]string{"msg": msg.String()})
			ws.Close()
		}
		// 判断该 uuid 是否已经存在
		if wsConns, ok := connPool[tarUUID]; ok {
			// 已经存在
			// 再判断该链接是否也存在过
			if _, o := wsConns[ws]; o {
				// 该链接已经存在，不做任何操作

			} else {
				connPool[tarUUID][ws] = true
			}
		} else {
			// 将新的 ws 加入到该 UUID 监控列表中
			connPool[tarUUID][ws] = true
			// 给新的 ws 设置定时器，超时会被移除
			libs.RegTime(tarUUID)
		}
	}
}

func wsModeHandler(w http.ResponseWriter, r *http.Request) {
	curHost := libs.GetFullHost(r)
	paths, ok := libs.SplitPath(libs.GetFullURL(r))
	if !ok {
		// 访问的根目录
		// 新用户
		_, redirectURL := newUser(curHost, cClientURLPostfix)
		redirect(redirectURL, w)
		return
	}

	// 获得 UUID 并统一为小写
	curUUID := strings.ToLower(paths[0])

	if len(curUUID) != cUUIDLength {
		// 不符合 UUID 规范，这应该是用户自己随便输入的
		// 新用户
		_, redirectURL := newUser(curHost, cClientURLPostfix)
		redirect(redirectURL, w)
		return
	}

	// 以下则是正确访问本系统的情况

	if connPool[curUUID] == nil {
		// 该 UUID 第一次访问系统
		// 用户第一次访问
		connPool[curUUID] = make(map[*websocket.Conn]bool)
	}
	// 老用户
	libs.RegTime(curUUID)
	// 完了后进行缓存清理
	defer cleanTimoutSession()

	// 处理不同端的访问
	if len(paths) > 1 {
		switch paths[1] {
		case cServerURLPostfix:
			// 服务器访问
			// 广播消息，不返回任何信息
			logrus.Info("Brodcast [" + curUUID + "] websocket msg")
			brodcastWS(curUUID, r)
		case cClientURLPostfix:
			// 客户端访问
			// 返回 html 页面
			logrus.Info(r.URL)
			showPage(curUUID, w, r)
		default:
			// 未知访问
			// 重定向到客户端访问
			tarURL := curHost + "/" + curUUID + "/" + cClientURLPostfix
			logrus.Info("Redirect to ", tarURL)
			redirect(tarURL, w)
		}
	}
}

// 新用户 生成 uuid 并返回，顺带返回跳转链接
func newUser(targetHost string, targetPostfix string) (
	tarUUID string, redirectURL string) {
	// 生成 uuid 并返回 uuid 和 跳转链接
	curUUID, err := uuid.NewV4()
	logrus.Error("Create uuid failed: ", err)
	return curUUID.String(), targetHost + "/" + curUUID.String() + "/" + targetPostfix
}

// 负责显示页面
func showPage(curUUID string, w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(wsTplFile.FilePath())
	if err != nil {
		logrus.Fatal("Parse template file "+wsTplFile.FilePath()+" err: ", err)
	}

	tarURL := libs.GetFullHost(r) + "/" + curUUID + "/" + cServerURLPostfix
	err = t.Execute(w, makeData(r, map[string]interface{}{"apiURL": tarURL}))
	if err != nil {
		logrus.Fatal("Execute template file "+wsTplFile.FilePath()+" err: ", err)
	}
	logrus.Info(r.URL)
}

// 负责在页面不正确时重定向
func redirect(tarURL string, w http.ResponseWriter) {
	w.Write([]byte("<head><meta http-equiv=\"refresh\" content=\"0;URL=" +
		tarURL + "\" /></head>"))
}

func brodcastWS(tarUUID string, r *http.Request) {
	data := makeData(r)
	wsConns := connPool[tarUUID]
	if wsConns == nil {
		return
	}

	// 如果该 uuid 上没有访问链接则直接返回
	if len(wsConns) == 0 {
		return
	}
	// 存在注册过的 ws 链接
	for key := range connPool[tarUUID] {
		err := websocket.JSON.Send(key, data)
		if err != nil {
			logrus.Error("Wrong ws connect", err)
			delete(connPool[tarUUID], key)
		}
	}
}

// 遍历清除超时的 session
func cleanTimoutSession() {
	for tmpUUID, Conns := range connPool {
		if libs.IsTimeout(tmpUUID, cTimeout) {
			for tmpWSConn := range Conns {
				if err := tmpWSConn.Close(); err != nil {
					logrus.Error("Close ws err: ", err)
				}
			}
			delete(connPool, tmpUUID)
			// 顺带清除定时器
			libs.DeleteTimerRecord(tmpUUID)
		}
	}
}
