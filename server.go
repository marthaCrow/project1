package main

import (
	"bufio"
	"flag"
	"net"
	"fmt"
	"strings"
)

type Message struct {
	sender  int
	message string
}

func handleError(err error) {
	// TODO: all
	// Deal with an error event.
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func acceptConns(ln net.Listener, conns chan net.Conn) {
	// TODO: all
	
	for {
		// Continuously accept a network connection from the Listener
		conn, err := ln.Accept()
		handleError(err)
		// and add it to the channel for handling connections.
		conns <- conn
	}
}

func handleClient(client net.Conn, clientid int, msgs chan Message) {
	// TODO: all
	// So long as this connection is alive:
	// Read in new messages as delimited by '\n's
	reader := bufio.NewReader(client)

	for {
		message , err := reader.ReadString('\n')
		handleError(err)
		// Tidy up each message and add it to the messages channel,
		// recording which client it came from.
		msgs <- Message{clientid, strings.TrimSpace(message)}
	}
	
}


func main() {
	// Read in the network port we should listen on, from the commandline argument.
	// Default to port 8030
	portPtr := flag.String("port", ":8030", "port to listen on")
	flag.Parse()

	//TODO Create a Listener for TCP connections on the port given above.
	ln, err := net.Listen("tcp", *portPtr)
	handleError(err)

	//Create a channel for connections
	conns := make(chan net.Conn)
	//Create a channel for messages
	msgs := make(chan Message)
	//Create a mapping of IDs to connections
	clients := make(map[int]net.Conn)

	i := 0

	//Start accepting connections
	go acceptConns(ln, conns)
	for {
		select {
		case conn := <-conns:
			//TODO Deal with a new connection
			// - assign a client ID
			clients[i] = conn
			
			// - start to asynchronously handle messages from this client
			go handleClient(conn, i, msgs)
			fmt.Println("Client connected:", i)
			i++
		case msg := <-msgs:
			for id, conn := range clients {
				if id != msg.sender {
					fmt.Fprintln(conn, msg.message)
				}
			}
		}
	}
}
