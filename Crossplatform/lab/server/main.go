package main

import "os"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		panic("Environment variable PORT not set")
	}
	server := NewServer(port)
	err := server.Listen()
	if err != nil {
		panic(err)
	}
}
