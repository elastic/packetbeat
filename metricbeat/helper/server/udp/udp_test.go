package udp

import (
	"fmt"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/metricbeat/helper/server"
)

func GetTestUdpServer(host string, port int) (server.Server, error) {
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", host, port))

	if err != nil {
		return nil, err
	}

	listener, err := net.ListenUDP("udp", addr)
	if err != nil {
		return nil, err
	}

	logp.Info("Started listening for UDP on: %s:%d", host, port)
	return &UdpServer{
		listener:          listener,
		receiveBufferSize: 1024,
		done:              make(chan struct{}),
		eventQueue:        make(chan server.Event),
	}, nil
}

func TestUdpServer(t *testing.T) {
	host := "127.0.0.1"
	port := 2003
	svc, err := GetTestUdpServer(host, port)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	svc.Start()
	defer svc.Stop()
	writeToServer(t, "test1", host, port)
	msg := <-svc.GetEvents()

	assert.True(t, msg.GetEvent() != nil)
	ok, _ := msg.GetEvent().HasKey("data")
	assert.True(t, ok)
	bytes, _ := msg.GetEvent()["data"].([]byte)
	assert.True(t, string(bytes) == "test1")

}

func writeToServer(t *testing.T, message, host string, port int) {
	servAddr := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.Dial("udp", servAddr)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	defer conn.Close()
	_, err = conn.Write([]byte(message))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
