package proxy

import (
	"net"
)

type IPNet struct {
	IP  net.IP
	Net *net.IPNet
}
