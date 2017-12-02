// Package channel provides mq service via golang channel
package channel

import (
	"github.com/chlins/Gomq/log"
)

// MQC global var
var MQC *MqC

// MqC is the message queue via channel
type MqC struct {
	group map[string]*Mq
}

// Mq is struct of msg queue
type Mq struct {
	MsgQueue chan string
	Cap      int
}

// NewMq struct of queue
func NewMq(c int) *Mq {
	return &Mq{
		MsgQueue: make(chan string, c),
		Cap:      c,
	}
}

// NewMqC is constructor of MqC
func NewMqC() *MqC {
	return &MqC{
		group: make(map[string]*Mq),
	}
}

// AddMq add mq to MqC
func AddMq(topic string, mq *Mq) {
	if _, found := MQC.group[topic]; !found {
		MQC.group[topic] = mq
	}
	return
}

// Push a new msg
func (m *MqC) Push(t string, v string) {
	if mq, found := MQC.group[t]; found {
		if len(mq.MsgQueue) >= mq.Cap {
			log.Error("[x] Channel is over limit")
			return
		}
		mq.MsgQueue <- v
		log.Info("[+] New message push: %s", v)
	}
}

// Pop a msg
func (m *MqC) Pop(t string) string {
	if mq, found := MQC.group[t]; found {
		if len(mq.MsgQueue) == 0 {
			log.Error("[x] Channel is empty")
			return ""
		}
		msg := <-mq.MsgQueue
		log.Info("[-] Pop message: %s", msg)
		return msg
	}
	return ""
}

// Empty check empty
func (m *MqC) Empty(t string) bool {
	return len(m.group[t].MsgQueue) == 0
}

// Full check channel full
func (m *MqC) Full(t string) bool {
	return len(m.group[t].MsgQueue) == m.group[t].Cap
}
