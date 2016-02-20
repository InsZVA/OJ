package judge

import (
    "strconv"
)

func NewJudgeHubC(lxcMaxNum int) *JudgeHub {
    judgeHub := NewJudgeHub("JudgeC", "https://images.linuxcontainers.org:8443", "ubuntu/trusty/amd64", lxcMaxNum)
    judgeHub.compiler = "gcc"
    template := JudgeC{
        Judge{
            containerName:      "JudgeC",
        },
    }
    template.Init(judgeHub.lxcTemplateName, judgeHub.lxcRemoteServer, judgeHub.lxcRemoteAlias)
    for i := 0;i < lxcMaxNum;i++ {
        newJudgeC := &JudgeC{
            Judge{
                containerName:      "JudgeC" + strconv.Itoa(judgeHub.autoIncreament),
            },
        }
        judgeHub.autoIncreament++
        newJudgeC.Init(judgeHub.lxcTemplateName, judgeHub.lxcRemoteServer, judgeHub.lxcRemoteAlias)
        judgeHub.lxcs = append(judgeHub.lxcs, newJudgeC)
    }
    return judgeHub
}

func (this *JudgeHub) Check(problemId string, compiler string, api string, submission string) []byte {
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
    //TODO: notify the judge result
    return nil
}