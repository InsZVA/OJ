package unix

import (
    "net"
    "net/http"
)

var address string

func unixDial(proto, addr string) (conn net.Conn, err error) {
    if address == "" {
        address = "/var/lib/lxd/unix.socket"
    }
    return net.Dial("unix", address)
}

func NewClient(addr string) *http.Client {
    address = addr
    tr := &http.Transport{ Dial: unixDial, }
    client := &http.Client{ Transport: tr, }
    return client
}