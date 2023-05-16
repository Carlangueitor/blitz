package blitz

//go:generate mockgen -destination=./mocks/mock_server.go -package=mocks github.com/carlangueitor/blitz Server

type Server interface {
	SetConfig(*Config)
	Start() error
}
