package sugusama

type NotOKError struct {
	At      string
	Status  string
	Message string
}

func NotOK(at, stat, msg string) error {
	return NotOKError{
		At:      at,
		Status:  stat,
		Message: msg,
	}
}

func (n NotOKError) Error() string {
	return n.At + " status " + n.Status + ", not ok: " + n.Message
}
