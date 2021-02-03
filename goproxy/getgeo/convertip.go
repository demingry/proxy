package getgeo


import (
	"fmt"
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"
)

func GetGeo(IP string)(ISO string){

	db,err:=geoip2.Open("GeoLite2-City.mmdb")
	if err!=nil {
		log.Fatal(err)
	}
	defer db.Close()
	ip:=net.ParseIP(IP)
	record,err:=db.City(ip)
	if err!=nil {
	log.Fatal(err)
	}

	fmt.Println("City Name: ",record.City.Names["en"])
	if len(record.Subdivisions)>0{
		fmt.Println("Subdivision name: ",record.Subdivisions[0].Names["en"])
	}
	fmt.Println("ISO country: ",record.Country.IsoCode)
	return record.Country.IsoCode
}
