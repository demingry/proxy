package chkpro

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

var (
	pre_results = []string{}
	results = map[string]string{}
)


func Split_string(){
	for index,i:= range pre_results{
		line := strings.Split(i," ")
		if len(line)!=2{
			fmt.Printf("Please check file line nearby %d.\n",index)
			os.Exit(1)
		}
		results[line[0]] = line[1]
	}
}

func Read_file(filepath string)(IPLine map[string]string,stat bool){
	if file,err := os.Open(filepath);err!=nil {
		panic(err)
	}else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan(){
			text := scanner.Text()
			if text==""{
				continue
			}
			pre_results = append(pre_results,text)
		}
		fmt.Printf("Total %d lines.",len(pre_results))
	}
	Split_string()
	return results,true
}
