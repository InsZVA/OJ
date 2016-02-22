package judge

import (
    "io/ioutil"
    "fmt"
    "time"
    "net/http"
    "net/url"
    "bytes"
)

const daemon_path = "./daemon"

type LXCer interface {
	Init(lxcTemplateName, lxcRemoteServer, lxcRemoteAlias string)
    Run()
	Restore()
    Working() bool
    SetWorking(b bool)
    IP() string
}

type JudgeHub struct {
    lxcTemplateName string      //A template lxc for copies
    lxcRemoteServer string      //Remote resources to pull when template not exist
    lxcRemoteAlias  string      //Remote image alias to pull
    lxcMaxNum       int         //Max LXC number in a JudgeHub
    lxcs            []LXCer    //LXC List
    autoIncreament  int         //Unique id to create lxc for name
    compiler        string
}

/*
{
    "type": "sync",
    "status": "Success",
    "status_code": 200,
    "metadata": {}                          # Extra resource/action specific metadata
}
*/

type Judge struct {
    containerName   string
    sanpshotName    string
    working         bool
}

func (this *Judge) Init(lxcTemplateName, lxcRemoteServer, lxcRemoteAlias string) {
    var uuid string
    fmt.Println("Initializing ", this.containerName)
    if containerExist(lxcTemplateName) {
        fmt.Println("Copying ", this.containerName)
        _, uuid = containerCopy(lxcTemplateName, this.containerName)
        wait(uuid)
        fmt.Println("Launching ", this.containerName)
        _, uuid = containerStart(this.containerName)
        wait(uuid)
        fmt.Println("Creating snapshot on ", this.containerName)
        this.sanpshotName = this.containerName + "snap"
        _, uuid = containerSnapshot(this.containerName, this.sanpshotName)
        wait(uuid)
        return
    } else {
        _, uuid =containerLaunch(this.containerName, lxcRemoteServer, lxcRemoteAlias)
    }
    fmt.Println("Launching ", this.containerName)
    wait(uuid)
    _, uuid = containerStart(this.containerName)
    fmt.Println("Starting ", this.containerName)
    wait(uuid)
    fmt.Println("Uploading daemon to ", this.containerName)
    bytes,_ := ioutil.ReadFile(daemon_path)
    containerPush(this.containerName, "/daemon", bytes)

    //TODO: subTypes add code to install compiler environment and create snapshot
}

func (this *Judge) Run() {
    _, uuid := containerStart(this.containerName)
    wait(uuid)
    _, uuid = containerExec(this.containerName, []string{"chmod", "777", "/daemon"})
    wait(uuid)
    _, uuid = containerExec(this.containerName, []string{"/daemon"})
    time.Sleep(2 * 1000 * time.Millisecond)
}

func (this *Judge) Restore() {
    _, uuid := containerRestore(this.containerName, this.sanpshotName)
    wait(uuid)
}

func (this *Judge) State() (string, []string, error) {
    return containerState(this.containerName)
}

func (this *Judge) Working() bool {
    return this.working
}

func (this *Judge) SetWorking(b bool) {
    this.working = b
}

func (this *Judge) IP() string {
    _, ips, _ := containerState(this.containerName)
    return ips[0]
}

func NewJudgeHub(lxcTemlateName, lxcRemoteServer, lxcRemoteAlias string, lxcMaxNum int) *JudgeHub {
    return &JudgeHub{
        lxcMaxNum:          lxcMaxNum,
        lxcRemoteAlias:     lxcRemoteAlias,
        lxcRemoteServer:    lxcRemoteServer,
        autoIncreament:     1,
        lxcTemplateName:    lxcTemlateName,
    }    
}


func (this *JudgeHub) Check(problemId string, compiler string, api string, submission string, notify string) {
    retry:
    //Peek a free Judge
    var selectedJudge LXCer
    for selectedJudge == nil {
        for _, judge := range this.lxcs {
            if !judge.Working() {
                selectedJudge = judge
                break
            }
        }
    }
    //TODO: notify the judge begin
    selectedJudge.SetWorking(true)
    //TODO: command Judge To Judge
    data := make(url.Values)
    data["porblemId"] = []string{problemId}
    data["compiler"] = []string{compiler}
    data["api"] = []string{api}
    data["submission"] = []string{submission}
    http.Get(notify + "/problem/" + problemId + "/start")   //notify Start
    res, err := http.PostForm("http://" + selectedJudge.IP() + ":1996/judge", data)
    if err != nil {
        //This conttainer might occur an error
        go func(selectedJudge LXCer) {
            selectedJudge.Restore()
            selectedJudge.Run()
            selectedJudge.SetWorking(false)
        }(selectedJudge)
        //Another try
        selectedJudge = nil
        goto retry
    }
    defer res.Body.Close()
    result, _ := ioutil.ReadAll(res.Body)
    //Restore selected judge
    go func(selectedJudge LXCer) {
            selectedJudge.Restore()
            selectedJudge.Run()
            selectedJudge.SetWorking(false)
        }(selectedJudge)
    //Notify the result
    http.Post(notify + "/problem/" + problemId + "/result", "application/json", bytes.NewReader(result))
}