package runner

import (
	"encoding/xml"
	"flag"
	"github.com/nothollyhigh/kiss/log"
	 "go_gate/config"
	. "go_gate/manger"
	"go_gate/manger/proxy"
	"io"
	"io/ioutil"
	"os"
	"time"
)

var (
	appVersion = ""
	//bornTime   = time.Now() /* 进程启动时间 */
	confpath   = flag.String("config", "./config.xml", "config file path, default is ./config.xml")

	logout = io.Writer(nil)
)

func initConfig() {
	flag.Parse()

	filename := *confpath

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error when Open xml config file: %s: %v\n", filename, err)
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("Error when Read xml config file: %s: %v\n", filename, err)
	}

	err = xml.Unmarshal(data, config.GlobalXmlConfig)
	if err != nil {
		log.Fatal("Error when xml.Unmarshal from xml config file: %s: %v\n    data: %s\n", filename, err, string(data))
	}

	log.Info("config: %v\n%v ", filename, string(data))
}

func initLog() {
	var (
		fileWriter = &log.FileWriter{
			RootDir:     config.GlobalXmlConfig.Options.LogDir + time.Now().Format("20060102150405/"),
			DirFormat:   "",
			FileFormat:  "20060102.log",
			MaxFileSize: 1024 * 1024 * 32,
			EnableBufio: false,
		}
	)
	if config.GlobalXmlConfig.Options.Debug {
		logout = io.MultiWriter(os.Stdout, fileWriter)
	} else {
		logout = fileWriter
	}

	log.SetOutput(logout)
}


func Run(version string) {
	//版本号
	appVersion = version

	//初始化配置信息
	initConfig()

	//初始化日志信息
	initLog()
	log.Info("gate start, app version: '%v'", version)


	//初始化代理信息
	ProxyMgr.InitProxy()

	//初始化接口配置
	InitApi()

	// 每分钟统计一次连接数情况
	proxy.ConnMgr.StartDataFlowRecord(time.Second * 60)



	////性能测试 cpu
	//cpuProfile, _ := os.Create("cpu_profile")
	//pprof.StartCPUProfile(cpuProfile)
	//defer pprof.StopCPUProfile()
	//
	////性能测试 内存
	//memProfile, _ := os.Create("mem_profile")
	//pprof.WriteHeapProfile(memProfile)
}

func Stop() {
	log.Info("gate stop")
}
