package response

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/jrmsdev/go-jcms/lib/internal/context/appctx"
)

type Response struct {
	buf       bytes.Buffer
	body      io.Writer
	size      int
	status    int
	errmsg    string
	headers   map[string]string
	location  string
	template  string
	tpllayout string
}

func New() *Response {
	r := &Response{
		size:      0,
		status:    http.StatusNotImplemented,
		errmsg:    "NOERRMSG",
		headers:   make(map[string]string),
		location:  "NOLOCATION",
		template:  "",
		tpllayout: "",
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

func (r *Response) SetError(
	ctx context.Context,
	status int,
	msg string,
) context.Context {
	r.status = status
	r.errmsg = msg
	r.buf.Reset()
	return appctx.Fail(ctx)
}

func (r *Response) Error() string {
	return r.errmsg
}

func (r *Response) Write(blob []byte) (int, error) {
	return r.body.Write(blob)
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

func (r *Response) Location() string {
	return r.location
}

func (r *Response) Redirect(
	ctx context.Context,
	status int,
	location string,
) context.Context {
	r.status = status
	r.location = location
	return appctx.SetRedirect(ctx)
}

func (r *Response) Template() string {
	return r.template
}

func (r *Response) SetTemplate(name string) {
	r.template = name
}

func (r *Response) TemplateLayout() string {
	return r.tpllayout
}

func (r *Response) SetTemplateLayout(name string) {
	r.tpllayout = name
}
