package smsintel

import (
	"context"

	"github.com/kihamo/smsintel/procedure"
)

func (s *SmsIntel) GetPhoneInfo(input *procedure.GetPhoneInfoInput) (*procedure.GetPhoneInfoOutput, error) {
	request := s.GetPhoneInfoRequest(input)
	return request.output.(*procedure.GetPhoneInfoOutput), request.Send()
}

func (s *SmsIntel) GetPhoneInfoWithContext(ctx context.Context, input *procedure.GetPhoneInfoInput) (*procedure.GetPhoneInfoOutput, error) {
	request := s.GetPhoneInfoRequest(input)
	return request.output.(*procedure.GetPhoneInfoOutput), request.SendWithContext(ctx)
}

func (s *SmsIntel) GetPhoneInfoRequest(input *procedure.GetPhoneInfoInput) *Request {
	if input == nil {
		input = &procedure.GetPhoneInfoInput{}
	}

	output := &procedure.GetPhoneInfoOutput{}
	request := s.newRequest(procedure.GetPhoneInfoPath, input, output)

	return request
}

func (s *SmsIntel) Info(input *procedure.InfoInput) (*procedure.InfoOutput, error) {
	request := s.InfoRequest(input)
	return request.output.(*procedure.InfoOutput), request.Send()
}

func (s *SmsIntel) InfoWithContext(ctx context.Context, input *procedure.InfoInput) (*procedure.InfoOutput, error) {
	request := s.InfoRequest(input)
	return request.output.(*procedure.InfoOutput), request.SendWithContext(ctx)
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
	return request.output.(*procedure.RequestSourceOutput), request.Send()
}

func (s *SmsIntel) RequestSourceWithContext(ctx context.Context, input *procedure.RequestSourceInput) (*procedure.RequestSourceOutput, error) {
	request := s.RequestSourceRequest(input)
	return request.output.(*procedure.RequestSourceOutput), request.SendWithContext(ctx)
}

func (s *SmsIntel) RequestSourceRequest(input *procedure.RequestSourceInput) *Request {
	if input == nil {
		input = &procedure.RequestSourceInput{}
	}

	output := &procedure.RequestSourceOutput{}
	request := s.newRequest(procedure.RequestSourcePath, input, output)

	return request
}

func (s *SmsIntel) SendSms(input *procedure.SendSmsInput) (*procedure.SendSmsOutput, error) {
	request := s.SendSmsRequest(input)
	return request.output.(*procedure.SendSmsOutput), request.Send()
}

func (s *SmsIntel) SendSmsWithContext(ctx context.Context, input *procedure.SendSmsInput) (*procedure.SendSmsOutput, error) {
	request := s.SendSmsRequest(input)
	return request.output.(*procedure.SendSmsOutput), request.SendWithContext(ctx)
}

func (s *SmsIntel) SendSmsRequest(input *procedure.SendSmsInput) *Request {
	if input == nil {
		input = &procedure.SendSmsInput{}
	}

	output := &procedure.SendSmsOutput{}
	request := s.newRequest(procedure.SendSmsPath, input, output)

	return request
}
