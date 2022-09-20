package main

import (
	"fmt"
	"github.com/oschwald/geoip2-golang"
	"go-test-examples/geoip/gateway"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"testing"
)

func TestGateway(t *testing.T) {
	firstByte, secondByte := gateway.GetLatency("multichain.storage:https")
	log.Println("Everything:", firstByte, secondByte)
	firstByte, secondByte = gateway.GetLatency("ipfs.io:https")
	log.Println("Everything:", firstByte, secondByte)

	db, err := geoip2.Open("/media/ccao/data/dataset/GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ips, _ := net.LookupIP("multichain.storage")
	for _, ip := range ips {
		if ipv4 := ip.To4(); ipv4 != nil {
			fmt.Printf("IPv4: %s\n", ipv4)
			record, _ := db.City(ip.To4())
			fmt.Printf("English subdivision name: %v\n", record.Subdivisions[0].Names["en"])
			if len(record.Subdivisions) > 0 {
				fmt.Printf("English subdivision name: %v\n", record.Subdivisions[0].Names["en"])
			}

		}
	}
}

func TestLocation(t *testing.T) {
	multichainUrl := "multichain.storage"
	ipfsIoUrl := "ipfs.io"
	multichainFirstByte, secondByte := gateway.GetLatency(multichainUrl + ":https")
	log.Println(multichainUrl, multichainFirstByte, secondByte)
	ipfsIoUrlFirstByte, ipfsIoUrlsecondByte := gateway.GetLatency(ipfsIoUrl + ":https")
	log.Println(ipfsIoUrl, ipfsIoUrlFirstByte, ipfsIoUrlsecondByte)

	ips, _ := net.LookupIP("ipfs.io")
	for _, ip := range ips {
		if ipv4 := ip.To4(); ipv4 != nil {
			fmt.Printf("IPv4: %s\n", ipv4)
			resp, err := http.Get("https://ipinfo.io/" + ipv4.String())
			if err != nil {
				log.Fatalln(err)
			}
			//We Read the response body on the line below.
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatalln(err)
			}
			//Convert the body to type string
			sb := string(body)
			log.Printf(sb)

		}
	}

}
