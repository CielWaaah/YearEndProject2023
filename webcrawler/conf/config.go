package conf

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	WCookie string
	WToken  string
)

func LoadWorkBench(file *ini.File) {
	WCookie = file.Section("work_bench").Key("Cookie").String()
	WToken = file.Section("work_bench").Key("Token").String()
}

func Init() {
	file, err := ini.Load("./webcrawler/conf/config.ini")
	if err != nil {
		fmt.Println("配置文件读取失败", err)
	}

	LoadWorkBench(file)

	//SetupLogger()
}

//func SetupLogger() {
//	logFileLocation, _ := os.OpenFile("./year_end_project.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
//	log.SetOutput(logFileLocation)
//}
