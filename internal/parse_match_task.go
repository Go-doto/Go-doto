package internal

type ParseMatchTask struct {
	StartNum     int64 `json:"startNum"`
	Amount       int   `json:"amount"`
	QueriesLimit int   `json:"queriesLimit"`
}
