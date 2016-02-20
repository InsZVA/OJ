package main
import (
    _"./judge"
    "net/http"
    "io/ioutil"
    "fmt"
    "net"
)

func unixDial(proto, addr string) (conn net.Conn, err error) {
    return net.Dial("unix", "/var/lib/lxd/unix.socket")
}

func main() {
    //var c judge.JudgeC
    //c.Init()
    tr := &http.Transport{ Dial: unixDial, }
    client := &http.Client{ Transport: tr, }
    r, err := client.Get("http://unix.socket/1.0/containers/JudgeqC")
    if err != nil {
        panic(err)
    }
    defer r.Body.Close()
    bytes, _ := ioutil.ReadAll(r.Body)
    fmt.Println(string(bytes))
}