package main

import (
    "net/http"
)

func problemData(w http.ResponseWriter, req *http.Request) {
    w.Write([]byte("{\"code\":0,\"msg\":\"ok\",\"body\":{\"stdin\":\"1\",\"stdout\":\"1\\n1\"}}"))
}

func main() {
    http.HandleFunc("/problem/2", problemData)
    http.ListenAndServe(":1234", nil)
}