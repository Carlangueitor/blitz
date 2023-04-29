package blitz

type Server interface {
	SetConfig(*Config)
	Start() error
}
