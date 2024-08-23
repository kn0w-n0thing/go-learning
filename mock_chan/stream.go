package main

type DataStreamer interface {
	DataStream() chan int
}

type MockStreamer struct {
	stream chan int
}

func (m *MockStreamer) DataStream() chan int {
	return m.stream
}
