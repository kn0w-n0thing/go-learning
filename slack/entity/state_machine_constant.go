package main

type PredicateType int

const (
	PredicateAll PredicateType = iota
	PredicateAny
)

type NodeStatus int

const (
	NodeStatusApproved NodeStatus = iota
	NodeStatusRejected
	NodeStatusWaiting
)

type StateMachineStatus int

const (
	StateMachineStatusApproved StateMachineStatus = iota
	StateMachineStatusRejected
	StateMachineStatusWaiting
	StateMachineStatusExpired
)
