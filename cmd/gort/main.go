package main

import (
	"fmt"
	"github.com/gort"
)

func main(){
	tcpPort, err := gort.GetFreePort()
	if err != nil{
		fmt.Println(fmt.Sprintf("Fatal error: %s", err.Error()))
	}
	fmt.Println(tcpPort)
}