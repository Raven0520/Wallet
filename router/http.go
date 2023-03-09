package router

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/raven0520/wallet/app"
)

var HTTPHandler *http.Server

// HttpServerStart Start Http / Https Service
func HTTPServerStart() {
	var (
		host = app.GetStringConfig("base.http.host")
		port = app.GetStringConfig("base.http.port")
		addr = host + ":" + port
	)
	gin.SetMode(app.BaseConf.Base.DebugMode)
	router := InitRouter()
	HTTPHandler = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    time.Duration(app.GetIntConfig("base.http.read_timeout")) * time.Second,
		WriteTimeout:   time.Duration(app.GetIntConfig("base.http.write_timeout")) * time.Second,
		MaxHeaderBytes: 1 << uint(app.GetIntConfig("base.http.max_header_bytes")),
	}
	err := StorePid()
	if err != nil {
		log.Printf(" [ERROR] HttpServerRun:%s err: %v\n", addr, errors.New("Store PID Failed"))
	}
	go func() {
		log.Printf(" [INFO] Http Server : %s\n", addr)
		if err := HTTPHandler.ListenAndServe(); err != nil {
			// if err := HTTPHandler.ListenAndServeTLS("./cert/"+host+".pem", "./cert/"+host+".key"); err != nil {
			if err != http.ErrServerClosed {
				log.Printf(" [ERROR] HttpServerRun:%s err: %v\n", addr, err)
			}
		}
	}()
}

// StorePid 保存 Pid
func StorePid() error {
	pid := os.Getpid()
	pidFile, err := os.OpenFile(app.BaseConf.Path.Pid, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return fmt.Errorf("Could not open pid file: %v", err)
	}
	defer pidFile.Close()
	_, err = pidFile.WriteString(fmt.Sprintf("%d", pid))
	if err != nil {
		return fmt.Errorf("Could not write to pid file: %s", err)
	}
	return nil
}

// HttpServerStop 关闭 Http 服务
func HTTPServerStop() {
	text, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := HTTPHandler.Shutdown(text); err != nil {
		fmt.Println("Shutdown Error ", err.Error())
		log.Fatalf(" [ERROR] HttpServerStop err: %v\n", err)
		return
	}
	log.Printf(" [INFO] HttpServerStop stopped \n")
}
