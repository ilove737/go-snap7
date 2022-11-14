# go-snap7


```go
package main

import (
	"fmt"

	s7 "github.com/ilove737/go-snap7"
)

func main() {
	Client := s7.Cli_Create()
	Client.Cli_ConnectTo("127.0.0.1", 0, 2)
	CP := Client.Cli_GetCpInfo()
	fmt.Println(CP.MaxPduLengt)

	CPU := Client.Cli_GetCpuInfo()
	fmt.Println(CPU.ModuleTypeName)

	fmt.Println(Client.Cli_GetExecTime())

	var wdata []byte
	for i := 0; i < 48; i++ {
		wdata = append(wdata, byte(i))
	}

	Client.Cli_DBWrite(1, 0, 48, wdata)

	fmt.Println(Client.Cli_DBRead(1, 1, 32))

	Client.Cli_MBWrite(1, 20, wdata)
	fmt.Println(Client.Cli_MBRead(1, 20))
}

```
