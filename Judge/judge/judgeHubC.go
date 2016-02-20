package judge

func NewJudgeHubC(lxcMaxNum int) *JudgeHub {
    return &JudgeHub{
        lxcTemplateName:    "JudgeCTemplate0",
        lxcRemoteServer:    "https://images.linuxcontainers.org:8443",
        lxcRemoteAlias:     "ubuntu/trusty/adm64",
        autoIncreament:      1,
    }
}