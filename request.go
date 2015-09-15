package smsintel

//go:generate goimports -w ./

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/kihamo/gotypes"
)

const (
	smsIntelApiUrl = "https://lcab.smsintel.ru/lcabApi/"
)

var client *http.Client

func init() {
	client = &http.Client{
		Timeout: time.Second * 5,
	}
}

type Request struct {
	procedure string
	login     string
	password  string
	response  *http.Response

	request       *http.Request
	requestParams *url.Values

	Input  interface{}
	Output interface{}
}

func NewRequest(procedure string, login string, password string, input interface{}, output interface{}) *Request {
	httpRequest, _ := http.NewRequest("GET", smsIntelApiUrl+procedure+".php", nil)

	return &Request{
		procedure:     procedure,
		login:         login,
		password:      password,
		request:       httpRequest,
		requestParams: &url.Values{},
		Input:         input,
		Output:        output,
	}
}

func (r *Request) setParamsFromInput() {
	input := reflect.Indirect(reflect.ValueOf(r.Input))
	if input.Kind() == reflect.Struct {
		for i := 0; i < input.NumField(); i++ {
			field := input.Type().Field(i)

			tag := field.Tag.Get("json")
			if tag == "" {
				continue
			}

			key := strings.Split(tag, ",")[0]
			value := input.FieldByName(field.Name)

			if value.Kind() == reflect.Ptr {
				value = reflect.Indirect(value)

				if !value.IsValid() {
					continue
				}
			}

			if value.Kind() == reflect.Bool {
				r.requestParams.Add(key, strconv.Itoa(gotypes.ToInt(value.Interface())))
			} else {
				r.requestParams.Add(key, gotypes.ToString(value.Interface()))
			}
		}
	}
}

func (r *Request) sign() {
	r.requestParams.Add("login", r.login)
	r.requestParams.Add("password", r.password)
}

func (r *Request) Send() (err error) {
	r.setParamsFromInput()
	r.sign()

	var buf bytes.Buffer

	if r.request.URL.RawQuery != "" {
		buf.WriteString(r.request.URL.RawQuery)
		buf.WriteByte('&')
	}

	buf.WriteString(r.requestParams.Encode())
	r.request.URL.RawQuery = buf.String()

	if r.response, err = client.Do(r.request); err != nil {
		return err
	}

	defer r.response.Body.Close()

	content, err := ioutil.ReadAll(r.response.Body)
	if err != nil {
		return err
	}

	var response interface{}
	err = json.Unmarshal(content, &response)
	if err != nil {
		return err
	}

	if response, ok := response.(map[string]interface{}); ok {
		code := gotypes.ToInt64(response["code"])
		message := gotypes.ToString(response["descr"])

		return NewApiError(code, message, nil)
	}

	converter := gotypes.NewConverter(response, r.Output)
	if !converter.Valid() {
		return errors.New(strings.Join(converter.GetInvalidFields(), ","))
	}

	return nil
}
