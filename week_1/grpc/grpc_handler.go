package grpc

import "github.com/fatih/color"

// Hello возвращает приветственное сообщение
func Hello() string {
	return color.GreenString("Hello from gRPC package!")
}
