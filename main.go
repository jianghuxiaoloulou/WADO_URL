package main

import (
	"WowjoyProject/WADO_URL/global"
	v1 "WowjoyProject/WADO_URL/internal/routers/api/v1"
	"net/http"
)

// @title WADO_URL服务
// @version 1.0.0.1
// @description 实现WADO服务
// @termsOfService https://github.com/jianghuxiaoloulou/wado_url.git
func main() {
	global.Logger.Info("*******WADO_URL服务********")

	http.HandleFunc("/wado", v1.FileDownload)
	http.HandleFunc("/WADO", v1.FileDownload)


	// router := &routers.MyMux{}
	// ser := &http.Server{
	// 	Addr:           ":" + global.ServerSetting.HttpPort,
	// 	Handler:        router,
	// 	ReadTimeout:    global.ServerSetting.ReadTimeout,
	// 	WriteTimeout:   global.ServerSetting.WriteTimeout,
	// 	MaxHeaderBytes: 1 << 20,
	// }
	global.Logger.Fatal(http.ListenAndServe(":"+global.ServerSetting.HttpPort, nil))

}
