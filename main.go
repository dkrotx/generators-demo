package main

import (
	"fmt"
	"os"

	"github.com/dkrotx/generators-demo/client"
)

func useClient(cli client.Client) {
	cli.CallWithString("Jan")
	cli.CallWithInt(50123574)
}

func main() {
	cli := client.NewClientWithLog(os.Stdout)

	fmt.Println("Hello")
	useClient(cli)
}
