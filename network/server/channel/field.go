package channel

import (
	"goms/nx"
	"time"
)

type Field struct {
	ips   map[string]time.Time
	MapNX *nx.MapNX
}

func NewField(ip string, mapNX *nx.MapNX) *Field {
	f := &Field{
		ips: make(map[string]time.Time),
	}
	f.ips[ip] = time.Now()
	f.MapNX = mapNX
	return f
}

func (f *Field) AddIP(newIP string) {
	_, ok := f.ips[newIP]
	if !ok {
		f.ips[newIP] = time.Now()
	}
}

func (f *Field) RemoveIP(ip string) {
	_, ok := f.ips[ip]
	if !ok {
		return
	}
	delete(f.ips, ip)
}

func (f *Field) GetIPsLength() int {
	return len(f.ips)
}

func (f *Field) GetIPs() map[string]time.Time {
	return f.ips
}
