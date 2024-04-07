package main

import (
	"log"

	cmap "github.com/orcaman/concurrent-map/v2"
)

var ma = cmap.New[string]()

func main() {
	log.Print("listening upd server on port 3008")
	listenUdp(3008, handler)
}
