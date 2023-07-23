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