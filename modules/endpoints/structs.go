package endpoints

type Error struct {
	ErrorCode    uint64 `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}

type ConnectionSuccess struct {
	ID string `json:"ID"`
}

type HeartBeatSuccess struct {
	Time any `json:"time"`
}

type EventFoundReturn struct {
	Recipient string `json:"recipient"`
	EventType string `json:"event_type"`
	Extra     string `json:"extra"`
}
