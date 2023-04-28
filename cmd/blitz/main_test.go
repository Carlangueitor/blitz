package main

import (
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/carlangueitor/blitz"
	"github.com/carlangueitor/blitz/mocks"
)

func TestStart(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	configLoader := mocks.NewMockConfigLoader(mockCtl)
	config := &blitz.Config{
		Port: 1000,
	}

	configLoader.
		EXPECT().
		Load().
		Return(config, nil).
		Times(1)

	start(configLoader)
}
