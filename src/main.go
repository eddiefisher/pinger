// Copyright 2018, Eddie Fisher. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// main.go [created: Mon,  8 Jan 2018]

package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/eddiefisher/pinger/src/version"
)

var (
	address     = "127.0.0.1:8081"
	machineName string
	timeMessage = 5 * time.Second
	timeReconn  = 3 * time.Second
)

func init() {
	logrus.Printf("commit: %s, build time: %s, release: %s", version.Commit, version.BuildTime, version.Release)
}

func main() {
	_machineName := flag.String("name", "kk", "enter name of machine")
	flag.Parse()
	machineName = *_machineName
	logrus.Println(*_machineName)
	for {
		conn, err := net.Dial("tcp", address)
		logrus.Println("open connection")
		if err != nil {
			logrus.Warnln("connection error:", err)
			time.Sleep(timeReconn)
		} else {
			err = Message(conn)
			if err != nil {
				logrus.Warnln("message error", err)
			}
		}
	}
}

//Message get and send message to server
func Message(conn net.Conn) (err error) {
	for {
		time.Sleep(1 * time.Second)
		fmt.Fprintf(conn, "text"+machineName+"\n")
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			break
		}
		fmt.Print("Message from server: " + message)
	}
	return
}
