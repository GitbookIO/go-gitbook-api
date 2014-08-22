package client

type Error struct {
	Msg  string `json:"error"`
	Code int    `json:"code"`
}

func (e *Error) Error() string {
	return e.Msg
}
