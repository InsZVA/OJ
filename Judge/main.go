package main
import (
    "./judge"
    "flag"
    "net/http"
)

var compilerCNum = flag.Int("c", 5, "C Compilers Num") 
var judgeHub *judge.JudgeHub

func handler(rw http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    go judgeHub.Check(r.Form["problemId"][0], r.Form["compiler"][0], r.Form["api"][0], r.Form["submission"][0], r.Form["notify"][0])
    rw.Write([]byte(`{"code":0,"msg":"ok","body":{}}`))
}

func main() {
    flag.Parse()
    judgeHub = judge.NewJudgeHubC(*compilerCNum)
    //TODO: Add HTTP Listen
    http.HandleFunc("/", handler)
    http.ListenAndServe(":2020", nil)
}