package main

import (
	"fmt"
	"time"
)

const (
	DefaultTimeOut time.Duration = 1 * time.Minute
	DefaultCaching bool          = false
)

type options struct {
	timeout time.Duration
	caching bool
}

type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

func WithTimeOut(t time.Duration) Option {
	return optionFunc(func(o *options) {
		o.timeout = t
	})
}

func WithCaching(c bool) Option {
	return optionFunc(func(o *options) {
		o.caching = c
	})
}

func Connect(opts ...Option) {
	options := options{
		timeout: DefaultTimeOut,
		caching: DefaultCaching,
	}
	for _, o := range opts {
		o.apply(&options)
	}
	fmt.Println(options)
}

func main() {
	Connect()
	Connect(WithTimeOut(1 * time.Microsecond))
	Connect(WithCaching(true))
	Connect(WithTimeOut(1*time.Microsecond), WithCaching(true))
}
