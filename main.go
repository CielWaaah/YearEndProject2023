package main

import "year_end_project/webcrawler/routers"

func main() {
	//初始化配置

	r := routers.NewRouter()

	r.Run()

}
