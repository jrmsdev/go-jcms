package response

import (
	"bytes"
	"io"
	"net/http"
)

type Response struct {
	buf    bytes.Buffer
	body   io.Writer
	size   int
	status int
	errmsg string
}

func New() *Response {
	r := &Response{}
	r.body = io.MultiWriter(&r.buf)
	r.size = 0
	r.status = http.StatusNotImplemented
	r.errmsg = "NOERRMSG"
	return r
}

func (r *Response) SetStatus(status int) {
	r.status = status
}

func (r *Response) Status() int {
	return r.status
}

func (r *Response) SetError(status int, msg string) {
	r.status = status
	r.errmsg = msg
}

func (r *Response) Error() string {
	return r.errmsg
}

func (r *Response) Write(s string) error {
	n, err := io.WriteString(r.body, s)
	if err != nil {
		r.size += n
	}
	return err
}

func (r *Response) Body() []byte {
	b := r.buf.Bytes()
	r.buf.Reset()
	return b
}
