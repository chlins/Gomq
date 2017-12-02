// Package mq provides message queue service
package mq

// MsQ is common struct
type MsQ interface {
	Push(string, string)
	Pop(string) string
	Empty(string) bool
	Full(string) bool
}
