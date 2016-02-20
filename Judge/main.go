package main
import (
    "./judge"
)

func main() {
    judge.TestJudge.Init("ubuntu2", "htts://images.linuxcontainers.org:8443", "ubuntu/trusty/amd64", 1)
}