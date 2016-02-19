package main

import (
    "./check"
    "net/http"
    "strconv"
    "encoding/json"
)

type ResponseBody struct {
    Accepted string
    ErrorType string
    ErrorMsg string
    MemoryUse int
    TimeUse int
}

type Response struct {
    Code int
    Msg string
    Body ResponseBody
}

func judge(w http.ResponseWriter, req *http.Request) {
    resp := Response{
        Code:   0,
        Msg:    "ok",
        Body:   ResponseBody{
            Accepted:   "false",
            ErrorType:      "",
            ErrorMsg:       "",
            MemoryUse:      0,
            TimeUse:        0,
        },
    }
    defer func() {
        re,_:= json.Marshal(resp)
        w.Write(re)
    }()
    req.ParseForm()
    problemIds, ok := req.Form["problemId"]
    if !ok {
        resp.Code = -1
        resp.Msg = "problemId missing!"
        return
    }
    problemIdS := problemIds[0]
    submissions, ok := req.Form["submission"]
    if !ok {
        resp.Code = -1
        resp.Msg = "submission missing!"
        return
    }
    submission := submissions[0]
    apis, ok := req.Form["api"]
    if !ok {
        resp.Code = -1
        resp.Msg = "api missing!"
        return
    }
    check.Api = apis[0]
    compilers, ok := req.Form["compiler"]
    if !ok {
        resp.Code = -1
        resp.Msg = "api missing!"
        return
    }
    compiler := compilers[0]
    
    problemId,err := strconv.Atoi(problemIdS)
    if err != nil {
        resp.Code = -1
        resp.Msg = "problemId is not number!"
        return
    }
    var checker check.Checker
    if compiler == "gcc" {
        checker = check.NewCheckC(problemId, submission)
    } else {
        resp.Code = -1
        resp.Msg = "unsupported compiler!"
        return
    }
    err = checker.Build()
    if err != nil {
        resp.Body.ErrorType = "Compiler Error"
        resp.Body.ErrorMsg  = err.Error()
        return
    }
    runinfo, err := checker.Check()
    if err != nil {
        resp.Body.ErrorType = "Runtime Error"
        resp.Body.ErrorMsg  = err.Error()
        resp.Body.MemoryUse = runinfo.Memory
        resp.Body.TimeUse   = runinfo.Time
        return
    }
    resp.Body.Accepted = "true"
    resp.Body.MemoryUse = runinfo.Memory
    resp.Body.TimeUse   = runinfo.Time
}

func main() {
    http.HandleFunc("/judge", judge)
    http.ListenAndServe(":1996", nil)
}