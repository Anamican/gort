# gort
Free Port finder using Golang

## Import as package

```go
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
```

## As an Executable

Once the gort is built using `go build` , use the executable as

```shell script
    $ gort
    $ SOME_PORT_NUMBER_PRINTED
```