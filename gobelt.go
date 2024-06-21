package gobelt

import (
	"fmt"
	"net"
	"strings"
)

type Checker struct {
	checks []Check
}

func NewChecker() *Checker {
	c := &Checker{}
	c.checks = []Check{
		Example,
		IPQuery,
		MappedDrives,
	}
	return c
}

func (c Checker) Checks() []Check {
	return c.checks
}

type Check func() Result

type ResultKind int

const (
	KindError ResultKind = iota
	KindInfo
)

type Result struct {
	Kind  ResultKind
	Data  []string
	Error error
}

func (r Result) String() string {
	switch r.Kind {
	case KindError:
		return fmt.Sprintf("ERROR: %v", r.Error)
	case KindInfo:
		return fmt.Sprintf("INFO: %v", strings.Join(r.Data, "\n"))
	default:
		return "unknown result type"
	}
}

func Example() Result {
	return Result{
		Kind: KindInfo,
		Data: []string{"hello", "world"},
	}
}

func IPQuery() Result {
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
	return Result{
		Kind: KindInfo,
		Data: value,
	}
}
