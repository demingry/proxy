package chkpro

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type HttpProxy struct {
}

func (HttpProxy) IsProxy(proxyIP string, proxyPort int, timeout int)(isPorxy bool, err error) {
	proxyURL:=fmt.Sprintf("http://%s:%d",proxyIP,proxyPort)
	proxy,err:=url.Parse(proxyURL)
	if err!=nil{
		return false,err
	}
	netTransport := &http.Transport {
		Proxy: http.ProxyURL(proxy),
	}
	client:=&http.Client {
		Timeout: time.Second * time.Duration(timeout),
		Transport: netTransport,
	}
	return CheckProxy(client)
}
