package procedure

const (
	SendSmsPath = "sendSms"
)

type SendSmsInput struct {
	Txt           *string `json:"txt"`
	To            *string `json:"to"`
	IdGroup       *string `json:"idGroup"`
	Source        *string `json:"source"`
	Flash         *bool   `json:"flash"`
	DateTimeSend  *string `json:"dateTimeSend"`
	OnlyDelivery  *bool   `json:"onlydelivery"`
	UseAlfaSource *bool   `json:"use_alfasource"`
	DiscountId    *bool   `json:"discountID"`
}

type SendSmsOutput struct {
	Code            int64   `json:"code"`
	Descr           string  `json:"descr"`
	ColsmsOfSending int64   `json:"colsmsOfSending"`
	PriceOfSending  float64 `json:"priceOfSending"`
}
