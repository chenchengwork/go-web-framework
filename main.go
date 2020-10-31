package main

import (
	"crypto/tls"
	"fmt"
	"github.com/chenchengwork/go-web-framework/models"
	"github.com/chenchengwork/go-web-framework/pkg/logging"
	"github.com/chenchengwork/go-web-framework/pkg/setting"
	"github.com/chenchengwork/go-web-framework/pkg/util"
	"github.com/chenchengwork/go-web-framework/pkg/validation"
	"github.com/chenchengwork/go-web-framework/routers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func init() {
	setting.Setup()
	logging.Setup()
	util.Setup()
	models.Setup()
	validation.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	go startHttp()

	startHttps()
}

func startHttp() {
	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	// http 方式
	server.ListenAndServe()
}

func startHttps() {
	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpsPort)
	maxHeaderBytes := 1 << 20

	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,

		//
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}

	log.Printf("[info] start http server listening %s", endPoint)

	// https 方式
	server.ListenAndServeTLS("cert.pem", "key.pem")
}
