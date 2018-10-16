package discordrpc

type HandshakeMessage struct {
	Version       int    `json:"v"`
	ApplicationID string `json:"client_id"`
}

type CommandMessage struct {
	Command string `json:"cmd"`
}

type CommandEventMessage struct {
	CommandMessage
	Event string `json:"evt"`
}

type CommandGetMessage struct {
	Nonce
	CommandMessage
	Args string `json:"args"`
}

type CommandRichPresenceMessage struct {
	Nonce
	CommandMessage
	Args *RichPresenceMessageArgs `json:"args"`
}
