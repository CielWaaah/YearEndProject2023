package main

import (
	"year_end_project/webcrawler/conf"
	"year_end_project/webcrawler/routers"
)

func main() {
	conf.Init()
	//初始化配置

	r := routers.NewRouter()

	r.Run()

}
