package main

type StateMachineModel struct {
	Version string             `json:"version"`
	BizId   uint               `json:"bizId"`
	TraceId uint               `json:"traceId"`
	Node    []ApproveNodeModel `json:"nodes"`
}

type ApproveNodeModel struct {
	Approvers []string      `json:"approvers"`
	Predicate PredicateType `json:"predicate"`
}
