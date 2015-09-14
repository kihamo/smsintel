package smsintel

//go:generate goimports -w ./

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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
	request   *http.Request
	response  *http.Response

	Input  interface{}
	Output interface{}
}

func NewRequest(procedure string, login string, password string, input interface{}, output interface{}) *Request {
	httpRequest, _ := http.NewRequest("GET", smsIntelApiUrl+procedure+".php", nil)

	return &Request{
		procedure: procedure,
		login:     login,
		password:  password,
		request:   httpRequest,
		Input:     input,
		Output:    output,
	}
}

func (r *Request) getParamsFromInput() *url.Values {
	params := &url.Values{}

	return params
}

func (r *Request) Send() (err error) {
	params := r.getParamsFromInput()
	params.Add("login", r.login)
	params.Add("password", r.password)

	r.request.URL, _ = url.Parse(r.request.URL.String() + "?" + params.Encode())
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

	fmt.Println(r.request.URL.String())
	fmt.Println(string(content))

	converter := gotypes.NewConverter(response, r.Output)
	if !converter.Valid() {
		err = errors.New(strings.Join(converter.GetInvalidFields(), ","))
	}

	return err
}
