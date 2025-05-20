package main

import (
	"github.com/goburrow/modbus"
	"time"
)

func main() {
	// Modbus RTU/ASCII
	handler := modbus.NewRTUClientHandler("/dev/ttyUSB0")
	handler.BaudRate = 115200
	handler.DataBits = 8
	handler.Parity = "N"
	handler.StopBits = 1
	handler.SlaveId = 1
	handler.Timeout = 5 * time.Second

	defer handler.Close()

	client := modbus.NewClient(handler)

	// 一分钟一次
	for {
		time.Sleep(1 * time.Minute)
		client.ReadDiscreteInputs(15, 2)
	}

}
