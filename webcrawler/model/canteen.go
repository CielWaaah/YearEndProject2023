package model

import (
	"io"
	"net/http"
	"strings"
)

type Canteen struct {
	//
}

func GetCanteenDataWithCookie(dataURL, authorization, cookie, startT, endT string) (string, error) {
	requestBody := "limit=10&page=1&tranType=&start=" + startT + "&end=" + endT

	// 创建包含Cookie的请求
	req, err := http.NewRequest("POST", dataURL, strings.NewReader(requestBody))
	if err != nil {
		return "", err
	}

	// 设置请求头中的Cookie
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", authorization)
	req.Header.Set("Cookie", cookie)

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 处理响应...
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
