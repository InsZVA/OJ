package check

import (
    "io"
    "net/http"
    "strconv"
    "io/ioutil"
    "encoding/json"
    "errors"
    "bytes"
)

var Api = "http://localhost:1234/problem/"

type StdCheck struct {
    stdin io.Reader
    stdout string
    Submission string
    executable []string
    ProblemId int
}

type RunInfo struct {
    Time int    //ms
    Memory int  //kb
}

type Checker interface {
    Build() error
    Check() (RunInfo, error)
    GetStandardInOut() error
}

func (this *StdCheck) GetStandardInOut() error {
    response, err := http.Get(Api + strconv.Itoa(this.ProblemId))
    defer response.Body.Close()
    if err != nil {
        return err
    }
    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return err
    }
    var jsonData map[string]interface{}
    err = json.Unmarshal(body, &jsonData)
    if err != nil {
        return err
    }
    if code, ok := jsonData["code"].(int);ok && code != 0 {
        msg, ok := jsonData["msg"].(string)
        if ok {
            return errors.New(msg)
        } else {
            return errors.New("In GetStandardInOut: Response code is not 0!")
        }
    }
    jsonBody, ok := jsonData["body"].(map[string]interface{})
    if !ok {
        return errors.New("In GetStandardInOut: Response body is missing!")
    }
    stdin, ok := jsonBody["stdin"].(string)
    if !ok {
        return errors.New("In GetStandardInOut: Response stdin is missing!")
    }
    stdout, ok := jsonBody["stdout"].(string)
    if !ok {
        return errors.New("In GetStandardInOut: Response stdout is missing!")
    }
    this.stdin = bytes.NewReader([]byte(stdin))
    this.stdout = stdout
    return nil
}

