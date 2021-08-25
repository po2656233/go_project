## kiss gate

[![MIT licensed][1]][2]
[![Go Report Card][3]][4]

[1]: https://img.shields.io/badge/license-MIT-blue.svg
[2]: LICENSE.md
[3]: https://goreportcard.com/badge/github.com/nothollyhigh/kissgate
[4]: https://goreportcard.com/report/github.com/nothollyhigh/kissgate


- kiss/net 网关，反代 tcp、websocket 协议到后端 tcp 线路，

- 支持线路检测、负载均衡、realip等，详见源码


## 安装

- go get github.com/nothollyhigh/kissgate


## 运行

- kissgate -config="config.xml"


## 配置


```xml
<setting>
     <!-- debug: 设置日志是否输出到控制台 -->
    <!-- logdir: 日志目录 -->
    <!-- redirect: 是否开启全局tcp重定向 -->
    <!-- heartbeat: 心跳设置 -->
    <options debug="true" logdir="./logs/" redirect="false">
        <heartbeat interval="10" timeout="50"></heartbeat>
    </options>
    <proxy>
        <!-- tcp 10000 端口 反代到 tcp 10001 10002 端口 -->
        <busline name="ws" addr=":9950" type="websocket" redirect="" tls="false" realipmode="http">
            <route path="/"></route>
            <line serverid="HALL" addr=":9950" type="websocket" redirect="" tls="false" realipmode="http">
                <node addr="" ip="127.0.0.1" port="9958" maxload="5000" enable="false"></node>
            </line>
            <line serverid="HALL1" addr="" type="websocket" redirect="" tls="false" realipmode="websocket">
                <node addr="" ip="127.0.0.1" port="9000" maxload="1" enable="true"></node>
            </line>
        </busline>
        <!-- tcp 10000 端口 反代到 tcp 10001 10002 端口 -->
        <busline name="tcp" addr=":20002" type="tcp" redirect="" tls="false" realipmode="tcp">
            <line serverid="HALL" addr=":20002" type="tcp" redirect="" tls="false" realipmode="tcp">
                <node addr="" ip="127.0.0.1" port="10021" maxload="50000" enable="false"></node>
            </line>
        </busline>
    </proxy>
    <!-- 对外服务端口 register:服务注册-->
    <api addr=":10001" type="http" registerpath="/register" querypath="/info" reloadpath="/reload" enablepath="/enableLine" disablepath=""></api>
</setting>
```
```
http://网关IP:10001/register  [服务注册]   
post方式
HEAD
Authorization: Basic IyNzc3NeXl46KFM/U1MmXi4xNA==
Content-Type: multipart/form-data;
以下是form-data格式数据
[key]:[value]
type:   websocket   //  tcp websocket
name:   HALL1       //  服务ID
ip  :   127.0.0.1   //  指定的服务IP
port:   9000        //  指定的服务端口
maxload:1           //  负载量
```
````
http://网关IP:10001/info      [网关信息] get
http://网关IP:10001/reload    [重新加载配置] post
http://网关IP:10001/enableLine [线路启用/停用] post 关联control.xml
````
## 示例

-  使用上面示例的配置启动网关

```sh
kissgate -config="config.xml"
```

- 启动后端tcp服务器，tcp/websocket各两个端口

```
cd kissgate/example
go run tcpserver.go
```

- 启动测试用客户端

```
cd kissgate/example
go run gateclient.go
```

- 观察网关、客户端、服务器日志，代码详见

1. [server](https://github.com/nothollyhigh/kissgate/blob/master/example/tcpserver.go)

2. [client](https://github.com/nothollyhigh/kissgate/blob/master/example/gateclient.go)
