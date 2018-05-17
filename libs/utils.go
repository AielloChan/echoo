package libs

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
)

var (
	_timer        = make(map[string]int64)
	_SplitPathReg = regexp.MustCompile(`https?://[^\/]+\/(?:([^\/\?]+)\/)*([^\?]*)`)
)

// GetFullURL 获得完整 url
func GetFullURL(r *http.Request) string {
	return GetFullHost(r) + r.RequestURI
}

// GetFullHost 获得完整 host
func GetFullHost(r *http.Request) string {
	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}
	return strings.Join([]string{scheme, r.Host}, "")
}

// ConcatMap 拼接两个 map
func ConcatMap(m1 []string, m2 []string) []string {
	for k, v := range m1 {
		m2[k] = v
	}
	return m2
}

// ReadAll 读取所有并转为 string 输出
func ReadAll(r io.Reader) string {
	content, err := ioutil.ReadAll(r)
	logrus.Debug("Read body data failed: ", err)
	return string(content)
}

// FmtCookies 格式化 Cookie
func FmtCookies(cookies []*http.Cookie) []string {
	str := []string{}
	for _, v := range cookies {
		str = append(str, fmt.Sprint(v))
	}
	return str
}

// StringReducer 数据结构调整
func StringReducer(stringableItems []interface{ String() string }) []string {
	strItems := []string{}
	for index, val := range stringableItems {
		strItems[index] = val.String()
	}
	return strItems
}

// MapReducer 数据结构调整
func MapReducer(mapData map[string][]string) map[string]string {
	newMap := map[string]string{}
	for key, val := range mapData {
		newMap[key] = strings.Join(val, ",")
	}
	return newMap
}

// PrepareFile 准备文件
func PrepareFile(filePath string) {
	// check
	if _, err := os.Stat(filePath); err != nil {
		err := os.MkdirAll(path.Dir(filePath), 0711)
		if err != nil {
			logrus.Fatal("Can't create dir for logging: ", err)
		}
	}

	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		logrus.Fatal("Can't write to log file: ", err)
	}
	defer f.Close()
}

// RegTime 登记计时器
func RegTime(uid string) {
	_timer[uid] = time.Now().Unix()
}

// IsTimeout 判断超时
func IsTimeout(uid string, timelimit int64) bool {
	if timstamp, ok := _timer[uid]; ok {
		timstampNow := time.Now().Unix()
		if timstampNow-timstamp > timelimit {
			return true
		}
		return false
	}
	return true
}

// DeleteTimerRecord 计时器垃圾收集
func DeleteTimerRecord(id string) {
	delete(_timer, id)
}

// SplitPath 获得 URL 中的地址项
// 如 http://localhost/main/secendary/.. 则会获取并返回 "main" "secendary" 等
func SplitPath(str string) ([]string, bool) {
	if len(str) == 0 {
		return nil, false
	}
	results := _SplitPathReg.FindStringSubmatch(str)
	if len(results) > 1 {
		return results[1:], true
	}
	return nil, false
}
