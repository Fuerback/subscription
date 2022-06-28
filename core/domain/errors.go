package domain

type Error struct {
	StatusCode int
	Message    error `json:"message"`
}

func (e *Error) Error() string {
	return e.Message.Error()
}
