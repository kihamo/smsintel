package smsintel

type SmsIntel struct {
	login    string
	password string
}

func NewSmsIntel(login string, password string) *SmsIntel {
	return &SmsIntel{
		login:    login,
		password: password,
	}
}

func (s *SmsIntel) newRequest(procedure string, input interface{}, output interface{}) *Request {
	return NewRequest(procedure, s.login, s.password, input, output)
}
