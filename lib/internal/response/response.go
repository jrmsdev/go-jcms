package response

import (
	"bytes"
	"io"
	"net/http"
)

type Response struct {
	buf     bytes.Buffer
	body    io.Writer
	size    int
	status  int
	errmsg  string
	headers map[string]string
}

func New() *Response {
	r := &Response{
		size:    0,
		status:  http.StatusNotImplemented,
		errmsg:  "NOERRMSG",
		headers: make(map[string]string),
	}
	r.body = io.MultiWriter(&r.buf)
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
	// TODO: resp SetError should call r.buf.Reset() for cleanup?
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

func (r *Response) Headers() map[string]string {
	return r.headers
}

func (r *Response) SetHeader(name, value string) {
	r.headers[name] = value
}
