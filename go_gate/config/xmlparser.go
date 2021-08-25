package config

import (
	"encoding/xml"
)

var (
	GlobalXmlConfig = &XMLConfig{}
)

///-----------config.xml(start)------------------///
//配置
type XMLConfig struct {
	XMLName xml.Name   `xml:"setting"`
	Options XMLOptions `xml:"options"`
	Proxy   XMLProxy   `xml:"proxy"`
	HttpApi XMLHttpApi `xml:"api"`
}

//设置项
type XMLOptions struct {
	XMLName   xml.Name     `xml:"options"`
	Debug     bool         `xml:"debug,attr"`
	LogDir    string       `xml:"logdir,attr"`
	Redirect  bool         `xml:"redirect,attr"`
	Heartbeat XMLHeartbeat `xml:"heartbeat"`
}

//心跳
type XMLHeartbeat struct {
	XMLName  xml.Name `xml:"heartbeat"`
	Interval int      `xml:"interval,attr"`
	Timeout  int      `xml:"timeout,attr"`
}

// HTTP API服务配置
type XMLHttpApi struct {
	XMLName      xml.Name `xml:"api"`
	Addr         string   `xml:"addr,attr"`
	Type         string   `xml:"type,attr"`
	RegisterPath string   `xml:"registerpath,attr"`
	QueryPath    string   `xml:"querypath,attr"`
	ReloadPath   string   `xml:"reloadpath,attr"`
	EnablePath   string   `xml:"enablepath,attr"`
	DisablePath  string   `xml:"disablepath,attr"`
}

//代理
type XMLProxy struct {
	XMLName  xml.Name      `xml:"proxy"`
	BusLines []*XMLBusLine `xml:"busline"`
	//Lines   []XMLLine `xml:"line"`
}

//---------------------------------------------//
//总线(相当于每条线路的默认配置)
type XMLBusLine struct {
	XMLName    xml.Name    `xml:"busline"`
	Name       string      `xml:"name,attr"`
	Addr       string      `xml:"addr,attr"`
	Type       string      `xml:"type,attr"`
	Redirect   string      `xml:"redirect,attr"`
	TLS        bool        `xml:"tls,attr"`
	RealIpMode string      `xml:"realipmode,attr"` //真实ip
	Routes     []*XMLRoute `xml:"route"`
	Certs      []*XMLCert  `xml:"cert,attr"`
	Lines      []*XMLLine  `xml:"line"`
}

//线路
type XMLLine struct {
	XMLName    xml.Name   `xml:"line"`
	ServerID   string     `xml:"serverid,attr"` //ServerID
	Addr       string     `xml:"-"`
	Type       string     `xml:"-"`
	Redirect   string     `xml:"-"`
	TLS        bool       `xml:"-"`
	RealIpMode string     `xml:"-"` //真实ip
	Routes     []XMLRoute `xml:"route"`
	Certs      []XMLCert  `xml:"cert"`
	Nodes      []XMLNode  `xml:"node"`
}

//节点
type XMLNode struct {
	XMLName xml.Name `xml:"node"`
	Addr    string   `xml:"-"`         // 实际使用的IP地址
	Ip      string   `xml:"ip,attr"`   // 配置上使用的IP
	Port    string   `xml:"port,attr"` // 配置上使用的端口
	Maxload int64    `xml:"maxload,attr"`
	Enable  bool     `xml:"enable,attr"`
	speed   int
}

//证书
type XMLCert struct {
	XMLName  xml.Name `xml:"cert"`
	Certfile string   `xml:"certfile,attr"`
	Keyfile  string   `xml:"keyfile,attr"`
}

//路由
type XMLRoute struct {
	XMLName xml.Name `xml:"route"`
	Path    string   `xml:"path,attr"`
}

///-----------config.xml(end)------------------///

///-----------control.xml(start)------------------///
// 控制
type XMLControl struct {
	XMLName  xml.Name    `xml:"control"`
	Name     string      `xml:"username"`
	PassWord string      `xml:"password"` //md5
	Enables  []XMLEnable `xml:"line"`
}

// 停用/启用
type XMLEnable struct {
	XMLName xml.Name `xml:"line"`
	Name    string   `xml:"name,attr"`
	ID      string   `xml:"serverID,attr"`
	IP      string   `xml:"ip,attr"`
	Port    string   `xml:"port,attr"`
	Enable  bool     `xml:"enable,attr"`
}

///-----------control.xml(end)------------------///
