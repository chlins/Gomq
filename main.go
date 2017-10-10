package main

import (
	"github.com/zcytop/Gomq/app"
)

func main() {
	App := app.Init()
	mq := app.NewMQ("testTopic")
	App.AddMQ(mq)
	App.Run()
}
