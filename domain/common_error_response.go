package domain

type Error struct {
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}
