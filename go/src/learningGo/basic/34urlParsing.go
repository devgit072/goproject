package main

import (
	"net/url"
	"fmt"
	"net"
)

func main() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u,err := url.Parse(s)
	fmt.Println(u, err)
	fmt.Println(url.Parse("www.google.com/999%22?122"))
	fmt.Println(url.Parse(".google.com/999%22?122"))
	fmt.Println(u.User)
	fmt.Println(u.User.Password())
	fmt.Println(u.User.Username())
	fmt.Println(u.Host)//it contains both hostname and port
	fmt.Println(u.Hostname())
	host,port,err := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)
	fmt.Println(u.Path)
	fmt.Println(u.RawQuery)//it will print k=v
}
