package procedure

const (
	RequestSourcePath = "requestSource"
)

type RequestSourceInput struct {
	Source string `json:"source"`
}

type RequestSourceOutput struct {
	Code  int64  `json:"code"`
	Descr string `json:"descr"`
}
