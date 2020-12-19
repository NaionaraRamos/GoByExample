package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	p := fmt.Println

	p(u.Scheme)
	p(u.User)
	p(u.User.Username())
	f, _ := u.User.Password()
	p(f)

	p(u.Host)

	host, port, _ := net.SplitHostPort(u.Host)
	p(host)
	p(port)

	p(u.Path)
	p(u.Fragment)

	p(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	p(m)
	p(m["k"][0])
}