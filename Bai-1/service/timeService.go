package service

import (
	"awesomeProject6/model"
	"time"
)

type TimeService struct{}

func NewTimeService() *TimeService {
	return &TimeService{}
}

func (s *TimeService) GetTime() *model.Model {
	return &model.Model{
		Status:  "OK",
		Version: "v1.0.0",
		Time:    time.Now().Format(time.RFC3339),
	}
}
