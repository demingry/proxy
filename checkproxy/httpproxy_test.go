package checkproxy

import (
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestCheckHttpProxy(t *testing.T) {
	go func(){
		proxy:=goproxy.NewProxyHttpServer()
		proxy.Verbose = true
		log.Fatal(http.ListenAndServe(":8080",proxy))
	}()
	time.Sleep(time.Second * 2)
	isProxy,err:= HttpProxy{}.IsProxy("127.0.0.1",8080)
	if !isProxy {
		t.Error("shold be a proxy",err)
	}
}
