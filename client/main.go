package main

import (
	"fmt"
	"net"
	"context"
)

// This is just boilerplate grabbed from some gist.

func GoogleDNSDialer(ctx context.Context, network, address string) (net.Conn, error) {
	d := net.Dialer{}
	return d.DialContext(ctx, "udp", "localhost:8089")
}

func main() {
	r := net.Resolver{
		PreferGo: true,
		Dial: GoogleDNSDialer,
	}
	ctx := context.Background()
	ipaddr, err := r.LookupIPAddr(ctx, "google.com")
	if err != nil {
	panic(err)
	}
	fmt.Println("DNS Result", ipaddr)
}

