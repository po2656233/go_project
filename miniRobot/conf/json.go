package conf

import (
	"encoding/json"
	"github.com/name5566/leaf/log"
	"io/ioutil"
)

var Server struct {
	LogLevel    string `json:"logLevel"`
	LogPath     string `json:"logPath"`
	WSAddr      string `json:"wsAddr"`
	CertFile    string
	KeyFile     string
	TCPAddr     string	`json:"tcpAddr"`
	MaxConnNum  int		`json:"maxConnNum"`
	ConsolePort int
	ProfilePath string `json:"profilePath"`
	CarryMainSubID bool  `json:"carryMainSubID"`//传输的时候 是否需要携带mainID 或 subID
}

func init() {
	data, err := ioutil.ReadFile("conf/server.json")
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &Server)
	if err != nil {
		log.Fatal("%v", err)
	}
}
