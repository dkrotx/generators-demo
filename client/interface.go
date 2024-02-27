package client

//go:generate gowrap gen -g -p github.com/dkrotx/generators-demo/client -i Client -t withlog.tmpl -o generated_withlog.go

type Client interface {
	CallWithString(s string) string
	CallWithInt(x int) int
}
