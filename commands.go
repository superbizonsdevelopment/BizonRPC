package bizonrpc

type AuthorizationCommand struct {
	Version       int    `json:"v"`
	ApplicationID string `json:"client_id"`
}
