package bizonrpc

import (
	"encoding/json"
	"net"
	"os"
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
	var err error

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

func (app *Application) Write(data []byte) (int, error) {
	tot, err := app.Conn.Write(data)
	if err != nil {
		return tot, err
	} else if tot <= 0 {
		app.Conn.Close()
		return tot, NoDataError
	}
	return tot, nil
}

func (app *Application) SendData(cmd Command) {

}

func (app *Application) Authorize() error {
	cmd := &AuthorizationCommand{
		Cmd:           "AUTHORIZE",
		Version:       1,
		ApplicationID: "430157626546847759",
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return err
	}
	_, err := app.Conn.Write(string(data))

	if err != nil {
		return err
	}
	return nil
}

func (app *Application) SetRichPresence(activity *Activity) error {
	cmd := &SetRPCCommand{
		Cmd: "SET_ACTIVITY",
		Args: &RPCMsgArgs{
			Activity: activity,
			Pid:      os.Getpid(),
		},
	}
	data, err := json.Marshal(cmd)

	if err != nil {
		return err
	}

	return app.Connection.Write(string(data))
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
