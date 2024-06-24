package gobelt

import (
	"fmt"
	"net"
	"strings"
)

type Check func() Result
type Checker struct {
	checks []Check
}
type Result struct {
	Kind  ResultKind
	Data  []string
	Error error
}
type ResultKind int

const (
	KindError ResultKind = iota
	KindInfo
)

func NewChecker() *Checker {
	c := &Checker{}
	c.checks = []Check{IPQuery}
	c.checks = append(c.checks, osSpecific()...)
	return c
}

func (c Checker) Checks() []Check {
	return c.checks
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
