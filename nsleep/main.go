package main

import (
	"flag"
	"time"
)

var (
	t time.Duration
)

func init() {
	flag.DurationVar(&t, "t", time.Second, "sleep time")
}

func main() {
	flag.Parse()
	time.Sleep(t)
}
