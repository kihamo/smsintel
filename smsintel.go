package smsintel

//go:generate goimports -w ./

const (
	smsIntelApiUrl = "https://lcab.smsintel.ru/lcabApi"
)

type SmsIntel struct {
	apiUrl   string
	login    string
	password string
}

func NewSmsIntel(login string, password string) *SmsIntel {
	return &SmsIntel{
		apiUrl:   smsIntelApiUrl,
		login:    login,
		password: password,
	}
}

func (s *SmsIntel) SetOptions(options map[string]string) error {
	if v, ok := options["api_url"]; ok {
		s.apiUrl = v
	}

	if v, ok := options["login"]; ok {
		s.login = v
	}

	if v, ok := options["password"]; ok {
		s.password = v
	}

	return nil
}

func (s *SmsIntel) newRequest(procedure string, input interface{}, output interface{}) *Request {
	return NewRequest(s.apiUrl, procedure, s.login, s.password, input, output)
}
