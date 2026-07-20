//go:build !js

package main

import (
	"net"
	"net/url"

	"golang.org/x/net/proxy"
)

var _ proxy.Dialer = &proxyDialer{}

type proxyDialer struct {
	proxyAddr string
}

func newProxyDialer(u *url.URL) proxy.Dialer { _ = "STUB: not implemented"; return *new(proxy.Dialer) }

func (d *proxyDialer) Dial(network, addr string) (net.Conn, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func newHTTPProxy() *url.URL { _ = "STUB: not implemented"; return nil }

func proxyHandleConn(clientConn net.Conn) { _ = "STUB: not implemented"; return }
