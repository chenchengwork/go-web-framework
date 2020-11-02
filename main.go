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

	if setting.ServerSetting.HttpPort > 0 && setting.ServerSetting.HttpsPort > 0 {
		go startServer(false)
		startServer(true)
	} else if setting.ServerSetting.HttpPort > 0 && setting.ServerSetting.HttpsPort == 0 {
		startServer(false)
	} else if setting.ServerSetting.HttpPort == 0 && setting.ServerSetting.HttpsPort > 0 {
		startServer(true)
	} else {
		fmt.Println("未指定http或https的端口!")
	}
}

func startServer(isHttps bool) {
	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	maxHeaderBytes := 1 << 20

	if !isHttps {
		endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
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
	} else {
		endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpsPort)
		server := &http.Server{
			Addr:           endPoint,
			Handler:        routersInit,
			ReadTimeout:    readTimeout,
			WriteTimeout:   writeTimeout,
			MaxHeaderBytes: maxHeaderBytes,

			//
			TLSConfig: &tls.Config{
				MinVersion:               tls.VersionTLS12,
				CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
				PreferServerCipherSuites: true,
				CipherSuites: []uint16{
					tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
					tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
					tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
					tls.TLS_RSA_WITH_AES_256_CBC_SHA,
				},
			},
			TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
		}

		log.Printf("[info] start https server listening %s", endPoint)

		// https 方式
		server.ListenAndServeTLS(setting.AppSetting.RuntimeRootPath+"ssl/cert.pem", setting.AppSetting.RuntimeRootPath+"ssl/key.pem")
	}
}
