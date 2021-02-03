package main


import (
	"fmt"
	"github.com/kirinlabs/HttpRequest"
	"os"
	"regexp"
	"goproxy/getgeo"
	"github.com/Unknwon/goconfig"
)


var (
	apikey string
	max_threads int
	request_number int
	result_file string
	Threads chan int
	Syncchan chan struct{}
)


func InitConf() {
	cfg,err:=goconfig.LoadConfigFile("./Configure")
	if err!=nil {
		fmt.Print("[!]No configuration.")
	}
	apikey = cfg.MustValue("default","apikey","null")
	max_threads = cfg.MustInt("default","max_threads",4)
	request_number = cfg.MustInt("default","request_number",20)
	result_file = cfg.MustValue("default","result_file","./newfiletext")
	Threads = make(chan int, max_threads)
	Syncchan = make(chan struct{}, request_number+10)
}

func main(){
	defer func(){
		if p:=recover();p!=nil {
			fmt.Println("Got panic: ",p)
		}
	}()
	InitConf()
	for i:=0;i<request_number;i++{
		go func(){
			Threads <- 1
			Syncchan <- struct{}{}
			request_url := fmt.Sprintf("http://api.scraperapi.com/?api_key=%s&url=http://httpbin.org/ip",apikey)
			res,_:=HttpRequest.Get(request_url)
			defer res.Close()
			body,_:=res.Body()
			html := string(body)
			fmt.Print(html)
			reg := regexp.MustCompile(`(?is:^{.*?origin": "(.*?)".*?)`)
			parsetxt := reg.FindStringSubmatch(html)
			f,err:=os.OpenFile(result_file,os.O_WRONLY|os.O_CREATE|os.O_APPEND,0666)
			defer f.Close()
			if err!=nil {
				fmt.Println("err: ",err)
			}else{
				ISO := getgeo.GetGeo(parsetxt[1])
				parsetxt[1]=parsetxt[1]+" "+ISO+"\n"
				_,err=f.Write([]byte(parsetxt[1]))
				if err!=nil {
					fmt.Println("err: ",err)
				}
			}
			<-Threads
		}()
		<-Syncchan
	}
	//defer f.Close()
}
