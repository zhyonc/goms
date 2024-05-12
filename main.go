package main

import (
	"flag"
	"fmt"
	"goms/config"
	"goms/transport"
	"goms/transport/server"
	"log"
)

const (
	loginServer   string = "login"
	channelServer string = "channel"
)

var defaultConfigPath map[string]string = map[string]string{
	loginServer:   "login_config.toml",
	channelServer: "channel_config.toml",
}

var (
	typePtr            *string
	configPathPtr      *string
	typeUsageTip       string = fmt.Sprintf("Denotes what type of server to start: %s|%s", loginServer, channelServer)
	configPathUsageTip string = "Denotes the path of server config file if empty will use default config"
)

func init() {
	typePtr = flag.String("type", "", typeUsageTip)
	configPathPtr = flag.String("config", "", configPathUsageTip)
	flag.Usage = func() {
		flag.PrintDefaults()
		fmt.Printf("\r\nFor example: goms -type=%s -config=%s", loginServer, defaultConfigPath[loginServer])
	}
	flag.Parse()
}

func main() {
	if *typePtr == "" {
		log.Println("Unkown server type:", *typePtr)
		return
	}
	var path string
	if *configPathPtr == "" {
		path = defaultConfigPath[*typePtr]
	} else {
		path = *configPathPtr
	}
	var s transport.Server
	switch *typePtr {
	case loginServer:
		conf := config.NewLoginConfig(path)
		s = server.NewLoginServer(conf)
	case channelServer:
		conf := config.NewChannelConfig(path)
		s = server.NewChannelServer(conf)
	}
	s.Run()
}
