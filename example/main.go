package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/maitredede/go-pigpiod"
)

var (
	piaddr string
)

func main() {
	flag.StringVar(&piaddr, "pi-address", "raspberrypi.local:8888", "Raspberry pi address")
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	c, err := pigpiod.Connect(ctx, piaddr)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	//read a gpio
	const gpio = 23
	level, err := c.Read(gpio)
	if err != nil {
		panic(err)
	}
	fmt.Printf("GPIO %v read level %v", gpio, level)
}
