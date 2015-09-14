package smsintel

import (
	"github.com/kihamo/smsintel/procedure"
)

func (s *SmsIntel) SendSms(input *procedure.SendSmsInput) (*procedure.SendSmsOutput, error) {
	request := s.SendSmsRequest(input)
	return request.Output.(*procedure.SendSmsOutput), request.Send()
}

func (s *SmsIntel) SendSmsRequest(input *procedure.SendSmsInput) *Request {
	if input == nil {
		input = &procedure.SendSmsInput{}
	}

	output := &procedure.SendSmsOutput{}
	request := s.newRequest(procedure.SendSmsPath, input, output)

	return request
}

func (s *SmsIntel) Info(input *procedure.InfoInput) (*procedure.InfoOutput, error) {
	request := s.InfoRequest(input)
	return request.Output.(*procedure.InfoOutput), request.Send()
}

func (s *SmsIntel) InfoRequest(input *procedure.InfoInput) *Request {
	if input == nil {
		input = &procedure.InfoInput{}
	}

	output := &procedure.InfoOutput{}
	request := s.newRequest(procedure.IntoPath, input, output)

	return request
}

func (s *SmsIntel) RequestSource(input *procedure.RequestSourceInput) (*procedure.RequestSourceOutput, error) {
	request := s.RequestSourceRequest(input)
	return request.Output.(*procedure.RequestSourceOutput), request.Send()
}

func (s *SmsIntel) RequestSourceRequest(input *procedure.RequestSourceInput) *Request {
	if input == nil {
		input = &procedure.RequestSourceInput{}
	}

	output := &procedure.RequestSourceOutput{}
	request := s.newRequest(procedure.RequestSourcePath, input, output)

	return request
}

func (s *SmsIntel) GetPhoneInfo(input *procedure.GetPhoneInfoInput) (*procedure.GetPhoneInfoOutput, error) {
	request := s.GetPhoneInfoRequest(input)
	return request.Output.(*procedure.GetPhoneInfoOutput), request.Send()
}

func (s *SmsIntel) GetPhoneInfoRequest(input *procedure.GetPhoneInfoInput) *Request {
	if input == nil {
		input = &procedure.GetPhoneInfoInput{}
	}

	output := &procedure.GetPhoneInfoOutput{}
	request := s.newRequest(procedure.GetPhoneInfoPath, input, output)

	return request
}
