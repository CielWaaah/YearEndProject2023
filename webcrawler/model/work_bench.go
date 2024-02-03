package model

import (
	"fmt"
	"io"
	"net/http"
	"year_end_project/webcrawler/conf"
)

type WorkBench struct {
	//
}

func GetMemberInfo() (string, error) {
	req, err := http.NewRequest("GET", "http://work.muxi-tech.xyz/api/v1.0/group/0/userList/", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Cookie", conf.WCookie)
	req.Header.Set("Token", conf.WToken)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func GetUserArticles(uid string) (string, error) {
	url := fmt.Sprintf("http://work.muxi-tech.xyz/api/v1.0/status/%v/list/1/", uid)
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Cookie", conf.WCookie)
	req.Header.Set("Token", conf.WToken)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil

}
