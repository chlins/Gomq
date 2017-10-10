package app

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/zcytop/Gomq/log"
	"github.com/zcytop/Gomq/util"
)

// App is the map of mq
type App struct {
	Mqs map[string]*MQ
	sync.Mutex
}

var wg sync.WaitGroup

// MQ message queue
type MQ struct {
	Config      util.Config
	Topic       string
	MessageChan chan util.Message
	Wg          sync.WaitGroup
}

// Init initialize an App
func Init() *App {
	return &App{
		Mqs: make(map[string]*MQ),
	}
}

// AddMQ add a MQ
func (a *App) AddMQ(m *MQ) {
	a.Lock()
	a.Mqs[m.Topic] = m
	a.Unlock()
}

// NewMQ initialize a new mq
func NewMQ(t string) *MQ {
	return &MQ{
		Config:      util.GetConfig(),
		Topic:       t,
		MessageChan: make(chan util.Message, 1000),
	}
}

// Run starts ...
func (a *App) Run() {
	log.Success("App running ......")
	for _, m := range a.Mqs {
		wg.Add(1)
		go m.run()
	}
	wg.Wait()
}

func (m *MQ) run() {
	log.Success("MQ running... topic is %s.", m.Topic)
	m.Wg.Add(2)
	go m.listenUDP()
	go m.listenHTTP()
	m.Wg.Wait()
	wg.Done()
}

func (m *MQ) listenUDP() {
	defer m.Wg.Done()
	addr, _ := net.ResolveUDPAddr("udp", m.Config.Ports.UDP)
	conn, _ := net.ListenUDP("udp", addr)
	log.Info("UDP Server listen on %s", addr)
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		n, _, _ := conn.ReadFromUDP(buf)
		message := util.TransByteToMessage(buf[0:n])
		m.MessageChan <- message
		log.Info("Send message to channel: %s", message)
		time.Sleep(1 * time.Millisecond)
	}
}

func (m *MQ) listenHTTP() {
	defer m.Wg.Done()
	log.Info("HTTP Server listen on %s", m.Config.Ports.HTTP)
	var msgs []util.Message
	http.HandleFunc("/"+m.Topic, func(writer http.ResponseWriter, request *http.Request) {
		for index := 0; index < len(m.MessageChan); index++ {
			msgs = append(msgs, <-m.MessageChan)
		}
		for i, v := range msgs {
			fmt.Fprintln(writer, fmt.Sprintf("Receive message : [%d] ===> %s.", i, v))
		}
	})
	err := http.ListenAndServe(m.Config.Ports.HTTP, nil)
	if err != nil {
		log.Error("Listening on %s failed.", m.Config.Ports.HTTP)
		return
	}
}
