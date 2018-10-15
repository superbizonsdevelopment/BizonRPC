package discordrpc_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	rpc "github.com/gitarzysta/discordrpc-go"
)

func TestMeme(t *testing.T) {

	//time.Sleep(time.Second * 3)

	//316245861074206730
	//368924690946850817
	win, err := rpc.New("430157626546847759")
	win.Open()
	fmt.Println(win.Connection.IsOpen())

	if err != nil {
		panic(err)
	}

	str, err := win.Connection.Read()
	fmt.Println("ERR", err)
	fmt.Println("STR", str)

	time.Sleep(5 * time.Second)
	counter := 0
	for {
		activity := &rpc.Activity{
			Details: "Dean" + strconv.Itoa(counter),
			State:   "Proud To Be A Go Developer",
			TimeStamps: &rpc.TimeStamps{
				StartTimestamp: time.Now().Unix(),
				EndTimestamp:   time.Now().Add(20 * time.Second).Unix(),
			},
		}
		counter++
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
