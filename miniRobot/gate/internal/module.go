package internal

import (
    "github.com/name5566/leaf/chanrpc"
    "github.com/name5566/leaf/log"
    "github.com/name5566/leaf/network"
    "math"
    "miniRobot/base"
    "miniRobot/conf"
    "miniRobot/game"
    "miniRobot/msg"
    "net"
    "reflect"
    "time"
)

type Module struct {
    *Gate
}

func (m *Module) OnInit() {
    m.Gate = &Gate{
        MaxConnNum:      conf.Server.MaxConnNum,
        PendingWriteNum: conf.PendingWriteNum,
        MaxMsgLen:       conf.MaxMsgLen,
        WSAddr:          conf.Server.WSAddr,
        HTTPTimeout:     conf.HTTPTimeout,
        CertFile:        conf.Server.CertFile,
        KeyFile:         conf.Server.KeyFile,
        TCPAddr:         conf.Server.TCPAddr,
        LenMsgLen:       conf.LenMsgLen,
        LittleEndian:    conf.LittleEndian,
        Processor:       msg.ProcessorProto, //消息处理器对象(proto|json)
        AgentChanRPC:    game.ChanRPC,
    }
}

type Gate struct {
    MaxConnNum      int
    PendingWriteNum int
    MaxMsgLen       uint32
    Processor       network.Processor
    AgentChanRPC    *chanrpc.Server

    // websocket
    WSAddr      string
    HTTPTimeout time.Duration
    CertFile    string
    KeyFile     string

    // tcp
    TCPAddr      string
    LenMsgLen    int
    LittleEndian bool
}

func (gate *Gate) Run(closeSig chan bool) {
    //var wsServer *network.WSClient
    //if gate.WSAddr != "" {
    //	wsServer = new(network.WSClient)
    //	wsServer.Addr = gate.WSAddr
    //	//wsServer.MaxConnNum = gate.MaxConnNum
    //	wsServer.PendingWriteNum = gate.PendingWriteNum
    //	wsServer.MaxMsgLen = gate.MaxMsgLen
    //	//wsServer.HTTPTimeout = gate.HTTPTimeout
    //	//wsServer.CertFile = gate.CertFile
    //	//wsServer.KeyFile = gate.KeyFile
    //	wsServer.NewAgent = func(conn *network.WSConn) network.Agent {
    //		a := &agent{conn: conn, gate: gate}
    //		if gate.AgentChanRPC != nil {
    //			gate.AgentChanRPC.Go("NewAgent", a)
    //		}
    //		return a
    //	}
    //}
    listWS := make( []*network.WSClient,0)
    for i := 0; i < base.ALLCount; i++ {

        var client *network.WSClient
        if gate.WSAddr != "" {
            client = new(network.WSClient)
            client.Addr = gate.WSAddr
            client.ConnNum = 1
            client.HandshakeTimeout = 10 * time.Second
            client.ConnectInterval = 3 * time.Second
            client.PendingWriteNum = conf.PendingWriteNum
            //client.LenMsgLen = 4
            client.AutoReconnect = true
            client.MaxMsgLen = math.MaxUint32
            client.NewAgent = func(conn *network.WSConn) network.Agent {
                a := &agent{conn: conn, gate: gate}
                if gate.AgentChanRPC != nil {
                    gate.AgentChanRPC.Go("NewAgent", a)
                }
                return a
            }
        }

        //if wsServer != nil {
        //	wsServer.Start()
        //}
        if client != nil {
            client.Start()
        }
        listWS = append(listWS,client)

        //if wsServer != nil {
        //	wsServer.Close()
        //}

    }
    //<-closeSig
    //for _,client:=range listWS{
    //    if client != nil {
    //        client.Close()
    //    }
    //}


}

func (gate *Gate) OnDestroy() {}

type agent struct {
    conn     network.Conn
    gate     *Gate
    userData interface{}
}

func (a *agent) Run() {
    for {
        data, err := a.conn.ReadMsg()
        if err != nil {
            log.Debug("read message: %v", err)
            break
        }

        if a.gate.Processor != nil {
            msg, err := a.gate.Processor.Unmarshal(data)
            if err != nil {
                log.Debug("unmarshal message error: %v", err)
                break
            }
            err = a.gate.Processor.Route(msg, a)
            if err != nil {
                log.Debug("route message error: %v", err)
                break
            }
        }
    }
}

func (a *agent) OnClose() {
    if a.gate.AgentChanRPC != nil {
        err := a.gate.AgentChanRPC.Call0("CloseAgent", a)
        if err != nil {
            log.Error("chanrpc error: %v", err)
        }
    }
}

func (a *agent) WriteMsg(msg interface{}) {
    if a.gate.Processor != nil {
        data, err := a.gate.Processor.Marshal(msg)
        if err != nil {
            log.Error("marshal message %v error: %v", reflect.TypeOf(msg), err)
            return
        }
        err = a.conn.WriteMsg(data...)
        if err != nil {
            log.Error("write message %v error: %v", reflect.TypeOf(msg), err)
        }
    }
}

func (a *agent) LocalAddr() net.Addr {
    return a.conn.LocalAddr()
}

func (a *agent) RemoteAddr() net.Addr {
    return a.conn.RemoteAddr()
}

func (a *agent) Close() {
    a.conn.Close()
}

func (a *agent) Destroy() {
    a.conn.Destroy()
}

func (a *agent) UserData() interface{} {
    return a.userData
}

func (a *agent) SetUserData(data interface{}) {
    a.userData = data
}
