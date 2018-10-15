package main

import (
	"fmt"
	"time"

	rpc "github.com/superbizonsdevelopment/bizonrpc"
)

func main() {
	win, err := rpc.New("ID")
	win.Open()
	fmt.Println(win.Connection.IsOpen())

	if err != nil {
		panic(err)
	}

	str, err := win.Connection.Read()
	fmt.Println("ERR", err)
	fmt.Println("STR", str)

	time.Sleep(5 * time.Second)

	for {
		activity := &rpc.Activity{
			Details: "Example Details",
			State:   "Go is the best programming language!",
			TimeStamps: &rpc.TimeStamps{
				StartTimestamp: time.Now().Unix(),
				EndTimestamp:   time.Now().Add(20 * time.Second).Unix(),
			},
		}

		err := win.SetRichPresence(activity)
		if err != nil {
			fmt.Println(err)
			continue
		}

		str, err := win.Connection.Read()

		fmt.Println(err)
		fmt.Println("STR1", str)

		fmt.Println("---\nDone?")
		time.Sleep(time.Second * 20)
	}
}
