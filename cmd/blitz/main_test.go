package main

import (
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/carlangueitor/blitz"
	"github.com/carlangueitor/blitz/mocks"
)

func TestStart(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	configLoader := mocks.NewMockConfigLoader(mockCtrl)
	config := &blitz.Config{
		Port: 1000,
	}

	configLoader.
		EXPECT().
		Load().
		Return(config, nil).
		Times(1)

	server := mocks.NewMockServer(mockCtrl)
	server.
		EXPECT().
		SetConfig(config).
		Times(1)
	server.
		EXPECT().
		Start().
		Return(nil).
		Times(1)

	start(configLoader, server)
}
