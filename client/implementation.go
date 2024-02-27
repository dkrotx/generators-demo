package client

import "fmt"

type impl struct {
}

// NewBaseClient returns a base client version
func NewBaseClient() Client {
	return &impl{}
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func (i impl) CallWithString(s string) (string, string) {
	fmt.Printf("Implementation.CallWithString(%s)\n", s)
	return s, reverse(s)
}

func (i impl) CallWithInt(x int) int {
	fmt.Printf("Implementation.CallWithInt(%d)\n", x)
	return -x
}
