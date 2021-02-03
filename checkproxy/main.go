package main


import (
	"fmt"
	"sync"
	"time"
	"flag"
	"strings"
	"strconv"
	"os"
)

import "checkproxy/chkpro"


/*
 target to scan struct
*/
type target struct{
	proxyIP string
	portslice []int
	issocks bool
}

/*
 target init
*/
func Newtarget(proxyIP string,is_socks bool)*target{
	return &target{proxyIP,[]int{},is_socks}
}

/*
 convert param port_range to real port
*/
func (t *target)Port_handle(port_range string){
	//port range
	if strings.Contains(port_range,"-"){
		stringsplit := strings.Split(port_range,"-")
		start_temp,err1 := strconv.Atoi(stringsplit[0])
		end_temp,err2 := strconv.Atoi(stringsplit[1])
		if err1!=nil || err2!=nil {
			panic("[!]port range invalid.\n")
		}else {
			if start_temp<0 || start_temp>65535 || end_temp<0 || end_temp>65535 || end_temp<start_temp {
				panic("[!]please check port range.\n")
			}else{
				fmt.Println("start: ",start_temp,"end: ",end_temp,"\n")
				for port:=start_temp;port<=end_temp;port++ {
					t.portslice = append(t.portslice,port)
				}
			}
		}//multi ports
	}else if strings.Contains(port_range,","){
		stringsplit := strings.Split(port_range,",")
		for _,i:= range stringsplit{
			port,err := strconv.Atoi(i)
			if err==nil{
				t.portslice = append(t.portslice,port)
			}
		}
	}else { //single port
		startend,_:=strconv.Atoi(port_range)
		if startend<0 || startend>65535 {
			panic("[!]please check port range.\n")
		}
		t.portslice = append(t.portslice,starten)
	}
}


/*
 global var
*/
var (
	f func(proxyIP string,proxyPort int,timeout int)(isProxy bool, err error)
	issocks bool
	timeout int
	targets []*target
)

/*
 print usage
*/
func PrintUsage() {
	fmt.Println("[!]usage:")
	fmt.Println("\tgo run main.go -h `proxyIP` -p `sigle port or port range or multi ports` -t `timeout` -s `socks else http` -f `file path`\n")
}


/*
 parse params and check
*/
func Param_parse() {
	host := flag.String("h","null","proxyIP")
	port_range := flag.String("p","null","port range")
	time_out := flag.String("t","20","timeout")
	is_socks := flag.Bool("s",false,"socks or http")
	file_path := flag.String("f","null","file path to check")

	flag.Parse()


	if((*host=="null" || *port_range=="null")&&*file_path=="null") {
		PrintUsage()
		os.Exit(1)
	}
	if((*host!="null"&&*file_path!="null")||(*port_range!="null"&&*file_path!="null")){
		PrintUsage()
		os.Exit(1)
	}


	if(*file_path=="null"){
		newtarget := Newtarget(*host,*is_socks)
		newtarget.Port_handle(*port_range)
		targets = append(targets,newtarget)
	}else{
		filepath := *file_path
		IPLine,_ := chkpro.Read_file(filepath)
		for index,value := range IPLine{
			newtarget := Newtarget(index,*is_socks)
			newtarget.Port_handle(value)
			targets = append(targets,newtarget)
		}
	}

	if timeout_temp,err:=strconv.Atoi(*time_out);err!=nil{
		panic("[!]numer of timeout is invalid.\n")
	}else{
		timeout = timeout_temp
	}

	return
}


/*
 type http or socks
*/
func Istype(issocks bool) {
	if issocks {
		f = chkpro.SocksProxy{}.IsProxy
	}else {
		f = chkpro.HttpProxy{}.IsProxy
			//(proxyIP string, proxyPort int, timeout int)(isProxy bool, err error)
	}
}




func main(){


	var current = sync.WaitGroup{}
	Param_parse()
	var now = time.Now().Unix()
	var Map = make(map[string]int)
	for _,host := range targets{
		Istype(host.issocks)
		for _,port:= range host.portslice{
			current.Add(1)
			go func(ip string, port int){
				isProxy,err := f(ip,port,timeout)
				if isProxy {
					Map[ip] = port
				}
				if err!=nil{
					fmt.Println(err)
				}
				current.Done()
			}(host.proxyIP,port)
		}
	}
	current.Wait()
	fmt.Printf("Totally %d host(s). Live %d .\n",len(targets),len(Map))
	fmt.Printf("Time els: %d.\n",time.Now().Unix()-now)
	fmt.Println(Map)
}
