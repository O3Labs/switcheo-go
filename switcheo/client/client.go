//Package client contains methods to make request to Switcheo API server.
package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type API struct {
	URL        string
	HTTPClient *http.Client
	UserAgent  string
}

//New initializes API with given URL. it also provides a way to overwrite *http.Client
func New(url string, httpClient *http.Client, userAgent string) *API {
	return &API{
		URL:        url,
		HTTPClient: httpClient,
		UserAgent:  userAgent,
	}
}

//Making a public request to Switcheo API server.
func (a *API) Request(method, endpoint string, params interface{}, out interface{}) error {
	url, _ := url.ParseRequestURI(a.URL)
	url.Path = url.Path + endpoint
	var req *http.Request

	if method == "GET" {
		//parse params to query string
		b, _ := json.Marshal(params)
		m := map[string]interface{}{}
		json.Unmarshal(b, &m)
		q := url.Query()
		for k, v := range m {
			switch t := v.(type) {
			case []interface{}:
				for _, value := range t {
					q.Set(fmt.Sprintf("%v", k), fmt.Sprintf("%v", value))
				}
			case string:
				q.Set(k, fmt.Sprintf("%v", v))
			default:
				q.Set(k, fmt.Sprintf("%v", v))
			}
		}
		url.RawQuery = q.Encode()
		req, _ = http.NewRequest(method, url.String(), nil)
	} else if method == "POST" {

		b, _ := json.Marshal(params)
		sortedKey := map[string]interface{}{}
		json.Unmarshal(b, &sortedKey)

		sortedB, _ := json.Marshal(sortedKey)
		log.Printf("post data %+v", string(sortedB))
		req, _ = http.NewRequest(method, url.String(), bytes.NewBuffer(sortedB))
	}
	log.Printf("%v %v", method, url.String())

	req.Header.Add("content-type", "application/json")
	req.Header.Add("UserAgent", a.UserAgent)
	res, err := a.HTTPClient.Do(req)
	defer res.Body.Close()
	if res.StatusCode != 200 {
		var e = possibleErrors[res.StatusCode]
		return errors.New(fmt.Sprintf("%v", e.Message))
	}

	if out != nil {
		err = json.NewDecoder(res.Body).Decode(&out)
	}

	return err
}
