package errs

func WrapError(e error, msg string) error {
	return &wrapError{
		err: e,
		msg: msg,
	}
}

type wrapError struct {
	err error
	msg string
}

func (e *wrapError) Error() string {
	return e.msg
}
func (e *wrapError) Unwrap() error {
	return e.err
}
