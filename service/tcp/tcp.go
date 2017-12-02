// Package tcp provides the mq service via udp transmit
package tcp

import (
	"encoding/json"
	"net"
	"time"

	"github.com/chlins/Gomq/log"
	"github.com/chlins/Gomq/mq"
	"github.com/chlins/Gomq/mq/channel"
)

// Service is the service mq via udp
type Service struct {
	mq   mq.MsQ
	port string
	ln   net.Listener
}

// Reg Info
type Reg struct {
	Role  string `json:"role"`
	Topic string `json:"topic"`
	Cap   int    `json:"cap"`
}

// NewService is the constructor of mq
func NewService(p string, m mq.MsQ) *Service {
	svc := new(Service)
	lisn, err := net.Listen("tcp", "0.0.0.0:"+p)
	if err != nil {
		log.Fatal("Start tcp server failed, %s", err)
	}
	svc.mq = m
	svc.port = p
	svc.ln = lisn
	return svc
}

// Start udp mq server
func (s *Service) Start() {
	log.Success("[TCP] message queue service starting...")
	go s.accept()
}

func (s *Service) accept() {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			log.Error("Accept Client failed, %s", err)
		}
		go s.handle(conn)
	}
}

func (s *Service) handle(conn net.Conn) {
	buf := make([]byte, 10240)
	n, _ := conn.Read(buf)
	var reg Reg
	if err := json.Unmarshal(buf[:n], &reg); err == nil {
		if reg.Role == "producer" {
			mq := channel.NewMq(reg.Cap)
			channel.AddMq(reg.Topic, mq)
			for {
				if s.mq.Full(reg.Topic) {
					conn.Write([]byte("Channel is up to cap limit"))
					conn.Close()
					return
				}
				n, _ := conn.Read(buf)
				s.mq.Push(reg.Topic, string(buf[:n]))
			}
		} else if reg.Role == "consumer" {
			for {
				if !s.mq.Empty(reg.Topic) {
					msg := s.mq.Pop(reg.Topic)
					if msg != "" {
						conn.Write([]byte(msg))
					}
				}
				<-time.After(10 * time.Millisecond)
			}
		}
		return
	}
	defer func() {
		conn.Write([]byte("Invalid request"))
		conn.Close()
	}()
}
