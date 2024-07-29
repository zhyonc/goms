package main

import (
	"flag"
	"fmt"
	"goms/config"
	"goms/logger"
	"goms/mongodb"
	"goms/network"
	"goms/network/server"
	"goms/nxfile"
	"sync"

	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	loginServerType   string = "login"
	worldServerType   string = "world"
	defaultConfigPath string = "config.toml"
	defaultNXFileDir  string = "./nxfile"
)

var (
	serverTypePtr      *string
	configPathPtr      *string
	typeUsageTip       string = fmt.Sprintf("Denotes what type of server to start: %s|%s", loginServerType, worldServerType)
	configPathUsageTip string = "Denotes the path of server config file if empty will use default config"
)

func init() {
	serverTypePtr = flag.String("server", loginServerType, typeUsageTip)
	configPathPtr = flag.String("config", defaultConfigPath, configPathUsageTip)
	flag.Usage = func() {
		flag.PrintDefaults()
		fmt.Printf("\r\ngoms -server=%s", loginServerType)
	}
	flag.Parse()
}

func main() {
	// Config
	if *serverTypePtr != loginServerType && *serverTypePtr != worldServerType {
		slog.Error("Unkown server type:", "server", *serverTypePtr)
		return
	}
	var path string
	if *configPathPtr == "" {
		path = defaultConfigPath
	} else {
		path = *configPathPtr
	}
	conf := config.NewConfig(path)
	// Logger
	logger := logger.NewLogger(conf.Logger.LogLevel)
	slog.SetDefault(logger)
	// DB
	dbClient := mongodb.NewDBClient(conf.DB.DBURI, conf.DB.DBName)
	defer dbClient.Disconnect()
	// Server
	var s network.Server
	switch *serverTypePtr {
	case loginServerType:
		done := make(chan bool, 1)
		go func() {
			nxfile.ExtractCharacter(defaultNXFileDir)
			done <- true
		}()
		<-done
		s = server.NewLoginServer(conf, dbClient)
	case worldServerType:
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			nxfile.ExtractMap(defaultNXFileDir)
		}()
		wg.Wait()
		s = server.NewWorldServer(conf, dbClient)
	}
	// Avoid unexpected exit
	sch := make(chan os.Signal, 1)
	signal.Notify(sch, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		for sig := range sch {
			switch sig {
			case syscall.SIGTERM, syscall.SIGINT:
				slog.Info("Server will be closed after 5 second")
				time.Sleep(5 * time.Second)
				s.Stop()
				return
			default:
				slog.Info("other signal", "syscall", sig)
			}
		}
	}()
	s.Run()
	<-sch
}
