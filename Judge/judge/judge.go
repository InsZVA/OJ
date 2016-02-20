package judge

import (
    "io/ioutil"
)

const daemon_path = "./daemon"

type LXCer interface {
	Init(lxcTemplateName, lxcRemoteServer, lxcRemoteAlias string, id int)
    Run()
	Restore()
	Submit()
}

type JudgeHub struct {
    lxcTemplateName string      //A template lxc for copies
    lxcRemoteServer string      //Remote resources to pull when template not exist
    lxcRemoteAlias  string      //Remote image alias to pull
    lxcMaxNum       int         //Max LXC number in a JudgeHub
    lxcs            []*LXCer    //LXC List
    autoIncreament  int         //Unique id to create lxc for name
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
    status          string
    sanpshotName    string
}

func (this *Judge) Init(lxcTemplateName, lxcRemoteServer, lxcRemoteAlias string, id int) {
    if containerExist(lxcTemplateName) {
        containerCopy(lxcTemplateName, this.containerName)
    } else {
        containerLaunch(this.containerName, lxcRemoteServer, lxcRemoteAlias)
    }
    waitForLaunch(this.containerName)
    containerStartConfirmed(this.containerName)
    bytes,_ := ioutil.ReadFile(daemon_path)
    containerPush(this.containerName, "/daemon", bytes)
    this.sanpshotName = this.containerName + "snap"
    containerSnapshot(this.containerName, this.sanpshotName)
}

func (this *Judge) Run() {
    containerStartConfirmed(this.containerName)
}

func (this *Judge) Restore() {
    containerRestore(this.containerName, this.sanpshotName)
}

var TestJudge = Judge{
    containerName:      "test",
}