package main

import (
	"atp/payment/app/endpoint"
	"atp/payment/app/usecase/blockchain"
	"atp/payment/pkg/adapter/levelDB"
	"atp/payment/pkg/repository/transaction"
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	log.Println("============= START ===============")
	ctx := context.Background()

	path := "database/"
	db, err := levelDB.NewConnection(path)
	if err != nil {
		log.Fatalf("FAILED connect to database:" + err.Error())
	}

	repoTrans := transaction.NewRepository(db, path+".key.db")
	appBC := blockchain.NewBlockChain(repoTrans)

	bc := appBC.CreateBlockchain(ctx, fmt.Sprintf("%x", [32]byte{}))

	echoNew := echo.New()
	echoNew.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:    true,
		LogMethod:    true,
		LogURI:       true,
		LogUserAgent: true,
		LogLatency:   true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			logEcho := fmt.Sprintf("{status:%v} {method:%v} {latency:%v} {uri:%v} {user_agent:%v}", values.Status, values.Method, values.Latency, values.URI, values.UserAgent)
			if values.Status != 200 {
				log.Printf("[error] [logEcho] %s", logEcho)
			} else {
				log.Printf("[info] [logEcho] %s", logEcho)
			}
			return nil
		},
	}))
	//endpoint
	endpoint.NewHandler(echoNew, "blockchain/", appBC, repoTrans, bc)

	errServer := make(chan error)
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

	echoNew.Server.TLSConfig = cfg
	echoNew.Server.Addr = ":8008"
	//optional
	timeout := 1 * time.Minute
	echoNew.Server.ReadTimeout = timeout
	echoNew.Server.WriteTimeout = timeout
	echoNew.Server.IdleTimeout = timeout

	runServer := func() {
		log.Printf("[info] server running on port:[%s]", echoNew.Server.Addr)
		errServer <- echoNew.Server.ListenAndServe()
	}

	go runServer()

	for {
		select {
		case <-ctx.Done():
			ctxShutDown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			go func(ctx context.Context) {
				defer cancel()
				// shutdown server
				if err := echoNew.Shutdown(ctxShutDown); err != nil {
					log.Fatalf("[fatal] server shutdown failed:%s" + err.Error())
				}
				log.Fatal("[fatal] server exited properlys")
			}(ctx)

		case err := <-errServer:
			log.Fatalf("[fatal] server error got:%s" + err.Error())
		}
	}
}
