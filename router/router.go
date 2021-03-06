package router

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"ws/common"
	"ws/util"
)

func WsPush() {
	var wsPort = common.Ws.WsPort
	if err := common.CheckPort(int(wsPort)); err != nil {
		log.Fatal(err.Error())
	}
	wsPush := http.NewServeMux()
	wsPush.HandleFunc("/", WsRouter())
	wsPush.HandleFunc("/all", AllNodeRouter())
	common.LogInfoSuccess(fmt.Sprintf("创建WebSocket服务端口:%d", wsPort))
	if err := http.ListenAndServe(":"+strconv.Itoa(int(wsPort)), wsPush); err != nil {
		log.Fatal("main:", err)
	}
}
func HttpPush() {
	var httpPort = common.Http.HttpPort
	if err := common.CheckPort(int(httpPort)); err != nil {
		log.Fatal(err.Error())
	}
	var httpTimeOut = common.Http.HttpTimeOut
	httpPush := http.NewServeMux()
	httpPush.HandleFunc("/", HttpRouter())
	httpPush.HandleFunc("/token", HttpGetToken())
	httpPushTimeOut := http.TimeoutHandler(httpPush, time.Duration(httpTimeOut)*time.Second, util.TimeOut)
	common.LogInfoSuccess(fmt.Sprintf("创建HTTP服务端口:%d", httpPort))
	if err := http.ListenAndServe(":"+strconv.Itoa(int(httpPort)), httpPushTimeOut); err != nil {
		log.Fatal(err)
	}
}
//当http服务和ws服务公用一个端口的时候启用
func goHttpRouteHandle() {

}
