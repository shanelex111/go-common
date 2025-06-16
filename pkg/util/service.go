package util

import (
	"net"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	headerXOriginalForwardedFor = "X-Original-Forwarded-For"
	headerXForwardedFor         = "X-Forwarded-For"
	headerProxyClientIP         = "Proxy-Client-IP"
	headerWLProxyClientIP       = "WL-Proxy-Client-IP"
	headerRealIP                = "X-Real-IP"
	unknown                     = "unknown"

	localhostIp   = "127.0.0.1"
	localhostIp16 = "[::1]"
)

func GetUUID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

func GetIP(c *gin.Context) string {
	var ip = ""
	ip = c.GetHeader(headerXOriginalForwardedFor)

	if len(strings.TrimSpace(ip)) == 0 || unknown == ip {
		ip = c.GetHeader(headerRealIP)
	}

	if len(strings.TrimSpace(ip)) == 0 || unknown == ip {
		ip = c.GetHeader(headerXForwardedFor)
	}

	if len(strings.TrimSpace(ip)) == 0 || unknown == ip {
		ip = c.GetHeader(headerProxyClientIP)
	}
	if len(strings.TrimSpace(ip)) == 0 || unknown == ip {
		ip = c.GetHeader(headerWLProxyClientIP)
	}

	if len(strings.TrimSpace(ip)) == 0 || unknown == ip {
		ip = c.Request.RemoteAddr
	}

	if strings.Contains(ip, localhostIp16) || strings.Contains(ip, localhostIp) || net.ParseIP(ip).IsLoopback() {
		addrs, _ := net.InterfaceAddrs()
		for _, address := range addrs {
			// 检查 ip 地址判断是否回环地址
			if ipnet, flag := address.(*net.IPNet); flag && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					ip = ipnet.IP.String()
					break
				}
			}
		}
	}

	if ip != "" {
		ip = strings.Split(ip, ",")[0]
		ip = strings.Split(ip, ":")[0]
	}
	return ip
}
