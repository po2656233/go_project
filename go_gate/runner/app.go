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
		logPrefix = ""
		fileWriter = &log.FileWriter{
			RootDir:     "./logs/",     //日志根目录
			DirFormat:   "",             //日志根目录下无子目录
			FileFormat:  "20060102.log", //日志文件命名规则，按天切割文件
			TimeBegin:   len(logPrefix), //解析日志中时间起始位置，用于目录、文件切割，以免日志生成的地方所用时间与logfile写入时间不一致导致的切割偏差
			TimePrefix:  "2006-01-02 15:04:05.000",     //解析日志中时间格式
			MaxFileSize: 0,              //单个日志文件最大size，0则不限制size
			EnableBufio: false,          //是否开启bufio
		}
	)
	log.SetLevel(log.LEVEL_WARN)
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
