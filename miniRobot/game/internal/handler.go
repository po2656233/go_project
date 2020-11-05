package internal

import (
    "crypto/rand" //真随机
    "github.com/name5566/leaf/gate"
    "github.com/name5566/leaf/log"
    "math/big"
    protoMsg "miniRobot/msg/go"
    "reflect"
    "time"
)

//初始化
func init() {
    //游戏处理
    handlerMsg(&protoMsg.EnterGameResp{}, handleEnterGame)         //反馈--->主页信息
    handlerMsg(&protoMsg.EnterGameQZCCResp{}, handleEnterGameQZCC) //反馈--->主页信息
    handlerMsg(&protoMsg.EnterGameTBCCResp{}, handleEnterGameTBCC) //反馈--->主页信息
    handlerMsg(&protoMsg.EnterGameZJHResp{}, handleEnterGameZJH)   //反馈--->主页信息
    handlerMsg(&protoMsg.EnterGameMJResp{}, handleEnterGameMJ)     //反馈--->主页信息
    handlerMsg(&protoMsg.EnterGameMJERResp{}, handleEnterGameMJER) //反馈--->主页信息
    handlerMsg(&protoMsg.EnterGameSGResp{}, handleEnterGameSG)     //反馈--->主页信息
    handlerMsg(&protoMsg.LandLordsSceneResp{}, handleEnterGameLandLords)     //反馈--->主页信息


    //
    handlerMsg(&protoMsg.QzcowcowStateFreeResp{}, handleQzcowcowStateFreeResp)   //反馈--->主页信息
    handlerMsg(&protoMsg.TbcowcowStateFreeResp{}, handleTbcowcowStateFreeResp)   //反馈--->主页信息
    handlerMsg(&protoMsg.MahjongStateFreeResp{}, handleMahjongStateFreeResp)     //反馈--->主页信息
    handlerMsg(&protoMsg.MahjongERStateFreeResp{}, handleMahjongERStateFreeResp) //反馈--->主页信息
    handlerMsg(&protoMsg.ZhajinhuaStateFreeResp{}, handleZhajinhuaStateFreeResp) //反馈--->主页信息
    handlerMsg(&protoMsg.SangongStateFreeResp{}, handleSangongStateFreeResp)     //反馈--->主页信息
    handlerMsg(&protoMsg.LandLordsStateFreeResp{}, handleLandLordsStateFreeResp) //反馈--->主页信息

    //
    handlerMsg(&protoMsg.BaccaratStatePlayingResp{}, handleBaccaratStatePlayingResp)         //反馈--->主页信息
    handlerMsg(&protoMsg.BrcowcowStatePlayingResp{}, handleBrcowcowStatePlayingResp)         //反馈--->主页信息
    handlerMsg(&protoMsg.BrtoubaoStatePlayingResp{}, handleBrtoubaoStatePlayingResp)         //反馈--->主页信息
    handlerMsg(&protoMsg.TuitongziStatePlayingResp{}, handleTuitongziStatePlayingResp)       //反馈--->主页信息
    handlerMsg(&protoMsg.TigerXdragonStatePlayingResp{}, handleTigerXdragonStatePlayingResp) //反馈--->主页信息

    ///对战出牌
    handlerMsg(&protoMsg.ZhajinhuaStatePlayingResp{}, handleZhajinhuaStatePlayingResp) //反馈--->主页信息
    handlerMsg(&protoMsg.MahjongStatePlayingResp{}, handleMahjongStatePlayingResp) //反馈--->主页信息
    handlerMsg(&protoMsg.MahjongERStatePlayingResp{}, handleMahjongERStatePlayingResp) //反馈--->主页信息


    //操作提示
    handlerMsg(&protoMsg.MahjongHintResp{}, handleMahjongHintResp) //反馈--->主页信息
    handlerMsg(&protoMsg.MahjongERHintResp{}, handleMahjongERHintResp) //反馈--->主页信息

}

//注册传输消息
func handlerMsg(m interface{}, h interface{}) {
    skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

///////////////////////入场/////////////////////////////
//百人类
func handleEnterGame(args []interface{}) {
    m := args[0].(*protoMsg.EnterGameResp)
    a := args[1].(gate.Agent)
    person := a.UserData().(*protoMsg.UserInfo)
    if m.UserInfo.UserID == person.UserID {
        log.Debug("进入游戏:%v", m)
    }

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
    person := a.UserData().(*protoMsg.UserInfo)
    if m.Player.MyInfo.UserID == person.UserID {
        log.Debug("进入游戏:%v", m)
        msg := &protoMsg.MahjongReadyReq{
            IsReady: true,
        }
        a.WriteMsg(msg)
    }

}

func handleEnterGameMJER(args []interface{}) {
    m := args[0].(*protoMsg.EnterGameMJERResp)
    a := args[1].(gate.Agent)
    person := a.UserData().(*protoMsg.UserInfo)
    if m.Player.MyInfo.UserID == person.UserID {
        log.Debug("进入游戏:%v", m)
        msg := &protoMsg.MahjongERReadyReq{
            IsReady: true,
        }
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

//////////////////////////对战类/////////////////////////////////////////////
func handleZhajinhuaStatePlayingResp(args []interface{}) {
    m := args[0].(*protoMsg.ZhajinhuaStatePlayingResp)
    a := args[1].(gate.Agent)
    person := a.UserData().(*protoMsg.UserInfo)
    if person.UserID == m.UserID {
        second, _ := rand.Int(rand.Reader, big.NewInt(int64(m.Times.TotalTime)))
        time.AfterFunc(time.Duration(second.Int64())*time.Second, func() {
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
    if person.UserID == m.UserID {
        second, _ := rand.Int(rand.Reader, big.NewInt(int64(m.Times.TotalTime)))
        time.AfterFunc(time.Duration(second.Int64())*time.Second, func() {
            msg := &protoMsg.MahjongOutCardReq{
                Card: m.Card,
            }
            a.WriteMsg(msg)
        })
    }
}

func handleMahjongERStatePlayingResp(args []interface{}) {
    m := args[0].(*protoMsg.MahjongERStatePlayingResp)
    a := args[1].(gate.Agent)
    person := a.UserData().(*protoMsg.UserInfo)
    if person.UserID == m.UserID {
        second, _ := rand.Int(rand.Reader, big.NewInt(int64(m.Times.TotalTime)))
        time.AfterFunc(time.Duration(second.Int64())*time.Second, func() {
            msg := &protoMsg.MahjongEROutCardReq{
                Card: m.Card,
            }
            a.WriteMsg(msg)
        })
    }
}


//////////////////操作///////////////////////////////////////////
func handleMahjongHintResp(args []interface{}) {
    m := args[0].(*protoMsg.MahjongHintResp)
    a := args[1].(gate.Agent)
    person := a.UserData().(*protoMsg.UserInfo)
    if person.UserID == m.UserID {
        size:=len(m.Hints)
        one,_:=rand.Int(rand.Reader, big.NewInt(int64(size)))
        index:=one.Int64()
        msg := &protoMsg.MahjongOperateReq{
            Code: m.Hints[int(index)].Code,
            Cards: m.Hints[int(index)].Cards,
        }
        a.WriteMsg(msg)
    }
}
func handleMahjongERHintResp(args []interface{}) {
    m := args[0].(*protoMsg.MahjongERHintResp)
    a := args[1].(gate.Agent)
    person := a.UserData().(*protoMsg.UserInfo)
    if person.UserID == m.UserID {
        size:=len(m.Hints)
        one,_:=rand.Int(rand.Reader, big.NewInt(int64(size)))
        index:=one.Int64()
        msg := &protoMsg.MahjongEROperateReq{
            Code: m.Hints[int(index)].Code,
            Cards: m.Hints[int(index)].Cards,
        }
        a.WriteMsg(msg)
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
func handleTuitongziStatePlayingResp(args []interface{}) {
    m := args[0].(*protoMsg.TuitongziStatePlayingResp)
    a := args[1].(gate.Agent)
    // person := a.UserData().(*protoMsg.UserInfo)
    secondR, _ := rand.Int(rand.Reader, big.NewInt(int64(m.Times.TotalTime)))
    areaR, _ := rand.Int(rand.Reader, big.NewInt(3))
    area := int32(areaR.Int64())
    scoreR, _ := rand.Int(rand.Reader, big.NewInt(100))
    score := scoreR.Int64() * 100
    time.AfterFunc(time.Duration(secondR.Int64())*time.Second, func() {
        msg := &protoMsg.TuitongziBetReq{
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
