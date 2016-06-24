package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"github.com/go-stomp/stomp"
)

var server string
var username string
var password string
var destname string
var messagefile string

func init() {
	flag.StringVar(&server, "server", "localhost", "The address of the server")
	flag.StringVar(&username, "username", "guest", "The address of the server")
	flag.StringVar(&password, "password", "guest", "The password used to ")
	flag.StringVar(&destname, "destination", "", "The name of the destination on the server")
	flag.StringVar(&messagefile, "file", "", "The name of the payload to send to the server")

}
func main() {
	fmt.Println("Sending message to server")
	flag.Parse()
	if destname == "" {
		fmt.Printf("Destination is required\n")
		return
	}
	if messagefile == "" {
		fmt.Printf("Payload file is required\n")
		return
	}
	fmt.Println("Ready to connect and send")
	conn,err := stomp.Dial("tcp",server,
	stomp.ConnOpt.Login(username,password))

	if err != nil {
		fmt.Printf("Error connecting %v",err)
		return
	}

	payload,err := ioutil.ReadFile(messagefile)
	if err != nil {
		fmt.Printf("Error reading %s %v\n",messagefile,err)
		return
	}

	err = conn.Send(destname,"application/binary",payload)
	if err != nil {
		fmt.Printf("Error sending content %v",err)
		return
	}
	conn.Disconnect()
}
