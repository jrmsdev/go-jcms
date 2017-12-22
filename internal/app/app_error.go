package app

type Error struct {
    status int
    msg string
}

func newError (status int, msg string) *Error {
    return &Error{status, msg}
}

func (e *Error) Status () int {
    return e.status
}

func (e *Error) Error () string {
    return e.msg
}
