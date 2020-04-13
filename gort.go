package gort

import (
	"fmt"
	"net"
)

// GetFreePort searches for FreePort and send back one
func GetFreePort() (int, error){
	tcpAddr, err := net.ResolveTCPAddr("tcp", "[::]:0")
	if err != nil {
		fmt.Println(fmt.Sprintf("Fatal error: %s", err.Error()))
		return 0, err
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println(fmt.Sprintf("Fatal error: %s", err.Error()))
		return 0, err
	}

	defer listener.Close()
	return listener.Addr().(*net.TCPAddr).Port, nil
}