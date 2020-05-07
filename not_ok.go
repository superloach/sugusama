package sugusama

type NotOKError struct {
	At      string
	Message string
}

func NotOK(at, msg string) error {
	return NotOKError{
		At:      at,
		Message: msg,
	}
}

func (n NotOKError) Error() string {
	return n.At + " status not ok: " + n.Message
}
