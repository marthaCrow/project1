package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

func read(conn net.Conn) {
	//TODO In a continuous loop, read a message from the server and display it.
	reader := bufio.NewReader(conn)
	for {
		msg, _ := reader.ReadString('\n')
		fmt.Println(msg)
	}
}

func write(conn net.Conn, timer chan<- bool) {
	stdin := bufio.NewReader((os.Stdin)) // reader for keyboard input
	//TODO Continually get input from the user and send messages to the server.
	for {
		fmt.Printf("Enter text: ")
		msg, _ := stdin.ReadString('\n')
		if strings.TrimSpace(msg) == "stop" {timer <- true}
		fmt.Fprintln(conn, msg)
	}
}

func main() {
	// Get the server address and port from the commandline arguments.
	addrPtr := flag.String("ip", "127.0.0.1:8030", "IP:port string to connect to")
	flag.Parse()

	//connects to tcp server
	conn, _ := net.Dial("tcp", *addrPtr)
	timer := make(chan bool)
	go write(conn, timer)
	go read(conn)

	<-timer

	//TODO Try to connect to the server
	//TODO Start asynchronously reading and displaying messages
	//TODO Start getting and sending user messages.
}
