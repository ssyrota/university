package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	defer func() {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}()

	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(`
Welcome to chat!

After setting up nickname you will be automatically connected to broadcast channel.

To send private message type [${name}] in start of your message

Enter your nickname: `)
	name, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	_, err = conn.Write([]byte(name))
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			msgBuff := make([]byte, 1024)
			_, err := conn.Read(msgBuff)
			if err != nil {
				log.Printf("%v", "closing")
				os.Exit(0)
			}
			log.Printf("%v", string(msgBuff))
		}
	}()
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		_, err = conn.Write([]byte(str))
		if err != nil {
			panic(err)
		}
	}
}
