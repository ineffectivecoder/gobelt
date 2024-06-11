package gobelt

import (
	"net"
	"strings"
)

func IPQuery() []string {
	addrs, _ := net.InterfaceAddrs()
	var value []string
	var ip net.IP
	for _, addr := range addrs {
		tmp := addr.String()
		tmp = tmp[0:strings.Index(tmp, "/")]

		if ip = net.ParseIP(tmp); (ip == nil) || ip.IsLoopback() {
			continue
		}
		if ip.IsLinkLocalMulticast() || ip.IsLinkLocalUnicast() {
			continue
		}
		value = append(value, tmp)
	}
	return value
}
