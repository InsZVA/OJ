package check

import (
    "os/exec"
    "bytes"
    "io/ioutil"
    "errors"
    "time"
    "strconv"
    "fmt"
)

type CheckC struct{
    StdCheck
    timeLimit   int //ms
    memoryLimit int //Kb
    memoryUsage int //Kb
    timeUsage   int //ms
    finished    bool     //When the checking program finished this will be true and running goroutine will end
}

func (this *CheckC) Build() error {
    if this.ProblemId == 0 {
        return errors.New("Uninitialized checker!")
    }
    var out_b bytes.Buffer
    var err_b bytes.Buffer
    err := ioutil.WriteFile("./a.c", []byte(this.Submission), 0666)
    if err != nil {
        return err
    }
    cmd := exec.Command("/usr/bin/gcc", "./a.c", "-o", "./a.o")
    cmd.Stdout = &out_b
    cmd.Stderr = &err_b
    cmd.Run()
    if err_b.String() != "" {
        return errors.New(string(err_b.String()))
    }
    this.executable = []string{"./a.o"}
    return nil;
}

func NewCheckC(problemId int, submission string) *CheckC {
    new := &CheckC {
        StdCheck:       StdCheck{
            stdin:      nil,
            stdout:     "",
            Submission: submission,
            executable: nil,
            ProblemId:  problemId, 
        },
        timeLimit:      1000,
        memoryLimit:    256 * 1024,  
        memoryUsage:    0,  
    }
    new.GetStandardInOut()
    return new
}

func getMemoryUsage(pid int) int {
    filePath := "/proc/" + strconv.Itoa(pid) + "/statm"
    data, err := ioutil.ReadFile(filePath)
    if err != nil {
        return -1
    }
    tmp := 0
    mem := 0
    els := 0
    fmt.Sscanf(string(data), "%d %d %d", &tmp, &mem, &els)
    return mem * 4  //4K per page
}

func (this *CheckC) running(cmd *exec.Cmd) {
    pid := cmd.Process.Pid
    startTime := time.Now().UnixNano()
    for !this.finished {
        mem := getMemoryUsage(pid)
        if mem > this.memoryUsage {
            this.memoryUsage = mem
        }
        if mem >= this.memoryLimit {
            cmd.Process.Kill()
            return
        }
        if int((time.Now().UnixNano() - startTime) / (1024 * 1024)) > this.timeLimit {
            cmd.Process.Kill()
            this.timeUsage = -1
            return
        }
    }
    endTime := time.Now().UnixNano()
    this.timeUsage = int((endTime - startTime) / (1024 * 1024))
}

func (this *CheckC) Check() (RunInfo, error) {
    if this.ProblemId == 0 {
        return RunInfo {
            Time:   0,
            Memory: 0,
        }, errors.New("Uninitialized checker!")
    }
    if this.executable == nil {
        return RunInfo {
            Time:   0,
            Memory: 0,
        }, errors.New("Please run build first!")
    }
    var check_out bytes.Buffer
    cmd := exec.Command(this.executable[0],this.executable[1:]...)
    cmd.Stdin = this.stdin
    cmd.Stdout = &check_out
    err := cmd.Start()
    if err != nil {
        return RunInfo {
            Time:   0,
            Memory: 0,
        }, err
    }
    go this.running(cmd)
    err = cmd.Wait()
    this.finished = true
    if err != nil {
        if this.timeUsage < 0 {
            this.timeUsage = int(cmd.ProcessState.UserTime().Nanoseconds() / (1024 * 1024))
            return RunInfo {
                Time:   this.timeUsage,
                Memory: this.memoryUsage,
            }, errors.New("Time Limit Error!")
        }
        this.timeUsage = int(cmd.ProcessState.UserTime().Nanoseconds() / (1024 * 1024))
        if this.memoryUsage < 0 {
            return RunInfo {
                Time:   this.timeUsage,
                Memory: this.memoryUsage,
            }, errors.New("Memory Limit Error!")
        }
        return RunInfo {
            Time:   this.timeUsage,
            Memory: this.memoryUsage,
        }, err
    }
    this.timeUsage = int(cmd.ProcessState.UserTime().Nanoseconds() / (1024 * 1024))
    if check_out.String() != this.stdout {
        return RunInfo {
            Time:   this.timeUsage,
            Memory: this.memoryUsage,
        },  errors.New("Wrong Answer!")
    }
    return RunInfo {
            Time:   this.timeUsage,
            Memory: this.memoryUsage,
        }, nil;
}