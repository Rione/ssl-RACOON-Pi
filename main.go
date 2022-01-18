package main

import (
	"flag"
	"github.com/Rione-SSL/RACOON-Pi/proto/pb_gen"
	"github.com/golang/protobuf/proto"
	"log"
	"net"
	"time"
)

var (
	mode = flag.String("m", "server", "mode: client or server")
	port = flag.String("p", "4000", "host: ip:port")
)

func main() {
	flag.Parse()

	switch *mode {
	case "server":
		RunServer()
	}
}

func RunServer() {
	serverAddr, err := net.ResolveUDPAddr("udp", ":"+*port)
	CheckError(err)

	serverConn, err := net.ListenUDP("udp", serverAddr)
	CheckError(err)
	defer serverConn.Close()

	buf := make([]byte, 1024)

	log.Println("Listening on port " + *port)
	for {
		n, addr, err := serverConn.ReadFromUDP(buf)
		packet := &pb_gen.GrSim_Packet{}
		err = proto.Unmarshal(buf[0:n], packet)
		log.Printf("Received %d sent at %s from %s", *packet.Serial, time.Unix(*packet.SentTime, 0), addr)

		if err != nil {
			log.Fatal("Error: ", err)
		}
	}
}

func CheckError(err error) {
	if err != nil {
		log.Fatal("Error: ", err)
	}
}
