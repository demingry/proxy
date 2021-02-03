package chkpro

import (
	"context"
	"fmt"
	"golang.org/x/net/proxy"
	"net"
	"net/http"
	"time"
)

type SocksProxy struct {
}

func (SocksProxy) IsProxy(proxyIP string, proxyPort int, timeout int)(isProxy bool, err error) {
	proxyAddr:=fmt.Sprintf("%s:%d",proxyIP,proxyPort)
	dialSocksProxy, err:=proxy.SOCKS5("tcp",proxyAddr,nil,proxy.Direct)
	if err!=nil{
	return false,nil
	}
	netTransport:=&http.Transport{DialContext: func(ctx context.Context, network, addr string)(conn net.Conn,e error) {
		c,e:=dialSocksProxy.Dial(network,addr)
		return c,e
	}}
	client:=&http.Client{
		Timeout: time.Second * time.Duration(timeout),
		Transport: netTransport,
	}
	return CheckProxy(client)
}
