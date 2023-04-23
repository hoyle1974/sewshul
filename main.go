package main

import (
	"flag"

	log "github.com/sirupsen/logrus"
)

var (
	port = flag.Int("port", 50051, "The server port")
	ip   = flag.String("ip", "0.0.0.0", "address")
)

func main() {
	flag.Parse()

	log.Println("Starting . . .")

}
