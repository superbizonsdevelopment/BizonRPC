package bizonrpc

import "net"

type Application struct {
	ID        string
	System    string
	TempPath  string
	Conn      net.Conn
	Connected bool
}

func New(applicationID string) Application {
	return Application{
    ID: applicationID,
    System: GetSystem(),
    TempPath: GetTempPath(),
    Connected: false
  }
}

func (app *Application) Connect() error {
	conn, err := net.Dial(app.System, app.TempPath+"discord-rpc")

	if err != nil {
		return err
	}
	return nil
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
