package zmq

import (
	"testing"
	"fmt"
	"time"
	"github.com/pebbe/zmq4"
)

func TestZMQ0(t *testing.T)  {
	//  Socket to talk to clients
	responder, _ := zmq4.NewSocket(zmq4.REP)
	defer responder.Close()
	responder.Bind("tcp://*:5555")

	for {
		//  Wait for next request from client
		msg, _ := responder.Recv(0)
		fmt.Println("Received ", msg)

		//  Do some 'work'
		time.Sleep(time.Second)

		//  Send reply back to client
		reply := "World"
		responder.Send(reply, 0)
		fmt.Println("Sent ", reply)
	}
}

func TestZMQ1(t *testing.T) {
	//  Socket to talk to server
	fmt.Println("Connecting to hello world server...")
	requester, _ := zmq4.NewSocket(zmq4.REQ)
	defer requester.Close()
	requester.Connect("tcp://localhost:5555")

	for request_nbr := 0; request_nbr != 10; request_nbr++ {
		// send hello
		msg := fmt.Sprintf("Hello %d", request_nbr)
		fmt.Println("Sending ", msg)
		requester.Send(msg, 0)

		// Wait for reply:
		reply, _ := requester.Recv(0)
		fmt.Println("Received ", reply)
	}
}

func TestZMQ2(t *testing.T) {
	major, minor, patch := zmq4.Version()
	fmt.Printf("Current 0MQ version is %d.%d.%d\n", major, minor, patch)
}
