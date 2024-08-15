package main

import "gorm.io/gorm"

type ApproveNodeStatusOrm struct {
	Id   uint   `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(20)"`
}

type StateMachineModelOrm struct {
	Id      uint
	Version string
	BizId   uint
	TraceId uint
	Node    string
	Content string
}

type StateMachineContextOrm struct {
	gorm.Model
	Id uint
	// -1: not start yet.
	// node.length + 1: already terminated
	CurrentNode int
	Status      StateMachineStatus
}

type ApproveNodeContextOrm struct {
	gorm.Model
	Index    uint
	Status   NodeStatus
	Approver *string
}

type StateMachineModelContextRef struct {
	StateMachineModelId          uint
	StateMachineExecuteContextId uint
}

type StateMachineContextNodeContextRef struct {
	StateMachineExecuteContextId uint
	ApproveNodeModelId           []int
}
