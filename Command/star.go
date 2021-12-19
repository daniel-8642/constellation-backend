package Command

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

//作为缓存
var responseBuffer = make(map[string][]byte, 15)
var timeOut time.Time

func Star(c *gin.Context) {
	//c.Header("Access-Control-Allow-Origin","*")
	consName := c.DefaultQuery("consName", "0") //查询关键字
	if consName == "0" {
		//todo:错误处理
	}
	type_ := "today" // c.DefaultQuery("type", "today")由于缓存原因,仅允许查询今日运势
	key := c.DefaultQuery("key", "e1f7fff20b301745c64b655e0ef231d7")
	if timeOut.Format("2006-01-02") != time.Now().Format("2006-01-02") {
		responseBuffer = make(map[string][]byte, 15)
		timeOut = time.Now()
	}
	var body []byte
	if _, ok := responseBuffer[consName]; !ok {
		body = netQuery(consName, type_, key)
	} else {
		body, _ = responseBuffer[consName]
	}
	ret := fmt.Sprintf("%s", body)
	fmt.Println(ret)
	if strings.Contains(ret, "\"resultcode\":\"200\"") {
		c.Set("log", true)
	} else {
		c.Set("log", false)
	}
	c.String(http.StatusOK, ret) //输出
}

func netQuery(consName string, type_ string, key string) []byte {
	fmt.Println("击穿缓存,联网查询")
	resp, err := http.Get("http://web.juhe.cn:8080/constellation/getAll?" +
		"consName=" + url.QueryEscape(consName) + "&" +
		"type=" + type_ + "&" +
		"key=" + key)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	body, _ := ioutil.ReadAll(resp.Body)
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	if strings.Contains(string(body), "\"resultcode\":\"200\",\"error_code\":0") {
		responseBuffer[consName] = body
	}
	return body
}
