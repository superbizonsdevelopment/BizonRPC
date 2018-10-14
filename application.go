package bizonrpc

import (
	"net"
	"strconv"
)

type Application struct {
	ID        string
	System    string
	TempPath  string
	Conn      net.Conn
	Connected bool
}

func New(applicationID string) Application {
	return Application{
		ID:        applicationID,
		System:    GetSystem(),
		TempPath:  GetTempPath(),
		Connected: false,
	}
}

func (app *Application) Connect() error {
	for i := 0; i < 10; i++ {
		conn, err := net.Dial(app.System, app.TempPath+"discord-rpc-"+strconv.Itoa(i))

		if err == nil {
			app.Conn = conn
			app.Connected = true
			return nil
		}
	}

	return err
}

func (app *Application) SetRichPresence() {

}

func (app *Application) IsConnect() {

}

func (app *Application) Disconnect() error {
	err := app.Conn.Close()

	if err != nil {
		return err
	}
	return nil
}
