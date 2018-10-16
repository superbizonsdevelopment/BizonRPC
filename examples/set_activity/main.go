package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	rpc "github.com/superbizonsdevelopment/bizonrpc"
)

var (
	client_id string
	details   string
	state     string
)

func init() {
	flag.StringVar(&client_id, "i", "", "RPC ID.")
	flag.StringVar(&details, "d", "", "Details for RPC.")
	flag.StringVar(&state, "s", "", "State for RPC.")
	flag.Parse()

	if client_id == "" {
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	win, err := rpc.New(client_id)
	win.Open()
	fmt.Println(win.Connection.IsOpen())

	if err != nil {
		panic(err)
	}

	str, err := win.Connection.Read()

	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}

	fmt.Println("STR:", str)

	time.Sleep(5 * time.Second)

	for {
		activity := &rpc.Activity{
			Details: details,
			State:   state,
			TimeStamps: &rpc.TimeStamps{
				StartTimestamp: time.Now().Unix(),
				EndTimestamp:   time.Now().Add(20 * time.Second).Unix(),
			},
		}

		err := win.SetRichPresence(activity)
		if err != nil {
			fmt.Println("Error: ", err.Error())
			continue
		}

		str, err := win.Connection.Read()

		if err != nil {
			fmt.Println("Error: ", err.Error())
			return
		}

		fmt.Println("STR1", str)

		fmt.Println("---\nDone?")
		time.Sleep(time.Second * 20)
	}
}
