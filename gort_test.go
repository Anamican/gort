package gort

import (
	"fmt"
	"net"
	"testing"
)

func TestGetFreePort(t *testing.T) {
	port, err := GetFreePort()
	if err != nil{
		t.Error(err)
	}

	// Check if the port is valid
	if port == 0 {
		t.Error("Invalid Port")
	}

	// Try to listen on the port
	listener, err := net.Listen("tcp", fmt.Sprintf("[::]:%d", port))
	if err != nil {
		t.Error(err)
	}
	defer listener.Close()

}