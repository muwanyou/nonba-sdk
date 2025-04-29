package enum

import (
	"net/http"
)

type Scheme string

const (
	SchemeHttp  Scheme = "http"
	SchemeHttps Scheme = "https"
)

const DefaultScheme = SchemeHttp

func (s Scheme) String() string {
	return string(s)
}

type Method string

const (
	MethodGet     Method = http.MethodGet
	MethodPost    Method = http.MethodPost
	MethodPut     Method = http.MethodPut
	MethodDelete  Method = http.MethodDelete
	MethodHead    Method = http.MethodHead
	MethodPatch   Method = http.MethodPatch
	MethodOptions Method = http.MethodOptions
)

const DefaultMethod = MethodGet

func (m Method) String() string {
	return string(m)
}
