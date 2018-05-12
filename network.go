package talent

import (
	"io"
	"net"
	"strings"
)

// 判断一个error是否是io.EOF
func IsEOF(err error) bool {
	if err == nil {
		return false
	} else if err == io.EOF {
		return true
	} else if oerr, ok := err.(*net.OpError); ok {
		if oerr.Err.Error() == "use of closed network connection" {
			return true
		}
	} else {
		if err.Error() == "use of closed network connection" {
			return true
		}
	}
	return true
}

// 获取本机ip
func LocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	return filterIP(addrs)
}

func TransfarIP() string {
	l, err := net.Interfaces()
	if err != nil {
		return ""
	}

	for _, i := range l {
		l, _ := i.Addrs()
		switch i.Name {
		case "br0", "eth0", "en0":
			ip := filterIP(l)
			if ip == "" {
				continue
			} else {
				return ip
			}
		default:

		}
	}

	for _, i := range l {
		l, _ := i.Addrs()
		if strings.HasPrefix(i.Name, "eth") {
			ip := filterIP(l)
			if ip == "" {
				continue
			} else {
				return ip
			}
		}
	}

	return ""
}

func filterIP(addrs []net.Addr) string {
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				if !strings.Contains(ipnet.IP.String(), "192.168") {
					return ipnet.IP.String()
				}
			}
		}
	}

	return ""
}
