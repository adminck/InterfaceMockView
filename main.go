package main

import (
	"InterfaceMockView/api/https"
	"InterfaceMockView/models"
	"InterfaceMockView/router"
	"InterfaceMockView/utils/common"
	"InterfaceMockView/utils/log"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"path/filepath"
	"time"
)

var config = NewConfig()

func main() {
	os.Chdir(filepath.Dir(common.GetCurrentProcessPath()))

	if err := config.Load(); err != nil {
		log.Error("config.json read failed")
		return
	}

	log.InitLog(config.Log)

	if err := common.CreateDir("./data"); err != nil {
		log.Error("CreateDir data failed")
		return
	} //创建存放数据目录

	DB, err := models.DBinit(config.Db)
	if err != nil {
		log.Error("DBinit failed")
		return
	} //初始化数据库连接
	defer DB.Close() // 程序结束前关闭数据库链接

	if err := RunHttpServer(); err != nil {
		log.Error("HttpServer start failed")
		return
	} //启动http server

	//进程退出
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Infoln("recv signal interrupt, exit...")
}

func RunHttpServer() error {
	Router := router.Routers()
	HttpAddress := fmt.Sprintf(":%d", 80)
	hs := &http.Server{
		Addr:           HttpAddress,
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	https.SSLServerMgr.Router = Router
	https.SSLServerMgr.Start()
	if err := hs.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
