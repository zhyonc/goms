package api

import (
	"encoding/json"
	"goms/util"
	"net"
	"time"
)

type APICode uint16

const (
	TestPacket     APICode = 0
	SkipSDOAuth    APICode = 1 // Launcher->Router->LoginServer
	KickGameClient APICode = 2 // Launcher->Router->LoginServer
	QueryGaugePx   APICode = 3 // LoginServer->WorldServer
	PrepareMigrate APICode = 4 // LoginServer->WorldServer
)

// Use for server to server
type Message struct {
	ClientIP string  `json:"client_ip"`
	APICode  APICode `json:"api_code"`
	Status   string  `json:"status"`
	Content  []byte  `json:"content"`
}

func NewMessage(clientIP string, code APICode, content []byte) Message {
	m := Message{
		ClientIP: clientIP,
		APICode:  code,
		Content:  content,
	}
	return m
}

func (m *Message) Send(addr string, xorKey []byte) error {
	conn, err := net.Dial("udp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()
	buf, err := json.Marshal(m)
	if err != nil {
		return err
	}
	util.SimpleXOR(buf, xorKey)
	_, err = conn.Write(buf)
	if err != nil {
		return err
	}
	buf = make([]byte, 1024)
	conn.SetReadDeadline(time.Now().Add(3 * time.Second))
	n, err := conn.Read(buf)
	if err != nil {
		return err
	}
	data := buf[0:n]
	util.SimpleXOR(data, xorKey)
	err = json.Unmarshal(data, m)
	if err != nil {
		return err
	}
	return nil
}

type SkipSDOAuthRequest struct {
	AccountID uint32 `json:"account_id"`
}

type QueryGaugePxResponse struct {
	GaugePx []uint32 `json:"gauge_px"`
}
