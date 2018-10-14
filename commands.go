package bizonrpc

type Command struct {
	Cmd string `json:"cmd"`
}

type AuthorizationCommand struct {
	Cmd           string `json:"cmd"`
	Version       int    `json:"v"`
	ApplicationID string `json:"client_id"`
}

type SetRPCCommand struct {
	Cmd  string      `json:"cmd"`
	Args *RPCMsgArgs `json:"args"`
}

/*
  Under are args for Commands
*/

type RPCMsgArgs struct {
	Pid      int       `json:"pid"`
	Activity *Activity `json:"activity"`
}

type Activity struct {
	State    string `json:"state"`
	Details  string `json:"details"`
	TimeStam *Timestamps
}

type Timestamps struct {
	Start string `json:"start"`
	End   string `json:"end"`
}
