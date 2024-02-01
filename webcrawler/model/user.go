package model

import (
	"errors"
	"fmt"
	"github.com/anaskhan96/soup"
	"net/http"
	"net/http/cookiejar"
	"regexp"
	"strings"
	"time"
)

type User struct {
	Number   string `json:"number" form:"number" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func GetCookieAndAuth(number, password string) (string, string) {
	ji, rp := NewCCNUClient(number, password)
	pt, _ := Index(ji, rp)

	autho := "Bearer " + pt

	cookie := "JSESSIONID=" + ji + "; routeportal=" + rp + "; PORTAL_TOKEN=" + pt

	return cookie, autho
}

func NewCCNUClient(studentID string, password string) (string, string) {
	htmlBody, _ := soup.Get("https://account.ccnu.edu.cn/cas/login?service=http%3A%2F%2Fone.ccnu.edu.cn%2Fcas%2Flogin_portal")
	doc := soup.HTMLParse(htmlBody)

	links1 := doc.Find("body", "id", "cas").FindAll("script")
	js := links1[2].Attrs()["src"][26:]

	links2 := doc.Find("div", "class", "logo").FindAll("input")
	st := links2[2].Attrs()["value"]

	jar, _ := cookiejar.New(&cookiejar.Options{})
	client := &http.Client{
		Jar:     jar,
		Timeout: 5 * time.Second,
	}

	url := fmt.Sprintf("https://account.ccnu.edu.cn/cas/login;jsessionid=%v?service=http", js) + "%3A%2F%2Fone.ccnu.edu.cn%2Fcas%2Flogin_portal"
	text := fmt.Sprintf("username=%v&password=%v&lt=%v&execution=e1s1&_eventId=submit&submit=", studentID, password, st) + "%E7%99%BB%E5%BD%95"
	body := strings.NewReader(text)

	req, _ := http.NewRequest("POST", url, body)

	req.Header.Set("Cookie", "JSESSIONID="+js)
	req.Header.Set("Host", "account.ccnu.edu.cn")
	req.Header.Set("Origin", "https://account.ccnu.edu.cn")
	req.Header.Set("Referer", "https://account.ccnu.edu.cn/cas/login?service=http%3A%2F%2Fone.ccnu.edu.cn%2Fcas%2Flogin_portal")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return "", ""
	}

	cookies := client.Jar.Cookies(resp.Request.URL)
	ji := cookies[0].Value //JSESSIONID
	rp := cookies[1].Value //routeportal

	return ji, rp
}

func Index(ji, rp string) (string, error) {
	// 创建包含Cookie的请求
	req, err := http.NewRequest("GET", "http://one.ccnu.edu.cn/index", nil)
	if err != nil {
		return "", err
	}

	cookie := "JSESSIONID=" + ji + "; " + "routeportal=" + rp

	// 设置请求头中的Cookie
	//JSESSIONID=08997B7DB45C90356905AFE484C1184B; routeportal=5fb17c08d7829466342e4fe60bd21884
	req.Header.Set("Cookie", cookie)

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 处理响应...
	pt := resp.Header.Get("Set-Cookie")

	// 使用正则表达式提取PORTAL_TOKEN
	re := regexp.MustCompile(`PORTAL_TOKEN=([^\s]+)`)
	match := re.FindStringSubmatch(pt)

	if len(match) >= 2 {
		pt = match[1]
		return pt, nil
	} else {
		return "", errors.New("未找到PORTAL_TOKEN")
	}
}
