package judge

import (
    "io/ioutil"
    "fmt"
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
