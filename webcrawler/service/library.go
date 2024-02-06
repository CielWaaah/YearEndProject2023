package service

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"io"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"time"
)

type LibService struct {
	//
}

// 获取SessionId
func (service *LibService) NewLibrayClient(studentID string, password string) string {
	htmlBody, _ := soup.Get("https://account.ccnu.edu.cn/cas/login?service=http://kjyy.ccnu.edu.cn/loginall.aspx?page=")
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

	url := fmt.Sprintf("https://account.ccnu.edu.cn/cas/login;jsessionid=%v?service=http://kjyy.ccnu.edu.cn/loginall.aspx?page=", js)
	text := fmt.Sprintf("username=%v&password=%v&lt=%v&execution=e1s1&_eventId=submit&submit=", studentID, password, st) + "%E7%99%BB%E5%BD%95"
	body := strings.NewReader(text)

	req, _ := http.NewRequest("POST", url, body)

	req.Header.Set("Cookie", "JSESSIONID="+js)
	req.Header.Set("Host", "account.ccnu.edu.cn")
	req.Header.Set("Origin", "https://account.ccnu.edu.cn")
	req.Header.Set("Referer", "https://account.ccnu.edu.cn/cas/login?service=http://kjyy.ccnu.edu.cn/loginall.aspx?page=")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	cookies := client.Jar.Cookies(resp.Request.URL)

	s := cookies[0].Value //SessionId
	si := "ASP.NET_SessionId=" + s
	return si
}

func (service *LibService) GetLibraryRecords(cookie string) (string, error) {
	url := "http://kjyy.ccnu.edu.cn/ClientWeb/pro/ajax/center.aspx?act=get_History_resv&strat=90&StatFlag=OVER"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("Cookie", cookie)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return string(body), nil
}
