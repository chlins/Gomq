// Package service includes all type of services
package service

import (
	"github.com/chlins/Gomq/log"
)

type servicer interface {
	Start()
}

// Service common
type Service struct {
	childService []servicer
}

// NewService is the constructor of service
func NewService() *Service {
	return &Service{
		childService: make([]servicer, 0),
	}
}

// AddService add a child service
func (s *Service) AddService(svc servicer) {
	s.childService = append(s.childService, svc)
}

// Start main service
func (s *Service) Start() {
	log.Success("[Main] main service starting ...")
	for _, cs := range s.childService {
		go cs.Start()
	}
}
