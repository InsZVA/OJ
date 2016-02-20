package judge

import "testing"
import "fmt"

func Test_containerLaunch(t *testing.T) {
    err := containerLaunch("ubuntu2", "https://images.linuxcontainers.org:8443", "ubuntu/trusty/amd64")
    if err != nil {
        t.Error(err)
    }
    fmt.Println("Launch ubuntu2")
}

func Test_containerCopy(t *testing.T) {
    err := containerCopy("ubuntu2", "ubuntu3")
    if err != nil {
        t.Error(err)
    }
    fmt.Println("Launch ubuntu3")
}
/*
func Test_containerStart(t *testing.T) {
    waitForLaunch("ubuntu2")
    containerStartConfirmed("ubuntu2")
    fmt.Println("Start ubuntu2")
}

func Test_containerState(t *testing.T) {
    state, ips, err := containerState("ubuntu2")
    if state != "Running" || err != nil {
        t.Error(state, ips, err)
    }
}

func Test_containerPush(t *testing.T) {
    err := containerPush("ubuntu2", "/1.txt", []byte("123"))
    if err != nil {
        t.Error(err)
    }
}

func Test_containerSnapshot(t *testing.T) {
    err := containerSnapshot("ubuntu2", "abcsnap")
    if err != nil {
        t.Error(err)
    }
}

func Test_containerSnapshotReady(t *testing.T) {
    t.Error(containerSnapshotReady("ubuntu2", "abcsnap"))
}

func Test_containerRestore(t *testing.T) {
     err := containerRestore("ubuntu2", "abcsnap")
     if err != nil {
         t.Error(err)
     }
}*/