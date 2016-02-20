package main
import (
    "./judge"
    "flag"
    _"net/http"
)

var compilerCNum = flag.Int("c", 5, "C Compilers Num")

func main() {
    flag.Parse()
    judge.NewJudgeHubC(*compilerCNum)
    //TODO: Add HTTP Listen
}