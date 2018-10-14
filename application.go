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
	return &Application{ID: applicationID, System: GetSystem(), TempPath: GetTempPath(), Connected: false}
}
