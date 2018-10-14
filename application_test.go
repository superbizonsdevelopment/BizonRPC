package bizonrpc_test

import (
	"log"
	"testing"
	"time"

	rpc "github.com/superbizonsdevelopment/bizonrpc"
)

const appID = "430157626546847759"

func TestApplication(t *testing.T) {
	app := rpc.New(appID)

	err := app.Connect()

	if err != nil {
		log.Println("Error: ", err.Error())
		return
	}

	duration, err := time.ParseDuration("1m30s")

	if err != nil {
		log.Println("Error: ", err.Error())
		return
	}

	time.Sleep(duration)

	err = app.Disconnect()

	if err != nil {
		log.Println("Error: ", err.Error())
		return
	}
}
