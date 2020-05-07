package sugusama

type wrapError struct {
	err error
	msg string
}

func (w wrapError) Error() string {
	return w.msg + ": " + w.err.Error()
}

func (w wrapError) Unwrap() error {
	return w.err
}

func wrap(err *error, msg string) {
	*err = wrapError{
		err: *err,
		msg: msg,
	}
}
