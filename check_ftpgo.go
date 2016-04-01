package main

import (
	"flag"
	"fmt"
	"github.com/jlaffaye/ftp"
	"os"
	"time"
)

var (
	address        = flag.String("address", "localhost", "The address to connect.")
	connectTimeout = flag.Int("connect-timeout", 10000, "Connect timeout in ms.")
	password       = flag.String("password", "", "FTP password.")
	port           = flag.String("port", "21", "Port to use.")
	timeout        = flag.Int("timeout", 10000, "Timeout in ms.")
	username       = flag.String("username", "", "FTP username.")
)

func nagiosExitOk(t time.Duration) {
	fmt.Println("Ok, time:", t)
	os.Exit(0)
}

func nagiosExitError(msg string) {
	fmt.Println(msg)
	os.Exit(2)
}

func main() {

	var totalResponseTime time.Duration

	flag.Parse()

	//Set timer for global time
	t0 := time.Now()

	c, err := ftp.DialTimeout(*address+":"+*port, time.Duration(*connectTimeout)*time.Millisecond)
	if err != nil {
		nagiosExitError(err.Error())
	}

	defer c.Quit()

	if *username != "" && *password != "" {
		err = c.Login(*username, *password)
		if err != nil {
			nagiosExitError(err.Error())
		}
	}

	//Get total time
	totalResponseTime = time.Now().Sub(t0)

	//Exit OK
	nagiosExitOk(totalResponseTime)
}
