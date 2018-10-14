package bizonrpc

import "net"

type Application struct {
	ID        string
	System    string
	Conn      net.Conn
	Connected bool
}

func New(applicationID string) Application {
	return &Application{ID: applicationID, System: bizonrpc.GetSystem()}
}

func (app *Application) Connect() error {
	conn, err := net.Dial()
}

func (app *Application) SetRichPresence() {

}

func (app *Application) IsConnect() {

}

func (app *Application) Disconnect() {

}
