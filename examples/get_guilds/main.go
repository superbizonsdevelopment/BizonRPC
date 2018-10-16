package main

import (
	"fmt"
	"log"
	"time"

	rpc "github.com/superbizonsdevelopment/bizonrpc"
)

func main() {

	win, err := rpc.New("")

	if err != nil {
		log.Println("Error: ", err.Error())
		return
	}

	str, err := win.Connection.Read()

	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}

	fmt.Println("STR:", str)

	time.Sleep(5 * time.Second)

	err = win.GetGuilds()

	if err != nil {
		log.Println("Error: ", err.Error())
		return
	}

	str, err = win.Connection.Read()

	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}

	fmt.Println("STR1", str)
}
