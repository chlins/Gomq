package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// VERSION flag
const VERSION string = "0.0.1"

var (
	help    bool
	version bool
	topic   string
	cap     int
)

func usage() {
	fmt.Fprintf(os.Stderr,
		`Usage: ./gomq [-v version] [-t topic] [-c capacity]
Options:
`)
	flag.PrintDefaults()
}

func init() {
	flag.BoolVar(&help, "h", false, "this help")
	flag.BoolVar(&version, "v", false, "show version and exit")
	flag.IntVar(&cap, "c", 1000, "set the capacity of this topic")
	flag.StringVar(&topic, "t", "", "set a topic")
	flag.Usage = usage
}

func main() {
	flag.Parse()
	if help {
		flag.Usage()
		os.Exit(0)
	}
	if version {
		fmt.Println("Gomq: ", VERSION)
		os.Exit(0)
	}
	topic = strings.Replace(topic, " ", "", -1)
	if topic == "" {
		fmt.Println("Topic can't be null, please set a topic")
		fmt.Println("[-h] see help")
		os.Exit(1)
	}
}
