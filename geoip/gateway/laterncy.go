package gateway

import (
	"io/ioutil"
	"net"
	"time"
)

func GetLatency(address string) (time.Duration, time.Duration) {
	//https://stackoverflow.com/questions/30526946/time-http-response-in-go
	conn, err := net.Dial("tcp", address)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	conn.Write([]byte("GET / HTTP/1.0\r\n\r\n"))

	start := time.Now()
	oneByte := make([]byte, 1)
	_, err = conn.Read(oneByte)
	if err != nil {
		panic(err)
	}
	firstByte := time.Since(start)
	_, err = ioutil.ReadAll(conn)
	if err != nil {
		panic(err)
	}
	secondByte := time.Since(start)
	return firstByte, secondByte

}
