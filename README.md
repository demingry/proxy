
## what's this?
*  : Call scraperAPI to gain your proxyip.
*  : Whether proxy is live or not. Seperated from goproxy.
*  : 通过调用scraper的api获取代理ip地址
*  : 验证代理是否存活

## idea
* 主要欧美国家的代理是很好找的，
  并且AWS、Digitalocean、Azure等提供的也主要是美国、德国、新加坡、香港等地的机房
  可以搭建SS、V2ray、Trojan，但有特殊情况我们也不得不使用如非洲、中亚、南美的http和socks代理，
  免费的代理站点很多，写这个主要是娱乐为主,过后可能会逐步完善，比如添加代理可匿性验证，
  甚至变成一个端口扫描器，抓banner和指纹。

## usage
#### 
*  : Edit Configure file, append your api key for scraperapi,
* max_threads number of goroutine, scraperapi free-plan max to 5.
* request_number number of request each time, free-plan max to 1000 requests per day.
* result_file output file name.
*  : go run main.go
*  : 在配置文件Configure中添加您的scraperapi key
* max_thread是起的抓取线程，scraper免费版好像最多5个routine
* request_number是每次抓取请求的个数，免费版每天最多请求1000次
* result_file是输出结果的文件
*  : go run main.go
#### 
* -h proxyip, 
* -p port(port range:7000-8000/sigle port: 8080/multi ports: 8080,9090,10000 are supported)
* -t timeout
* -f input file(ipaddr+space+port(all supported), line by line) without -h and -p
* -h 待检验ip地址
* -p 待检验端口,支持单个/多个端口以及端口范围
* -t 超时时间
* -f 输入文件(ip+空格+port)，不能使用-h和-p了

## issue
* post your issue here, or georgedeming@outlook.com

##about

* runoob here....... demingry@hzuna
* our home -> [hzuna](https://github.com/hzuna)
