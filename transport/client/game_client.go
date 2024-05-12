package client

import (
	"encoding/binary"
	"fmt"
	"goms/opcode"
	"goms/packet/crypt"
	"goms/transport"
	"log/slog"
	"net"
	"strconv"
	"strings"
)

type IChild interface {
	handlePacket(op uint16, data []byte)
}

type gameClient struct {
	conn    net.Conn
	crypter crypt.Crypter
	IChild
}

func NewGameClient(conn net.Conn, child IChild) transport.GameClient {
	c := &gameClient{
		conn:    conn,
		crypter: crypt.NewCrypter(),

		IChild: child,
	}
	return c
}

// Disconnect implements transport.Client.
func (c *gameClient) Disconnect() {
	c.conn.Close()
	c = nil
	slog.Info("GameClient was disconnected")
}

// RecvPacket implements transport.GameClient.
func (c *gameClient) RecvPacket() {
	isHeader := true
	readSize := crypt.EncryptHeaderLen
	defer c.Disconnect()
	for {
		buf := make([]byte, readSize)
		_, err := c.conn.Read(buf)
		if err != nil {
			slog.Error("Failed to receive packet", "err", err, "addr", c.conn.RemoteAddr())
			break
		}
		if isHeader {
			readSize = crypt.DecodeHeader(buf)
		} else {
			readSize = crypt.EncryptHeaderLen
			c.crypter.Decrypt(buf)
			op := binary.LittleEndian.Uint16(buf[0:2])
			data := buf[2:]
			logPacket("In", op, data)
			c.IChild.handlePacket(op, data)
		}
		isHeader = !isHeader
	}
}

// SendPacket implements transport.GameClient.
func (c *gameClient) SendPacket(buf []byte) {
	op := binary.LittleEndian.Uint16(buf[0:2])
	data := buf[2:]
	_, err := c.conn.Write(c.crypter.Encrypt(buf))
	if err != nil {
		slog.Error("Failed to send packet", "err", err, "addr", c.conn.RemoteAddr())
		return
	}
	logPacket("Out", op, data)
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
	slog.Debug("["+tag+"]", "opcode", hex, "field", field, "data", formatData(data))

}

func formatData(data []byte) string {
	length := len(data)
	if length == 0 {
		return ""
	}
	lastIndex := length - 1
	var builder strings.Builder
	for i, v := range data {
		builder.WriteString(strconv.Itoa(int(v)))
		if i < lastIndex {
			builder.WriteString(",")
		}
	}
	return builder.String()
}
