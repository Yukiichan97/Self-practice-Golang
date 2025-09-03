package service

import (
	"awesomeProject6/model"
	"time"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetTime() *model.Model {
	return &model.Model{
		Status:  "OK",
		Version: "v1.0.0",
		Time:    time.Now().Format(time.RFC3339),
	}
}
