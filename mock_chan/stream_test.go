package main

import (
	"github.com/golang/mock/gomock"
	mocks "mock_chan/mock"
	"testing"
)

func TestChannelReturn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockObj := mocks.NewMockDataStreamer(ctrl)
	mockChan := make(chan int, 1)

	// Setup expectation
	mockObj.EXPECT().DataStream().Return(mockChan)

	// Test function that uses the DataStream method
	go func() {
		mockChan <- 42
		close(mockChan)
	}()

	// Read from the channel returned by the mock
	for val := range mockObj.DataStream() {
		if val != 42 {
			t.Errorf("Expected 42, got %d", val)
		}
	}
}
