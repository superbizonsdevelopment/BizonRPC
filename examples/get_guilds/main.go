package main

import (
	"fmt"
	"log"

	rpc "github.com/superbizonsdevelopment/bizonrpc"
)

func main() {

	win, err := rpc.New("430157626546847759")
	win.Open()

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

	err := win.Authenticate("gcUTs_kUbunc2ByWRWH0UNnmLGaIwJlh")

	if err != nil {
		log.Println("Error: ", err.Error())
		return
	}

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
