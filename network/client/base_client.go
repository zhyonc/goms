package client

import (
	"encoding/binary"
	"fmt"
	"goms/network"
	"goms/network/crypter"
	"goms/opcode"
	"goms/packet/inpacket"
	"goms/packet/outpacket"
	"goms/util"
	"log/slog"
	"net"
	"time"
)

const (
	pingInterval uint16 = 0
	pingTimeout  uint16 = 60
)

type IChild interface {
	handlePacket(op uint16, data []byte)
}

type baseClient struct {
	conn        net.Conn
	crypter     crypter.Crypter
	pingAckTime time.Time
	server      network.Server
	IChild
}

func NewBaseClient(conn net.Conn, recvIV, sendIV [4]byte, child IChild, server network.Server) network.Client {
	c := &baseClient{
		conn:    conn,
		crypter: crypter.NewCrypter(recvIV, sendIV),
		IChild:  child,
		server:  server,
	}
	return c
}

// Disconnect implements network.Client.
func (c *baseClient) Disconnect() {
	if c.conn != nil {
		c.server.KickClient(c.GetClientIP())
		c.conn.Close()
		c.conn = nil
		slog.Info("Client was disconnected")
	}
}

// IsDisconnected implements network.Client.
func (c *baseClient) IsDisconnected() bool {
	// if nil, server will delete the client from clientMap on clock cycle
	return c.conn == nil
}

// GetClientIP implements network.Client.
func (c *baseClient) GetClientIP() string {
	addr := c.conn.RemoteAddr()
	ip, _, _ := net.SplitHostPort(addr.String())
	return ip
}

// RecvPacket implements network.Client.
func (c *baseClient) RecvPacket() {
	isHeader := true
	readSize := crypter.EncryptHeaderLen
	defer c.Disconnect()
	for {
		buf := make([]byte, readSize)
		_, err := c.conn.Read(buf)
		if err != nil {
			slog.Error("Failed to receive packet", "err", err, "addr", c.conn.RemoteAddr())
			break
		}
		if isHeader {
			readSize = crypter.DecodeHeader(buf)
		} else {
			readSize = crypter.EncryptHeaderLen
			c.crypter.Decrypt(buf)
			op := binary.LittleEndian.Uint16(buf[0:2])
			data := buf[2:]
			logPacket("In", op, data)
			switch op {
			case opcode.CClientSocket_OnAliveReq_Callback:
				c.ClientPingServer(data)
				continue
			case opcode.CWvsApp_SendBackupPacket:
				c.logClientError(data)
				continue
			}
			c.IChild.handlePacket(op, data)
		}
		isHeader = !isHeader
	}
}

// SendPacket implements network.Client.
func (c *baseClient) SendPacket(buf []byte) {
	op := binary.LittleEndian.Uint16(buf[0:2])
	data := buf[2:]
	_, err := c.conn.Write(c.crypter.Encrypt(buf))
	if err != nil {
		slog.Error("Failed to send packet", "err", err, "addr", c.conn.RemoteAddr())
		return
	}
	logPacket("Out", op, data)
}
func (c *baseClient) ServerPingClient() {
	if pingInterval == 0 {
		return
	}
	defer c.Disconnect()
	pingIntervalDuration := time.Duration(pingInterval) * time.Second
	pingTimeoutDuration := time.Duration(pingTimeout) * time.Second
	c.pingAckTime = time.Now()
	ticker := time.NewTicker(pingIntervalDuration)
	defer ticker.Stop()
	for range ticker.C {
		if c == nil {
			break
		}
		c.SendPacket(outpacket.NewAliveReq())
		since := time.Since(c.pingAckTime)
		if since > pingTimeoutDuration {
			break
		}
	}
}

func (c *baseClient) ClientPingServer(buf []byte) {
	if pingInterval == 0 {
		return
	}
	in := inpacket.NewAliveReqCallback(buf)
	slog.Debug("Client ping Server", "NewIV", in.IV)
	c.pingAckTime = time.Now()
}

func logPacket(tag string, op uint16, data []byte) {
	hex := fmt.Sprintf("%d/0x%x", op, byte(op))
	var field string
	if tag == "In" {
		_, ok := opcode.NotLogInSet[op]
		if ok {
			return
		}
		field = opcode.InMap[op]
	}
	if tag == "Out" {
		_, ok := opcode.NotLogOutSet[op]
		if ok {
			return
		}
		field = opcode.OutMap[op]
	}
	if field == "" {
		field = "Unknown"
	}
	slog.Debug("["+tag+"]", "opcode", hex, "field", field, "data", util.FormatData(data))

}

func (c *baseClient) logClientError(data []byte) {
	in := inpacket.NewBackupPacket(data)
	callTypeStr := "UnknownType"
	switch in.CallType {
	case 1:
		callTypeStr = "SendBackupPacket"
	case 2:
		callTypeStr = "CrashReport"
	case 3:
		callTypeStr = "Exception"
	}
	iv := in.BackupBuffer[0:4]
	op := binary.LittleEndian.Uint16(in.BackupBuffer[4:6])
	hex := fmt.Sprintf("%d/0x%x", op, byte(op))
	packet := in.BackupBuffer[6:]
	slog.Debug("[ClientError]", "addr", c.conn.RemoteAddr(), "CallType", callTypeStr, "ErrorCode",
		in.ErrorCode, "IV", util.FormatData(iv), "opcode", hex, "buffer", util.FormatData(packet))
}
