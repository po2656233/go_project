package internal

import (
    "crypto/rand" //真随机
    "github.com/name5566/leaf/gate"
    "github.com/name5566/leaf/log"
    "math/big"
    . "miniRobot/base"
    protoMsg "miniRobot/msg/go"
    "reflect"
    "sort"
    "time"
)

var mjercards map[uint64][]int32
var mjcards map[uint64][]int32

//初始化
func init() {
    //游戏处理
    handlerMsg(&protoMsg.EnterGameResp{}, handleEnterGame)               //反馈--->主页信息
    handlerMsg(&protoMsg.EnterGameQZCCResp{}, handleEnterGameQZCC)       //反馈--->主页信息
    handlerMsg(&protoMsg.EnterGameTBCCResp{}, handleEnterGameTBCC)       //反馈--->主页信息
    handlerMsg(&protoMsg.EnterGameZJHResp{}, handleEnterGameZJH)         //反馈--->主页信息
    handlerMsg(&protoMsg.EnterGameMJResp{}, handleEnterGameMJ)           //反馈--->主页信息
    handlerMsg(&protoMsg.EnterGameMJERResp{}, handleEnterGameMJER)       //反馈--->主页信息
    handlerMsg(&protoMsg.EnterGameSGResp{}, handleEnterGameSG)           //反馈--->主页信息
    handlerMsg(&protoMsg.LandLordsSceneResp{}, handleEnterGameLandLords) //反馈--->主页信息
    handlerMsg(&protoMsg.TuitongziSceneResp{}, handleEnterGameTuitongzi) //反馈--->主页信息

    //准备
    handlerMsg(&protoMsg.QzcowcowStateFreeResp{}, handleQzcowcowStateFreeResp)   //反馈--->主页信息
    handlerMsg(&protoMsg.TbcowcowStateFreeResp{}, handleTbcowcowStateFreeResp)   //反馈--->主页信息
    handlerMsg(&protoMsg.MahjongStateFreeResp{}, handleMahjongStateFreeResp)     //反馈--->主页信息
    handlerMsg(&protoMsg.MahjongERStateFreeResp{}, handleMahjongERStateFreeResp) //反馈--->主页信息
    handlerMsg(&protoMsg.ZhajinhuaStateFreeResp{}, handleZhajinhuaStateFreeResp) //反馈--->主页信息
    handlerMsg(&protoMsg.SangongStateFreeResp{}, handleSangongStateFreeResp)     //反馈--->主页信息
    handlerMsg(&protoMsg.LandLordsStateFreeResp{}, handleLandLordsStateFreeResp) //反馈--->主页信息

    //发牌
    handlerMsg(&protoMsg.MahjongDealResp{}, handleMahjongDealResp)     //反馈--->主页信息
    handlerMsg(&protoMsg.MahjongERDealResp{}, handleMahjongERDealResp) //反馈--->主页信息

    //百人下注
    handlerMsg(&protoMsg.BaccaratStatePlayingResp{}, handleBaccaratStatePlayingResp)         //反馈--->主页信息
    handlerMsg(&protoMsg.BrcowcowStatePlayingResp{}, handleBrcowcowStatePlayingResp)         //反馈--->主页信息
    handlerMsg(&protoMsg.BrtoubaoStatePlayingResp{}, handleBrtoubaoStatePlayingResp)         //反馈--->主页信息
    handlerMsg(&protoMsg.BrTuitongziStatePlayingResp{}, handleBrTuitongziStatePlayingResp)   //反馈--->主页信息
    handlerMsg(&protoMsg.TigerXdragonStatePlayingResp{}, handleTigerXdragonStatePlayingResp) //反馈--->主页信息

    ///对战出牌
    handlerMsg(&protoMsg.ZhajinhuaStatePlayingResp{}, handleZhajinhuaStatePlayingResp) //反馈--->主页信息
    handlerMsg(&protoMsg.MahjongStatePlayingResp{}, handleMahjongStatePlayingResp)     //反馈--->主页信息
    handlerMsg(&protoMsg.MahjongERStatePlayingResp{}, handleMahjongERStatePlayingResp) //反馈--->主页信息
    handlerMsg(&protoMsg.LandLordsStatePlayingResp{}, handleLandLordsStatePlayingResp) //反馈--->主页信息
    handlerMsg(&protoMsg.TuitongziStatePlayingResp{}, handleTuitongziStatePlayingResp) //反馈--->主页信息

    //操作提示
    handlerMsg(&protoMsg.MahjongHintResp{}, handleMahjongHintResp)     //反馈--->主页信息
    handlerMsg(&protoMsg.MahjongERHintResp{}, handleMahjongERHintResp) //反馈--->主页信息
}

//注册传输消息
func handlerMsg(m interface{}, h interface{}) {
    skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

///////////////////////入场/////////////////////////////
//百人类
func handleEnterGame(args []interface{}) {
    //  m := args[0].(*protoMsg.EnterGameResp)
    //a := args[1].(gate.Agent)
    //  log.Debug("进入游戏:%v", m.UserInfo.UserID)

}

//抢庄牛牛
func handleEnterGameQZCC(args []interface{}) {
    m := args[0].(*protoMsg.EnterGameQZCCResp)
    a := args[1].(gate.Agent)
    person := a.UserData().(*protoMsg.UserInfo)
    if m.Player.MyInfo.UserID == person.UserID {
        log.Debug("进入游戏:%v", m)
        msg := &protoMsg.QzcowcowReadyReq{
            IsReady: true,
        }
        a.WriteMsg(msg)
    }

}

//
func handleEnterGameTBCC(args []interface{}) {
    m := args[0].(*protoMsg.EnterGameTBCCResp)
    a := args[1].(gate.Agent)
    person := a.UserData().(*protoMsg.UserInfo)
    if m.Player.MyInfo.UserID == person.UserID {
        log.Debug("进入游戏:%v", m)
        msg := &protoMsg.TbcowcowReadyReq{
            IsReady: true,
        }
        a.WriteMsg(msg)
    }

}

//
func handleEnterGameZJH(args []interface{}) {
    m := args[0].(*protoMsg.EnterGameZJHResp)
    a := args[1].(gate.Agent)
    person := a.UserData().(*protoMsg.UserInfo)
    if m.Player.MyInfo.UserID == person.UserID {
        log.Debug("进入游戏:%v", m)
        msg := &protoMsg.ZhajinhuaReadyReq{
            IsReady: true,
        }
        a.WriteMsg(msg)
    }

}
func handleEnterGameMJ(args []interface{}) {
    m := args[0].(*protoMsg.EnterGameMJResp)
    a := args[1].(gate.Agent)
    if len(mjcards) == 0 || mjcards == nil {
        mjcards = make(map[uint64][]int32, 0)
    }

    person := a.UserData().(*protoMsg.UserInfo)
    if m.Player.MyInfo.UserID == person.UserID {
        log.Debug("进入游戏:%v", m)
        msg := &protoMsg.MahjongReadyReq{
            IsReady: true,
        }
        mjcards[person.UserID] = make([]int32, 0)
        a.WriteMsg(msg)
    }

}

func handleEnterGameMJER(args []interface{}) {
    m := args[0].(*protoMsg.EnterGameMJERResp)
    a := args[1].(gate.Agent)
    if len(mjercards) == 0 || mjercards == nil {
        mjercards = make(map[uint64][]int32, 0)
    }
    person := a.UserData().(*protoMsg.UserInfo)
    if m.Player.MyInfo.UserID == person.UserID {
        log.Debug("进入游戏:%v", m)
        msg := &protoMsg.MahjongERReadyReq{
            IsReady: true,
        }
        mjercards[person.UserID] = make([]int32, 0)
        a.WriteMsg(msg)
    }

}

func handleEnterGameSG(args []interface{}) {
    m := args[0].(*protoMsg.EnterGameSGResp)
    a := args[1].(gate.Agent)
    person := a.UserData().(*protoMsg.UserInfo)
    if m.Player.MyInfo.UserID == person.UserID {
        log.Debug("进入游戏:%v", m)
        msg := &protoMsg.SangongReadyReq{
            IsReady: true,
        }
        a.WriteMsg(msg)
    }

}

func handleEnterGameLandLords(args []interface{}) {
    m := args[0].(*protoMsg.LandLordsSceneResp)
    a := args[1].(gate.Agent)
    //person := a.UserData().(*protoMsg.UserInfo)
    log.Debug("进入游戏:%v", m)
    msg := &protoMsg.LandLordsReadyReq{
        IsReady: true,
    }
    a.WriteMsg(msg)

}

func handleEnterGameTuitongzi(args []interface{}) {
    m := args[0].(*protoMsg.TuitongziSceneResp)
    a := args[1].(gate.Agent)
    //person := a.UserData().(*protoMsg.UserInfo)
    log.Debug("进入游戏:%v", m)
    msg := &protoMsg.TuitongziReadyReq{
        IsReady: true,
    }
    a.WriteMsg(msg)

}

//////////////////////////////////////////////////////////

//////////////////////准备///////////////////////////////////
func handleQzcowcowStateFreeResp(args []interface{}) {
    m := args[0].(*protoMsg.QzcowcowStateFreeResp)
    a := args[1].(gate.Agent)
    //	person := a.UserData().(*protoMsg.UserInfo)
    second, _ := rand.Int(rand.Reader, big.NewInt(int64(m.Times.TotalTime)))
    time.AfterFunc(time.Duration(second.Int64())*time.Second, func() {
        msg := &protoMsg.QzcowcowReadyReq{
            IsReady: true,
        }
        a.WriteMsg(msg)
    })

}
func handleTbcowcowStateFreeResp(args []interface{}) {
    m := args[0].(*protoMsg.TbcowcowStateFreeResp)
    a := args[1].(gate.Agent)
    //	person := a.UserData().(*protoMsg.UserInfo)
    second, _ := rand.Int(rand.Reader, big.NewInt(int64(m.Times.TotalTime)))
    time.AfterFunc(time.Duration(second.Int64())*time.Second, func() {
        msg := &protoMsg.TbcowcowReadyReq{
            IsReady: true,
        }
        a.WriteMsg(msg)
    })

}
func handleMahjongStateFreeResp(args []interface{}) {
    m := args[0].(*protoMsg.MahjongStateFreeResp)
    a := args[1].(gate.Agent)
    second, _ := rand.Int(rand.Reader, big.NewInt(int64(m.Times.TotalTime)))
    time.AfterFunc(time.Duration(second.Int64())*time.Second, func() {
        msg := &protoMsg.MahjongReadyReq{
            IsReady: true,
        }
        a.WriteMsg(msg)
    })

}
func handleMahjongERStateFreeResp(args []interface{}) {
    m := args[0].(*protoMsg.MahjongERStateFreeResp)
    a := args[1].(gate.Agent)
    second, _ := rand.Int(rand.Reader, big.NewInt(int64(m.Times.TotalTime)))
    time.AfterFunc(time.Duration(second.Int64())*time.Second, func() {
        msg := &protoMsg.MahjongERReadyReq{
            IsReady: true,
        }
        a.WriteMsg(msg)
    })

}
func handleZhajinhuaStateFreeResp(args []interface{}) {
    m := args[0].(*protoMsg.ZhajinhuaStateFreeResp)
    a := args[1].(gate.Agent)
    second, _ := rand.Int(rand.Reader, big.NewInt(int64(m.Times.TotalTime)))
    time.AfterFunc(time.Duration(second.Int64())*time.Second, func() {
        msg := &protoMsg.ZhajinhuaReadyReq{
            IsReady: true,
        }
        a.WriteMsg(msg)
    })
}
func handleSangongStateFreeResp(args []interface{}) {
    m := args[0].(*protoMsg.SangongStateFreeResp)
    a := args[1].(gate.Agent)
    second, _ := rand.Int(rand.Reader, big.NewInt(int64(m.Times.TotalTime)))
    time.AfterFunc(time.Duration(second.Int64())*time.Second, func() {
        msg := &protoMsg.SangongReadyReq{
            IsReady: true,
        }
        a.WriteMsg(msg)
    })

}
func handleLandLordsStateFreeResp(args []interface{}) {
    m := args[0].(*protoMsg.LandLordsStateFreeResp)
    a := args[1].(gate.Agent)
    second, _ := rand.Int(rand.Reader, big.NewInt(int64(m.Times.TotalTime)))
    time.AfterFunc(time.Duration(second.Int64())*time.Second, func() {
        msg := &protoMsg.LandLordsReadyReq{
            IsReady: true,
        }
        a.WriteMsg(msg)
    })
}

/////////////////////////////发牌/////////////////////////////////////////////////
func handleMahjongDealResp(args []interface{}) {
    m := args[0].(*protoMsg.MahjongDealResp)
    a := args[1].(gate.Agent)
    person := a.UserData().(*protoMsg.UserInfo)
    if person.UserID == m.UserID {
        if _, ok := mjcards[m.UserID]; ok {
            mjcards[m.UserID] = append(mjcards[m.UserID], m.HandCards...)
        }
    }
}
func handleMahjongERDealResp(args []interface{}) {
    m := args[0].(*protoMsg.MahjongERDealResp)
    a := args[1].(gate.Agent)
    person := a.UserData().(*protoMsg.UserInfo)
    if person.UserID == m.UserID {
        if _, ok := mjercards[m.UserID]; ok {
            mjercards[m.UserID] = append(mjercards[m.UserID], m.HandCards...)
        }
    }
}

//////////////////////////对战类/////////////////////////////////////////////
func handleZhajinhuaStatePlayingResp(args []interface{}) {
    m := args[0].(*protoMsg.ZhajinhuaStatePlayingResp)
    a := args[1].(gate.Agent)
    person := a.UserData().(*protoMsg.UserInfo)
    if person.UserID == m.UserID {
        second, _ := rand.Int(rand.Reader, big.NewInt(int64(m.Times.TotalTime/2)))
        time.AfterFunc(time.Duration(second.Int64())*time.Second, func() {
            //是否看牌
            if 0 == second.Int64()%7 {
                msgLook := &protoMsg.ZhajinhuaLookReq{}
                a.WriteMsg(msgLook)
            }
            switch second.Int64() % 3 {
            case 0:
                msg := &protoMsg.ZhajinhuaFollowReq{}
                a.WriteMsg(msg)
            case 1:
                scoreR, _ := rand.Int(rand.Reader, big.NewInt(100))
                score := scoreR.Int64() * 100
                msg := &protoMsg.ZhajinhuaRaiseReq{
                    Score: score,
                }
                a.WriteMsg(msg)
            case 2:
                msg := &protoMsg.ZhajinhuaGiveupReq{}
                a.WriteMsg(msg)
            }
        })
    }
}

func handleMahjongStatePlayingResp(args []interface{}) {
    m := args[0].(*protoMsg.MahjongStatePlayingResp)
    a := args[1].(gate.Agent)
    person := a.UserData().(*protoMsg.UserInfo)
    if person.UserID == m.UserID && 0 < m.Card {
        msg := &protoMsg.MahjongOutCardReq{
            Card: m.Card,
        }
        //添加牌值并排序
        if _, ok := mjcards[m.UserID]; ok {
            mjcards[m.UserID] = append(mjcards[m.UserID], m.Card)
            sort.Slice(mjcards[m.UserID], func(i, j int) bool {
                return mjcards[m.UserID][i] < mjcards[m.UserID][j]
            })
        }
        //随机时间出牌
        second, _ := rand.Int(rand.Reader, big.NewInt(int64(m.Times.TotalTime/2)))
        time.AfterFunc(time.Duration(second.Int64()+1)*time.Second, func() {
            size := len(mjcards[m.UserID])
            if 0 < size {
                msg.Card = mjcards[m.UserID][size-1]
                //删除牌值
                mjcards[m.UserID] = DeleteValue(mjcards[m.UserID], msg.Card).([]int32)
            }
            a.WriteMsg(msg)
        })
    }
}

func handleMahjongERStatePlayingResp(args []interface{}) {
    m := args[0].(*protoMsg.MahjongERStatePlayingResp)
    a := args[1].(gate.Agent)
    person := a.UserData().(*protoMsg.UserInfo)
    if person.UserID == m.UserID && 0 < m.Card {
        msg := &protoMsg.MahjongEROutCardReq{
            Card: m.Card,
        }
        //添加牌值并排序
        if _, ok := mjercards[m.UserID]; ok {
            mjercards[m.UserID] = append(mjercards[m.UserID], m.Card)
            sort.Slice(mjercards[m.UserID], func(i, j int) bool {
                return mjercards[m.UserID][i] < mjercards[m.UserID][j]
            })
        }
        //随机时间出牌
        second, _ := rand.Int(rand.Reader, big.NewInt(int64(m.Times.TotalTime/2)))
        time.AfterFunc(time.Duration(second.Int64()+1)*time.Second, func() {
            size := len(mjercards[m.UserID])
            if 0 < size {
                msg.Card = mjercards[m.UserID][size-1]
                //删除牌值
                mjercards[m.UserID] = DeleteValue(mjercards[m.UserID], msg.Card).([]int32)
            }
            a.WriteMsg(msg)
        })
    }
}

func handleLandLordsStatePlayingResp(args []interface{}) {
    m := args[0].(*protoMsg.LandLordsStatePlayingResp)
    a := args[1].(gate.Agent)
    person := a.UserData().(*protoMsg.UserInfo)
    log.Debug("斗地主机器人")
    if person.UserID == m.TurnID {
        msg := &protoMsg.LandLordsTrusteeReq{IsTrustee: true}
        a.WriteMsg(msg)
        log.Debug("机器人：%v托管", person.UserID)
        //second, _ := rand.Int(rand.Reader, big.NewInt(int64(m.Times.TotalTime/2)))
        //time.AfterFunc(time.Duration(second.Int64())*time.Second, func() {
        //    msg := &protoMsg.LandLordsOutCardReq{
        //        Card: m.Card,
        //    }
        //    a.WriteMsg(msg)
        //})
    }
}

func handleTuitongziStatePlayingResp(args []interface{}) {
    m := args[0].(*protoMsg.TuitongziStatePlayingResp)
    a := args[1].(gate.Agent)
    person := a.UserData().(*protoMsg.UserInfo)
    log.Debug("推筒子机器人:%v", person.UserID)
    msg := &protoMsg.LandLordsTrusteeReq{IsTrustee: true}
    a.WriteMsg(msg)
    second, _ := rand.Int(rand.Reader, big.NewInt(int64(m.Times.TotalTime/2)))
    time.AfterFunc(time.Duration(second.Int64())*time.Second, func() {
        twice, _ := rand.Int(rand.Reader, big.NewInt(int64(5)))
        msg := &protoMsg.TuitongziBetReq{
            BetScore: twice.Int64(),
        }
        a.WriteMsg(msg)
    })
}

//////////////////操作///////////////////////////////////////////
func handleMahjongHintResp(args []interface{}) {
    m := args[0].(*protoMsg.MahjongHintResp)
    a := args[1].(gate.Agent)
    person := a.UserData().(*protoMsg.UserInfo)
    if person.UserID == m.UserID {
        //删除pass
        for _, v := range m.Hints {
            if v.Code == protoMsg.MJOperate_Pass {
                m.Hints = DeleteValue(m.Hints, v).([]*protoMsg.MahjongHint)
                break
            }
        }
        size := len(m.Hints)
        one, _ := rand.Int(rand.Reader, big.NewInt(int64(size)))
        index := int(one.Int64())
        for idx, hit := range m.Hints {
            if hit.Code == protoMsg.MJOperate_Hu || hit.Code == protoMsg.MJOperate_ZiMo {
                index = idx
                break
            }
        }
        msg := &protoMsg.MahjongOperateReq{
            Code:  m.Hints[index].Code,
            Cards: m.Hints[index].Cards,
        }
        a.WriteMsg(msg)

        //删除牌值
        if _, ok := mjcards[m.UserID]; ok {
            for _, card := range m.Hints[index].Cards {
                mjcards[m.UserID] = DeleteValue(mjcards[m.UserID], card).([]int32)
            }
            //
            sort.Slice(mjcards[m.UserID], func(i, j int) bool {
                return mjcards[m.UserID][i] < mjcards[m.UserID][j]
            })
            if m.Hints[index].Code == protoMsg.MJOperate_Chi || m.Hints[index].Code == protoMsg.MJOperate_Pong {
                time.AfterFunc(time.Duration(2)*time.Second, func() {
                    sizeX := len(mjcards[m.UserID])
                    if 0 < sizeX {
                        a.WriteMsg(&protoMsg.MahjongOutCardReq{
                            Card: mjcards[m.UserID][sizeX-1],
                        })
                        mjcards[m.UserID] = DeleteValue(mjcards[m.UserID], mjcards[m.UserID][sizeX-1]).([]int32)
                    }

                })
            }

        }

    }
}

func handleMahjongERHintResp(args []interface{}) {
    m := args[0].(*protoMsg.MahjongERHintResp)
    a := args[1].(gate.Agent)
    person := a.UserData().(*protoMsg.UserInfo)
    if person.UserID == m.UserID {
        //删除pass
        for _, v := range m.Hints {
            if v.Code == protoMsg.MJOperate_Pass {
                m.Hints = DeleteValue(m.Hints, v).([]*protoMsg.MahjongERHint)
                break
            }
        }
        size := len(m.Hints)
        one, _ := rand.Int(rand.Reader, big.NewInt(int64(size)))
        index := int(one.Int64())
        for idx, hit := range m.Hints {
            if hit.Code == protoMsg.MJOperate_Hu || hit.Code == protoMsg.MJOperate_ZiMo {
                index = idx
                break
            }
        }
        msg := &protoMsg.MahjongEROperateReq{
            Code:  m.Hints[index].Code,
            Cards: m.Hints[index].Cards,
        }
        a.WriteMsg(msg)
        //删除牌值
        if _, ok := mjercards[m.UserID]; ok {
            for _, card := range m.Hints[index].Cards {
                mjercards[m.UserID] = DeleteValue(mjercards[m.UserID], card).([]int32)
            }
            //
            sort.Slice(mjercards[m.UserID], func(i, j int) bool {
                return mjercards[m.UserID][i] < mjercards[m.UserID][j]
            })
            if m.Hints[index].Code == protoMsg.MJOperate_Chi || m.Hints[index].Code == protoMsg.MJOperate_Pong {
                size := len(mjercards[m.UserID])
                if 0 < size {
                    time.AfterFunc(time.Duration(2)*time.Second, func() {
                        a.WriteMsg(&protoMsg.MahjongEROutCardReq{
                            Card: mjercards[m.UserID][size-1],
                        })
                        mjercards[m.UserID] = DeleteValue(mjercards[m.UserID], mjercards[m.UserID][size-1]).([]int32)
                    })
                }
            }

        }

    }
}

/////////////////////////百人类////////////////////////////////
func handleBaccaratStatePlayingResp(args []interface{}) {
    m := args[0].(*protoMsg.BaccaratStatePlayingResp)
    a := args[1].(gate.Agent)
    //	person := a.UserData().(*protoMsg.UserInfo)
    secondR, _ := rand.Int(rand.Reader, big.NewInt(int64(m.Times.TotalTime)))
    areaR, _ := rand.Int(rand.Reader, big.NewInt(8))
    area := int32(areaR.Int64())
    scoreR, _ := rand.Int(rand.Reader, big.NewInt(100))
    score := scoreR.Int64() * 100
    time.AfterFunc(time.Duration(secondR.Int64())*time.Second, func() {
        msg := &protoMsg.BaccaratBetReq{
            BetArea:  area,
            BetScore: score,
        }
        a.WriteMsg(msg)
    })
}
func handleBrcowcowStatePlayingResp(args []interface{}) {
    m := args[0].(*protoMsg.BrcowcowStatePlayingResp)
    a := args[1].(gate.Agent)
    // person := a.UserData().(*protoMsg.UserInfo)
    secondR, _ := rand.Int(rand.Reader, big.NewInt(int64(m.Times.TotalTime)))
    areaR, _ := rand.Int(rand.Reader, big.NewInt(4))
    area := int32(areaR.Int64())
    scoreR, _ := rand.Int(rand.Reader, big.NewInt(100))
    score := scoreR.Int64() * 100
    time.AfterFunc(time.Duration(secondR.Int64())*time.Second, func() {
        msg := &protoMsg.BrcowcowBetReq{
            BetArea:  area,
            BetScore: score,
        }
        a.WriteMsg(msg)
    })
}

func handleBrtoubaoStatePlayingResp(args []interface{}) {
    m := args[0].(*protoMsg.BrtoubaoStatePlayingResp)
    a := args[1].(gate.Agent)
    //  person := a.UserData().(*protoMsg.UserInfo)
    secondR, _ := rand.Int(rand.Reader, big.NewInt(int64(m.Times.TotalTime)))
    areaR, _ := rand.Int(rand.Reader, big.NewInt(25))
    area := int32(areaR.Int64())
    scoreR, _ := rand.Int(rand.Reader, big.NewInt(100))
    score := scoreR.Int64() * 100
    time.AfterFunc(time.Duration(secondR.Int64())*time.Second, func() {
        msg := &protoMsg.BrtoubaoBetReq{
            BetArea:  area,
            BetScore: score,
        }
        a.WriteMsg(msg)
    })
}

func handleBrTuitongziStatePlayingResp(args []interface{}) {
    m := args[0].(*protoMsg.BrTuitongziStatePlayingResp)
    a := args[1].(gate.Agent)
    // person := a.UserData().(*protoMsg.UserInfo)
    secondR, _ := rand.Int(rand.Reader, big.NewInt(int64(m.Times.TotalTime)))
    areaR, _ := rand.Int(rand.Reader, big.NewInt(3))
    area := int32(areaR.Int64())
    scoreR, _ := rand.Int(rand.Reader, big.NewInt(100))
    score := scoreR.Int64() * 100
    time.AfterFunc(time.Duration(secondR.Int64())*time.Second, func() {
        msg := &protoMsg.BrTuitongziBetReq{
            BetArea:  area,
            BetScore: score,
        }
        a.WriteMsg(msg)
    })
}

func handleTigerXdragonStatePlayingResp(args []interface{}) {
    m := args[0].(*protoMsg.TigerXdragonStatePlayingResp)
    a := args[1].(gate.Agent)
    secondR, _ := rand.Int(rand.Reader, big.NewInt(int64(m.Times.TotalTime)))
    areaR, _ := rand.Int(rand.Reader, big.NewInt(13))
    area := int32(areaR.Int64())
    scoreR, _ := rand.Int(rand.Reader, big.NewInt(100))
    score := scoreR.Int64() * 100
    time.AfterFunc(time.Duration(secondR.Int64())*time.Second, func() {
        msg := &protoMsg.TigerXdragonBetReq{
            BetArea:  area,
            BetScore: score,
        }
        a.WriteMsg(msg)
    })
}
