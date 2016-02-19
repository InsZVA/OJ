package judge

import (
    "os/exec"
    "bytes"
    "fmt"
)

type JudgeC struct {
    Judge
}

func (c *JudgeC) Init() {
    cmd := exec.Command("/bin/bash", "-c", `"lxc launch images:ubuntu/trusty/amd64 judgeC"`)
    var buffer bytes.Buffer
    cmd.Stdout = &buffer
    err := cmd.Run()
    if err != nil {
        panic(err)
    }
    fmt.Println(buffer.String())
}