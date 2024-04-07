package main

import (
	"errors"
	"fmt"
	"log"
	"net"
)

func listenUdp(port int, handler func(data []byte) ([]byte, error)) {
	addr, _ := net.ResolveUDPAddr("udp", fmt.Sprint(":", port))
	sock, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}
	if err := errors.Join(sock.SetWriteBuffer(1024*1024), sock.SetReadBuffer(1024*1024)); err != nil {
		panic(err)
	}
	for {
		buf := make([]byte, 1024)
		_, client, err := sock.ReadFrom(buf)
		if err != nil {
			log.Printf("oops, packet read failure: %s", err.Error())
		}
		sendBack := func(data []byte) {
			_, err := sock.WriteTo(data, client)
			if err != nil {
				log.Printf("oops, packet write failure: %s", err.Error())
			}
		}
		go func() {
			res, err := handler(buf)
			if err != nil {
				sendBack([]byte(err.Error()))
			}
			sendBack(res)
		}()
	}
}
