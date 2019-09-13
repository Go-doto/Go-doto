package internal

type parseMatchTask struct {
	StartNum int64 `json:"startNum"`
	Amount   int   `json:"amount"`
}

func NewTask(startNum int64, amount int) *parseMatchTask {
	return &parseMatchTask{StartNum: startNum, Amount: amount}
}
