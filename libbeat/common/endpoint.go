package common

// Representing an endpoint in the communication.
type Endpoint struct {
	Ip      string
	Port    uint16
	Name    string
	Cmdline string
	Proc    string
}
