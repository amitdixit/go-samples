package main

import "fmt"

type OptsFunc func(*Opts)

type Opts struct {
	maxConn int
	id      string
	tls     bool
}

type Server struct {
	Opts
}

func defaultOpts() Opts {
	return Opts{
		maxConn: 100,
		id:      "default",
		tls:     false,
	}
}

func withTls(opts *Opts) {
	opts.tls = true
}

func withMaxConn(maxConn int) OptsFunc {
	return func(opts *Opts) {
		opts.maxConn = maxConn
	}
}

func withId(id string) OptsFunc {
	return func(opts *Opts) {
		opts.id = id
	}
}

func newServer(opts ...OptsFunc) *Server {
	o := defaultOpts()
	for _, fn := range opts {
		fn(&o)
	}

	return &Server{Opts: o}
}

func main() {
	s := newServer(
		withTls,
		withMaxConn(1000),
		withId("test"),
	)
	fmt.Printf("%#v\n", s) // &main.Opts(s)
}
