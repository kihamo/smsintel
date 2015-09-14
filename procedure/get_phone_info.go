package procedure

const (
	GetPhoneInfoPath = "getPhoneInfo"
)

type GetPhoneInfoInput struct {
	Phone string `json:"phone"`
}

type GetPhoneInfoOutput struct {
	Code    int64  `json:"code"`
	Descr   string `json:"descr"`
	Phone   string `json:"phone"`
	Country string `json:"country"`
	Opsos   string `json:"opsos"`
}
