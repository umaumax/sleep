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
		log.Fatalln("1st arg is micro sleep second...")
	}

	usstr := flag.Arg(0)
	us, err := strconv.Atoi(usstr)
	if err != nil {
		log.Fatalln(err)
	}

	time.Sleep(time.Duration(us) * time.Microsecond)
}
