package main

import (
	"flag"
	"log"
	"strconv"
	"time"
)

var ()

func init() {
}

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		log.Fatalln("1st arg is milli sleep second...")
	}

	msstr := flag.Arg(0)
	ms, err := strconv.Atoi(msstr)
	if err != nil {
		log.Fatalln(err)
	}

	time.Sleep(time.Duration(ms) * time.Millisecond)
}
