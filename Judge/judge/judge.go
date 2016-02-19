package judge

type LXCer interface {
	Init()
	Restore()
	Submit()
}

type Judge struct {
    judgeId int
    problemId int
}
