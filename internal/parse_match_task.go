package internal

type parseMatchTask struct {
	startNum int64
	amount   int
}

func NewTask(startNum int64, amount int) *parseMatchTask {
	return &parseMatchTask{startNum: startNum, amount: amount}
}
