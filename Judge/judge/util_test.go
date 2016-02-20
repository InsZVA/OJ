package judge

import "testing"
import "fmt"
/*
func Test_containerLaunch(t *testing.T) {
    err, uuid := containerLaunch("ubuntu2", "https://images.linuxcontainers.org:8443", "ubuntu/trusty/amd64")
    if err != nil {
        t.Error(err)
    }
    wait(uuid)
    fmt.Println("Launch ubuntu2")
}

func Test_containerCopy(t *testing.T) {
    err, uuid := containerCopy("ubuntu2", "ubuntu3")
    if err != nil {
        t.Error(err)
    }
    wait(uuid)
    fmt.Println("Launch ubuntu3")
}

func Test_containerStart(t *testing.T) {
    err, uuid := containerStart("ubuntu2")
    if err != nil {
        t.Error(err)
    }
    wait(uuid)
    fmt.Println("Start ubuntu2")
}

func Test_containerState(t *testing.T) {
    state, ips, err := containerState("ubuntu2")
    if state != "Running" || err != nil {
        t.Error(state, ips, err)
    }
    fmt.Println("Get ubuntu2 state")
}

func Test_containerPush(t *testing.T) {
    err := containerPush("ubuntu2", "/1.txt", []byte("123"))
    if err != nil {
        t.Error(err)
    }
    fmt.Println("Push a file on ubuntu2")
}

func Test_containerSnapshot(t *testing.T) {
    err, uuid := containerSnapshot("ubuntu2", "abcsnap")
    if err != nil {
        t.Error(err)
    }
    wait(uuid)
    fmt.Println("Created a snapshot on ubuntu2")
}


func Test_containerRestore(t *testing.T) {
     err, uuid := containerRestore("ubuntu2", "abcsnap")
     if err != nil {
         t.Error(err)
     }
     wait(uuid)
     fmt.Println("Restore that snapshot on ubuntu2")
}
*/
func Test_containerExec(t *testing.T) {
     _, uuid := containerStart("ubuntu2")
     wait(uuid)
     fmt.Println("Start ubuntu2")
     err, uuid := containerExec("ubuntu2", []string{"apt-get", "update"})
     if err != nil {
         t.Error(err)
     }
     wait(uuid)
     fmt.Println("Execute a command on ubuntu2")
}
