package util

import (
	"encoding/json"

	"github.com/zcytop/Gomq/log"
)

// Message is the type of message
type Message struct {
	Body string `json:"body"`
}

// TransByteToMessage transform from bytes to message
func TransByteToMessage(value []byte) Message {
	message := Message{}
	err := json.Unmarshal(value, &message)
	if err != nil {
		log.Error("Unable to transform from bytes to message, ", err)
	}
	return message
}
