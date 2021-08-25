package manger

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/nothollyhigh/kiss/log"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"go_gate/config"
	. "go_gate/manger/proxy"
	"io/ioutil"
	shnet "net"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type httpHandler struct {
	Register   string // 注册服务
	Query      string
	Reload     string
	Enableline string
}

// 获取服务状态信息
type StatusServer struct {
	Percent  StatusPercent
	CPU      []CPUInfo
	Mem      MemInfo
	Swap     SwapInfo
	Load     *load.AvgStat
	Network  map[string]InterfaceInfo
	BootTime uint64
	Uptime   uint64
}

// 获取利用率
type StatusPercent struct {
	CPU  float64
	Disk float64
	Mem  float64
	Swap float64
}

// 获取CPU信息
type CPUInfo struct {
	ModelName string
	Cores     int32
}

// 获取内存信息
type MemInfo struct {
	Total     uint64
	Used      uint64
	Available uint64
}

// 获取系统分页空间信息system paging space information
type SwapInfo struct {
	Total     uint64
	Used      uint64
	Available uint64
}

// 获取接口信息
type InterfaceInfo struct {
	Addrs    []string
	ByteSent uint64
	ByteRecv uint64
}

func SaveConfig() bool {
	// 保存当前配置
	oldPath := fmt.Sprintf("config_%v.xml", time.Now().Unix())
	os.Rename("config.xml", oldPath)
	newConfig, err := xml.MarshalIndent(config.GlobalXmlConfig, "", "\t")
	if err != nil {
		log.Info("SaveConfig finish! but not save err:%v", err)
		return false
	}
	log.Info("SaveConfig finish! ok")
	//清空后写入 ModeAppend 也会清空
	ioutil.WriteFile("config.xml", newConfig, os.ModeAppend)
	return true
}

func (h *httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//权限校验
	checkAuth := func() bool {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			w.Header().Set("WWW-Authenticate", `Basic realm="Dotcoo User Login"`)
			return false
		}
		auths := strings.SplitN(auth, " ", 2)
		if len(auths) != 2 {
			fmt.Println("error")
			return false
		}
		authMethod := auths[0]
		authB64 := auths[1]
		switch authMethod {
		case "Basic":
			authstr, err := base64.StdEncoding.DecodeString(authB64)
			if err != nil {
				fmt.Println(err)
				return false
			}
			log.Info("鉴权成功:%s", authstr)
			userPwd := strings.SplitN(string(authstr), ":", 2)
			if len(userPwd) != 2 {
				fmt.Println("error")
				return false
			}
			username := userPwd[0]
			password := userPwd[1]
			// 先写死了
			if username != "##sss^^^" || password != "(S?SS&^.14" {
				fmt.Println("Username:", username)
				fmt.Println("Password:", password)
				return false
			}
		default:
			fmt.Println("error")
			return false
		}
		return true
	}

	reload := func() {
		var readData []byte
		var err error
		if readData, err = ioutil.ReadAll(r.Body); err != nil {
			w.Write([]byte("please check if the document(config.xml) is valid."))
			log.Error("Reload Error when xml.Unmarshal from xml config file:len:%v, data:%v", r.ContentLength, err.Error())
			return
		}
		if err = xml.Unmarshal(readData, config.GlobalXmlConfig); err != nil {
			w.Write([]byte("unable to parse the XML file."))
			log.Error("Reload Error unable to parse the XML file", err.Error())
			return
		}

		//log.Info("HTTP API:Read data line ok:\n%v", string(readData))
		//节点信息比对
		var busLine *config.XMLBusLine
		var line *config.XMLLine
		var pLine *Line
		var proxy IProxy
		var node config.XMLNode
		confProxy := config.GlobalXmlConfig.Proxy
		options := config.GlobalXmlConfig.Options
		DEFAULT_TCP_REDIRECT = options.Redirect
		DEFAULT_TCP_CHECKLINE_INTERVAL = time.Second * time.Duration(options.Heartbeat.Interval)
		DEFAULT_TCP_CHECKLINE_TIMEOUT = time.Second * time.Duration(options.Heartbeat.Timeout)

		for _, busLine = range confProxy.BusLines { //[0
			if proxy, _ = ProxyMgr.GetProxy(busLine.Name); nil != proxy { //[1 获取到的值不能为空，否则新增处理
				// 查找线路
				for _, line = range busLine.Lines { //[2
					for _, node = range line.Nodes { //[3
						node.Addr = fmt.Sprintf("%s:%s", node.Ip, node.Port)
						if pLine = proxy.GetLine(line.ServerID, node.Addr); pLine != nil {
							continue
						}
						// 新增线路
						switch busLine.Type { //[4
						case PT_TCP:
							proxy.(*ProxyTcp).AddLine(line.ServerID, node.Addr, DEFAULT_TCP_CHECKLINE_TIMEOUT, DEFAULT_TCP_CHECKLINE_INTERVAL, node.Maxload, config.GlobalXmlConfig.Options.Redirect)
						case PT_WEBSOCKET:
							proxy.(*ProxyWebsocket).AddLine(line.ServerID, node.Addr, DEFAULT_TCP_CHECKLINE_TIMEOUT, DEFAULT_TCP_CHECKLINE_INTERVAL, node.Maxload, config.GlobalXmlConfig.Options.Redirect)
						} //4]
						// 日志
						log.Info("Reload id:%v addr:%v", line.ServerID, node.Addr)
					} //3]
				} //2]
				ProxyMgr.Proxys[busLine.Name] = proxy
				//保留有效线路
				proxy.ReserveLines(busLine.Lines)

				switch busLine.Type { //[2'
				case PT_TCP:
					proxy.(*ProxyTcp).StartCheckLines()
				case PT_WEBSOCKET:
					proxy.(*ProxyWebsocket).StartCheckLines()
				} //2']
			} else { //1] [1'
				// 新增 busline
				ProxyMgr.AddProxy(busLine)
			} //1']
		} //0]
		if SaveConfig() {
			w.Write([]byte("update config finish! ok!"))
		} else {
			w.Write([]byte("warning: update config finish! But cannot save config!"))
		}
	}

	register := func() {

		if err := r.ParseMultipartForm(32 << 20); err != nil {
			w.Write([]byte("warning: Parameter is empty!"))
			return
		}
		strType := r.Form.Get("type")
		strServerName := r.Form.Get("name")
		strIp := r.Form.Get("ip")
		strPort := r.Form.Get("port")
		strMaxLoad := r.Form.Get("maxload")
		address := shnet.ParseIP(strIp)
		maxload, _ := strconv.ParseInt(strMaxLoad, 10, 64)
		_, err1 := strconv.Atoi(strPort)
		if strType == "" || strServerName == "" || address == nil || err1 != nil || maxload <= 0 {
			w.Write([]byte("warning: Parameter is invalid!"))
			return
		}
		isOk := false
		for _, busLine := range config.GlobalXmlConfig.Proxy.BusLines { //[0
			if busLine.Type == strType {
				if proxy, _ := ProxyMgr.GetProxy(busLine.Name); nil != proxy { //[1 获取到的值不能为空，否则新增处理
					// 查找线路
					var pline *config.XMLLine = nil
					addr := fmt.Sprintf("%s:%s", strIp, strPort)
					for _, line := range busLine.Lines { //[2
						pLine := proxy.GetLine(line.ServerID, addr)
						if pLine != nil && pLine.Remote == addr {
							isOk = true
							break
						}
						if line.ServerID == strServerName {
							if pLine != nil {
								isOk = true
								break
							}
							pline = line
						}
					} //2]
					// 已经有了的,就不让注册了
					if isOk {
						w.Write([]byte("warning: the service already exists! register"))
						return
					}

					// 新增线路
					switch busLine.Type { //[4
					case PT_TCP:
						proxy.(*ProxyTcp).AddLine(strServerName, addr, DEFAULT_TCP_CHECKLINE_TIMEOUT, DEFAULT_TCP_CHECKLINE_INTERVAL, maxload, config.GlobalXmlConfig.Options.Redirect)
					case PT_WEBSOCKET:
						proxy.(*ProxyWebsocket).AddLine(strServerName, addr, DEFAULT_TCP_CHECKLINE_TIMEOUT, DEFAULT_TCP_CHECKLINE_INTERVAL, maxload, config.GlobalXmlConfig.Options.Redirect)
					} //4]
					// 日志
					log.Info("register serverId:%v addr:%v", strServerName, addr)
					switch busLine.Type { //[2'
					case PT_TCP:
						proxy.(*ProxyTcp).StartCheckLines()
					case PT_WEBSOCKET:
						proxy.(*ProxyWebsocket).StartCheckLines()
					} //2']
					if pline == nil {
						pline = &config.XMLLine{
							Addr:     busLine.Addr,
							ServerID: strServerName,
							Type:     strType,
							Nodes:    make([]config.XMLNode, 0),
						}
						busLine.Lines = append(busLine.Lines, pline)
					}
					pline.Nodes = append(pline.Nodes, config.XMLNode{
						Ip:      strIp,
						Port:    strPort,
						Maxload: maxload,
						Enable:  true,
					})
					if strType == "websocket" {
						pline.RealIpMode = "http"
					} else {
						pline.RealIpMode = "tcp"
					}

					isOk = true
				}
			}
		}
		// 开一条线
		if !isOk {
			log.Info("register type:%v serverId:%v addr:%v port:%v maxload:%v  failed!", strType, strServerName, strIp, strPort, maxload)
			w.Write([]byte("register failed! "))
		} else {
			log.Info("register type:%v serverId:%v addr:%v port:%v maxload:%v   successful!", strType, strServerName, strIp, strPort, maxload)
			w.Write([]byte("register successful! "))
			if !SaveConfig() {
				w.Write([]byte("warning: update config finish! But cannot save config!"))
			}
		}
		return
	}

	//查询信息
	query := func() {
		w.Write(getInfosJSON())
		for _, proxy := range ProxyMgr.Proxys {
			w.Write(proxy.LinesForJSON())
		}
		w.Write([]byte("\n"))
		w.Write([]byte(ConnMgr.LogDataFlowRecord()))
	}

	//启用
	enable := func() {
		//格式:<enable name=  serverID=  ip=  port=  enable=\>
		if r.Method != "POST" {
			fmt.Fprintf(w, "Invalid request mode.")
			return
		}
		var readData []byte
		var err error
		xmlControl := &config.XMLControl{}
		if readData, err = ioutil.ReadAll(r.Body); err != nil {
			fmt.Fprintf(w, "please check if the document(control.xml) is valid.")
			log.Error("Reload Error when xml.Unmarshal from xml config file:len:%v, data:%v", r.ContentLength, err.Error())
			return
		}
		if err = xml.Unmarshal(readData, xmlControl); err != nil {
			w.Write([]byte("unable to parse the control.xml!"))
			log.Error("Reload Error unable to parse the XML file", err.Error())
			return
		}
		// 二次鉴权
		ss := md5.Sum([]byte("^^*SDASD)A)$%"))
		if 0 != strings.Compare(fmt.Sprintf("%X", ss), xmlControl.PassWord) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		var proxy IProxy
		var addr string
		var pLine *Line
		// 启用\停用列表
		for _, enable := range xmlControl.Enables {
			if proxy, _ = ProxyMgr.GetProxy(enable.Name); nil != proxy { //[1 获取到的值不能为空，否则新增处理
				addr = fmt.Sprintf("%s:%s", enable.IP, enable.Port)
				addrs, _ := shnet.LookupHost(addr)
				if 0 < len(addrs) {
					addr = addrs[0]
				}
				// 查找线路
				if pLine = proxy.GetLine(enable.ID, addr); nil != pLine {
					if enable.Enable {
						pLine.UnPause()
					} else {
						pLine.Pause()
					}
				}

			}
		}
		fmt.Fprintf(w, "finish-> %v", xmlControl.Enables)
	}

	//鉴权
	if !checkAuth() {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// 接口解析调用
	switch r.Method {
	case "GET":
		switch r.URL.Path {
		case h.Query:
			query()
			log.Info("HTTP API:query ok.")
		default:
			http.NotFound(w, r)
		}
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		switch r.URL.Path {
		case h.Register:
			register()
			log.Info("HTTP API:register server ok.")
		case h.Reload:
			reload()
			log.Info("HTTP API:Reload ok.")
		case h.Enableline:
			enable()
			log.Info("HTTP API:enable line ok.")
		}

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		http.NotFound(w, r)
	}

}

func getInfosJSON() []byte {
	vm, _ := mem.VirtualMemory()
	sm, _ := mem.SwapMemory()
	cpuStat, _ := cpu.Info()
	cc, _ := cpu.Percent(time.Second, true)
	du, _ := disk.Usage("/")
	hi, _ := host.Info()
	nv, _ := net.IOCounters(true)
	la, _ := load.Avg()
	interStat, _ := net.Interfaces()
	ss := new(StatusServer)
	ss.Load = la
	ss.Uptime = hi.Uptime
	ss.BootTime = hi.BootTime
	ss.Percent.Mem = vm.UsedPercent
	ss.Percent.CPU = cc[0]
	ss.Percent.Swap = sm.UsedPercent
	ss.Percent.Disk = du.UsedPercent
	ss.CPU = make([]CPUInfo, len(cpuStat))
	for index, ci := range cpuStat {
		ss.CPU[index].ModelName = ci.ModelName
		ss.CPU[index].Cores = ci.Cores
	}
	ss.Mem.Total = vm.Total
	ss.Mem.Available = vm.Available
	ss.Mem.Used = vm.Used
	ss.Swap.Total = sm.Total
	ss.Swap.Available = sm.Free
	ss.Swap.Used = sm.Used
	ss.Network = make(map[string]InterfaceInfo)
	for _, v := range nv {
		var ii InterfaceInfo
		ii.ByteSent = v.BytesSent
		ii.ByteRecv = v.BytesRecv
		ss.Network[v.Name] = ii
	}
	for _, v := range interStat {
		if ii, ok := ss.Network[v.Name]; ok {
			ii.Addrs = make([]string, len(v.Addrs))
			for index, vv := range v.Addrs {
				ii.Addrs[index] = vv.Addr
			}
			ss.Network[v.Name] = ii
		}
	}
	b, err := json.Marshal(ss)
	if err != nil {
		log.Error("infos to JSON error:%v", err.Error())
		return nil
	}
	return b
}

func info(i interface{}) {
	t := reflect.TypeOf(i)
	fmt.Println("Type: ", t.Name())
	v := reflect.ValueOf(i)
	fmt.Println("Fields: ")
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}
}

func InitApi() {
	mux := http.NewServeMux()
	hh := &httpHandler{
		Register:   config.GlobalXmlConfig.HttpApi.RegisterPath,
		Query:      config.GlobalXmlConfig.HttpApi.QueryPath,
		Reload:     config.GlobalXmlConfig.HttpApi.ReloadPath,
		Enableline: config.GlobalXmlConfig.HttpApi.EnablePath,
	}

	v := reflect.ValueOf(*hh)

	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		switch f.Kind() {
		case reflect.String:
			mux.Handle(f.String(), hh)
		}
	}

	go func() {
		if err := http.ListenAndServe(config.GlobalXmlConfig.HttpApi.Addr, mux); nil != err {
			log.Fatal("api err:%v", err.Error())
		}
	}()

}
