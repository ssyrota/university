package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"regexp"
	"strings"

	cmap "github.com/orcaman/concurrent-map/v2"
)

var isPrivateR = regexp.MustCompile(`\[.*\]`)

type Client struct {
	name string
	conn net.Conn
}

func (c *Client) handle() {
	input := bufio.NewScanner(c.conn)
	for input.Scan() {
		inputTxt := input.Text()
		if isPrivateR.Match([]byte(inputTxt)) {
			values := strings.SplitN(strings.TrimPrefix(inputTxt, "["), "]", 2)
			to := values[0]
			msg := values[1]
			messages <- postMessage(to, msg, c)
			continue
		}

		messages <- postMessage("", inputTxt, c)
	}
	clients.Remove(c.name)
	leaving <- postLeft(c)
}

var clients = cmap.New[Client]()
var leaving = make(chan message)
var messages = make(chan message)

type message struct {
	text    string
	address string
	to      string
}

func newClient(conn net.Conn) (*Client, error) {
	input := bufio.NewScanner(conn)
	input.Scan()
	name := input.Text()
	client := &Client{conn: conn, name: name}
	if ok := clients.SetIfAbsent(client.name, *client); !ok {
		return nil, fmt.Errorf("name %s is already taken", name)
	}
	messages <- postJoined(client)
	return client, nil
}

func postMessage(to, msg string, client *Client) message {
	addr := client.conn.RemoteAddr().String()
	return message{
		text:    client.name + ": " + msg,
		address: addr,
		to:      to,
	}
}
func postJoined(client *Client) message {
	addr := client.conn.RemoteAddr().String()
	return message{
		text:    client.name + " joined.",
		address: addr,
	}
}
func postLeft(client *Client) message {
	addr := client.conn.RemoteAddr().String()
	return message{
		text:    client.name + " left chat.",
		address: addr,
	}
}
func broadcaster() {
	for {
		select {
		case msg := <-messages:
			if msg.to != "" {
				conn, ok := clients.Get(msg.to)
				if !ok {
					continue
				}
				if msg.address != conn.conn.RemoteAddr().String() {
					fmt.Fprint(conn.conn, "[private] "+msg.text)
				}
			} else {
				for _, conn := range clients.Items() {
					if msg.address == conn.conn.RemoteAddr().String() {
						continue
					}
					fmt.Fprint(conn.conn, msg.text)
				}
			}
		case msg := <-leaving:
			for _, conn := range clients.Items() {
				fmt.Fprint(conn.conn, msg.text)
			}
		}
	}
}
func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go func(conn net.Conn) {
			defer conn.Close()
			client, err := newClient(conn)
			if err != nil {
				conn.Write([]byte(err.Error()))
				return
			}
			client.handle()
		}(conn)
	}
}
