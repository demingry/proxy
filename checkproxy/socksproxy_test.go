package checkproxy

import (
	"github.com/armon/go-socks5"
	"testing"
	"time"
)

func TestIsSocksProxy(t *testing.T) {
	go func() {
		conf := &socks5.Config{}
		server,err:=socks5.New(conf)
		if err!=nil{
			panic(err)
		}
		if err:= server.ListenAndServe("tcp","127.0.0.1:8002");err!=nil {
			panic(err)
		}
	}()
	time.Sleep(time.Second * 2)
	isProxy,err:=SocksProxy{}.IsProxy("127.0.0.1",8002)
	if !isProxy {
		t.Error("should be a proxy",err)
	}
}
