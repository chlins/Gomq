package main

import (
	"os"
	"os/signal"

	"github.com/chlins/Gomq/mq/channel"
	"github.com/chlins/Gomq/service"
	"github.com/chlins/Gomq/service/tcp"
)

func main() {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, os.Kill)
	svc := service.NewService()
	channel.MQC = channel.NewMqC()
	tcpSvc := tcp.NewService("8001", channel.MQC)
	svc.AddService(tcpSvc)
	svc.Start()
	<-exit
}
