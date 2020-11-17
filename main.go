package main

import (
	"InterfaceMockView/models"
	"InterfaceMockView/router"
	"InterfaceMockView/utils/common"
	"InterfaceMockView/utils/log"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"runtime/debug"
	"time"
)

type Certificates struct {
	CertFile	string
	KeyFile		string
}

var config = NewConfig()

func main() {
	os.Chdir(filepath.Dir(common.GetCurrentProcessPath()))

	go func() {
		for {
			debug.FreeOSMemory()
			time.Sleep(time.Minute)
		}
	}()

	initLog(config.Log)
	if err := config.Load(); err != nil {
		log.Error("config.json read failed")
		return
	}

	if err := common.CreateDir("./data"); err != nil {
		log.Error("CreateDir data failed")
		return
	} //创建存放数据目录

	DB,err := models.DBinit(config.Db)
	if err != nil {
		log.Error("DBinit failed")
		return
	} //初始化数据库连接
	defer DB.Close()  // 程序结束前关闭数据库链接

	if err := RunHttpServer();err != nil {
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
	HttpsAddress := fmt.Sprintf(":%d", 443)
	hs := &http.Server{
		Addr:           HttpAddress,
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	hss := &http.Server{
		Addr:           HttpsAddress,
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if Certs := GetCerts(); len(Certs) != 0 {
		go func() {
			if err := ListenAndServeTLSSNI(hss,Certs); err != nil {
				log.Error(err)
			}
		}()
	}else {
		log.Warn("Certs len not")
	}
	if err := hs.ListenAndServe();err != nil {
		return err
	}
	return nil
}

func initLog(cfg *LogConfig) {
	level := log.StringToLevel(cfg.LogLevel)
	log.SetLevel(level)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var (
		writer io.Writer
		err    error
	)
	writer, err = log.NewFileWriter(
		path.Join(filepath.Dir(common.GetCurrentProcessPath()), "log", "InterfaceMockView.log"),
		log.ReserveDays(cfg.ReserveDays),
		log.RotateByDaily(true),
		log.LogFileMaxSize(cfg.MaxSize),
	)
	if err != nil {
		log.Errorln("create file writer error:", err)
		return
	}

	if cfg.PrintScreen {
		writer = io.MultiWriter(writer, os.Stdout)
	}

	log.SetOutput(writer)
}

func ListenAndServeTLSSNI(srv *http.Server,Certs []Certificates) error {
	addr := ":https"
	certs := Certs
	config := &tls.Config{}
	if srv.TLSConfig != nil {
		*config = *srv.TLSConfig
	}
	if config.NextProtos == nil {
		config.NextProtos = []string{"http/1.1"}
	}

	var err error

	config.Certificates = make([]tls.Certificate, len(certs))
	for i, v := range certs {
		config.Certificates[i], err = tls.LoadX509KeyPair(v.CertFile, v.KeyFile)
		if err != nil {
			return err
		}
	}

	config.BuildNameToCertificate()

	conn, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	tlsListener := tls.NewListener(conn, config)
	err = srv.Serve(tlsListener)
	if err!=nil {
		return err
	}
	return nil
}

func GetCerts() []Certificates{
	var certs []Certificates
	var domains []models.Domain
	if err := models.DB.Find(&domains).Error; err != nil {
		fmt.Println("error:certslist Query Fail")
		return nil
	}
	for _,k := range domains{
		certs = append(certs, Certificates{
			CertFile: "./data/Domain/" + k.CrtFilePath,
			KeyFile: "./data/Domain/" + k.KeyFilePath,
		})
	}
	if len(certs) == 0 {
		fmt.Println("error:certs len not")
		return nil
	}
	return certs
}