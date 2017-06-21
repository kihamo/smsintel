package smsintel

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/kihamo/gotypes"
)

type Request struct {
	request       *http.Request
	requestParams *url.Values

	output interface{}
}

func NewRequest(apiUrl, procedure, login, password string, input interface{}, output interface{}) *Request {
	httpRequest, _ := http.NewRequest("GET", fmt.Sprintf("%s/%s.php", apiUrl, procedure), nil)

	r := &Request{
		request:       httpRequest,
		requestParams: &url.Values{},
		output:        output,
	}

	reflectInput := reflect.Indirect(reflect.ValueOf(input))
	if reflectInput.Kind() == reflect.Struct {
		for i := 0; i < reflectInput.NumField(); i++ {
			field := reflectInput.Type().Field(i)

			tag := field.Tag.Get("json")
			if tag == "" {
				continue
			}

			key := strings.Split(tag, ",")[0]
			value := reflectInput.FieldByName(field.Name)

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

	r.requestParams.Add("login", login)
	r.requestParams.Add("password", password)

	var buf bytes.Buffer

	if r.request.URL.RawQuery != "" {
		buf.WriteString(r.request.URL.RawQuery)
		buf.WriteByte('&')
	}

	buf.WriteString(r.requestParams.Encode())
	r.request.URL.RawQuery = buf.String()

	return r
}

func (r *Request) Send() error {
	return r.SendWithContext(context.Background())
}

func (r *Request) SendWithContext(ctx context.Context) error {
	req := r.request.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
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

		if code != 1 {
			message := gotypes.ToString(response["descr"])
			return NewApiError(code, message, nil)
		}
	}

	converter := gotypes.NewConverter(response, r.output)
	if !converter.Valid() {
		return errors.New(strings.Join(converter.GetInvalidFields(), ","))
	}

	return nil
}
