package main

import (
	"encoding/json"
	"log"
	"net"
	"strconv"
)

// Msg message struct
type Msg struct {
	Body string `json:"body"`
}

func main() {
	socket, err := net.DialUDP("udp4", nil, &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1), Port: 8001})
	if err != nil {
		log.Fatal("Udp connect failed, ", err)
	}
	defer socket.Close()
	var sendData []byte
	for i := 0; i < 1000; i++ {
		data := &Msg{Body: "Hello, i am message " + strconv.Itoa(i)}
		sendData, _ = json.Marshal(data)
		_, err = socket.Write(sendData)
		log.Printf("Send message %d\n", i)
		if err != nil {
			log.Fatal("Send message failed, ", err)
		}
	}
}
