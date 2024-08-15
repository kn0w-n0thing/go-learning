package main

import "time"

type SlackLeaveRequestRecord struct {
	Id         uint `gorm:"primaryKey"`
	TraceId    uint // combined to state machine
	RequestId  uint // combined to leave request
	MessageId  string
	ChannelId  string
	CreatedAt  time.Time
	ResponseAt *time.Time
	StatusId   uint
}
