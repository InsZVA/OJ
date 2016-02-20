package judge

import "fmt"
import "time"

type JudgeC struct {
    Judge
}

func (this *JudgeC) Init(lxcTemplateName, lxcRemoteServer, lxcRemoteAlias string) {
    this.Judge.Init(lxcTemplateName, lxcRemoteServer, lxcRemoteAlias)
    if this.containerName != lxcTemplateName {
        return
    }
    fmt.Println("Wait 3000 ms")
    time.Sleep(time.Millisecond * 3000)
    fmt.Println("Innstalling gcc on ", this.containerName)
    _, uuid := containerExec(this.containerName, []string{"apt-get", "update"})
    wait(uuid)
    _, uuid = containerExec(this.containerName, []string{"apt-get", "-y", "--force-yes", "install", "gcc"})
    wait(uuid)
    _, ips, _ := containerState(this.containerName)
    fmt.Println("Starting daemon on ", this.containerName, " ip: ", ips[0])
    _, uuid = containerExec(this.containerName, []string{"chmod", "777", "/daemon"})
    wait(uuid)
    _, uuid = containerExec(this.containerName, []string{"/daemon"})
    fmt.Println("Wait 1000 ms", this.containerName)
    time.Sleep(time.Millisecond * 1000)
    fmt.Println("Creating snapshot on ", this.containerName)
    this.sanpshotName = this.containerName + "snap"
    _, uuid = containerSnapshot(this.containerName, this.sanpshotName)
    wait(uuid)
}