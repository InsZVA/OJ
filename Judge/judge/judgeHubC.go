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
    if !containerExist(template.containerName)  {
        template.Init(judgeHub.lxcTemplateName, judgeHub.lxcRemoteServer, judgeHub.lxcRemoteAlias)
    }
    
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