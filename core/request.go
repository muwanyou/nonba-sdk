package core

import (
	"context"
	"fmt"
	"net/url"

	"github.com/muwanyou/nonba-sdk/enum"
)

type Request interface {
	SetContext(context.Context) *BaseRequest
	SetScheme(enum.Scheme) *BaseRequest
	SetHost(string) *BaseRequest
	SetPort(int32) *BaseRequest
	SetVersion(string) *BaseRequest
	SetPath(string) *BaseRequest
	SetMethod(enum.Method) *BaseRequest
	SetHeader(map[string]string) *BaseRequest
	SetQuery(map[string]string) *BaseRequest
	SetBody([]byte) *BaseRequest
	GetContext() context.Context
	GetScheme() enum.Scheme
	GetHost() string
	GetPort() int32
	GetVersion() string
	GetPath() string
	GetUrl() string
	GetMethod() enum.Method
	GetHeader() map[string]string
	GetQuery() map[string]string
	GetBody() []byte
}

type BaseRequest struct {
	context context.Context
	scheme  enum.Scheme
	host    string
	port    int32
	version string
	path    string
	method  enum.Method
	header  map[string]string
	query   map[string]string
	body    []byte
}

func NewBaseRequest() *BaseRequest {
	return new(BaseRequest)
}

func (r *BaseRequest) Init() *BaseRequest {
	r.header = make(map[string]string)
	r.query = make(map[string]string)
	r.body = make([]byte, 0)
	return r
}

func (r *BaseRequest) SetContext(context context.Context) *BaseRequest {
	r.context = context
	return r
}

func (r *BaseRequest) SetScheme(scheme enum.Scheme) *BaseRequest {
	r.scheme = scheme
	return r
}

func (r *BaseRequest) SetHost(host string) *BaseRequest {
	r.host = host
	return r
}

func (r *BaseRequest) SetPort(port int32) *BaseRequest {
	r.port = port
	return r
}

func (r *BaseRequest) SetPath(path string) *BaseRequest {
	r.path = path
	return r
}

func (r *BaseRequest) SetVersion(version string) *BaseRequest {
	r.version = version
	return r
}

func (r *BaseRequest) SetMethod(method enum.Method) *BaseRequest {
	r.method = method
	return r
}

func (r *BaseRequest) SetHeader(header map[string]string) *BaseRequest {
	r.header = header
	return r
}

func (r *BaseRequest) SetQuery(query map[string]string) *BaseRequest {
	r.query = query
	return r
}

func (r *BaseRequest) SetBody(body []byte) *BaseRequest {
	r.body = body
	return r
}

func (r *BaseRequest) GetContext() context.Context {
	return r.context
}

func (r *BaseRequest) GetScheme() enum.Scheme {
	return r.scheme
}

func (r *BaseRequest) GetHost() string {
	return r.host
}

func (r *BaseRequest) GetPort() int32 {
	return r.port
}

func (r *BaseRequest) GetVersion() string {
	return r.version
}

func (r *BaseRequest) GetPath() string {
	return r.path
}

func (r *BaseRequest) GetMethod() enum.Method {
	return r.method
}

func (r *BaseRequest) GetHeader() map[string]string {
	return r.header
}

func (r *BaseRequest) GetQuery() map[string]string {
	return r.query
}

func (r *BaseRequest) GetBody() []byte {
	return r.body
}

func (r *BaseRequest) WithScheme(scheme enum.Scheme) *BaseRequest {
	r.scheme = scheme
	return r
}

func (r *BaseRequest) WithHost(host string) *BaseRequest {
	r.host = host
	return r
}

func (r *BaseRequest) WithPort(port int32) *BaseRequest {
	r.port = port
	return r
}

func (r *BaseRequest) WithVersion(version string) *BaseRequest {
	r.version = version
	return r
}

func (r *BaseRequest) GetUrl() string {
	if len(r.query) > 0 {
		return fmt.Sprintf("%s://%s:%d/v%s%s?%s", r.scheme, r.host, r.port, r.version, r.path, UrlQueriesEncoded(r.query))
	} else {
		return fmt.Sprintf("%s://%s:%d/v%s%s", r.scheme, r.host, r.port, r.version, r.path)
	}
}

func UrlQueriesEncoded(query map[string]string) string {
	values := url.Values{}
	for key, value := range query {
		values.Add(key, value)
	}
	return values.Encode()
}
