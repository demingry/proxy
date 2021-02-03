package chkpro

import (
	"io/ioutil"
	"net/http"
	"strings"
)

const TestURL = "http://httpbin.org/ip"

type Proxyer interface {
	IsProxy(ProxyIP string, proxyPort int)(isProxy bool,err error)
}

func CheckProxy(client *http.Client)(isProxy bool,err error){
	res,err:=client.Get(TestURL)
	if err!=nil{
	return false,err
	}else{
	defer res.Body.Close()
	if res.StatusCode == 200 {
		body,err:=ioutil.ReadAll(res.Body)
		if err==nil&&strings.Contains(string(body),"origin") {
			return true,nil
		}else{
			return false,err
		}
	}else{
		return false,nil
	}
	}
}
