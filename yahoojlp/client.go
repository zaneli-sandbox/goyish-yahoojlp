package yahoojlp

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
)

type httpMethod string

const (
	apiUrl            = "http://jlp.yahooapis.jp/%s/%s/%s"
	get    httpMethod = "GET"
	post   httpMethod = "POST"
)

type Client struct {
	appid string
}

func NewClient(appid string) *Client {
	return &Client{appid}
}

func (c *Client) callApi(
	httpMethod httpMethod, serviceName string, apiVersion string, apiName string, params map[string]string, result interface{}) error {
	if result == nil {
		return fmt.Errorf("Invalid arguments: result type is nil")
	}

	endPoint := fmt.Sprintf(apiUrl, serviceName, apiVersion, apiName)
	params["appid"] = c.appid

	var res *http.Response
	var err error
	switch httpMethod {
	case get:
		url, err := createUrlWithQuery(endPoint, params)
		if err != nil {
			return err
		}
		res, err = http.Get(url)
		if err != nil {
			return err
		}
	case post:
		res, err = http.PostForm(endPoint, createValues(params))
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("Invalid http method: %s", httpMethod)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Invalid status: %s", res.Status)
	}
	if err = xml.NewDecoder(res.Body).Decode(result); err != nil {
		return err
	}
	return nil
}

func (c *Client) String() string {
	return "yahoojlp.Client"
}

func createUrlWithQuery(endPoint string, params map[string]string) (string, error) {
	u, err := url.Parse(endPoint)
	if err != nil {
		return "", err
	}
	q := u.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()
	return fmt.Sprint(u), nil
}

func createValues(params map[string]string) url.Values {
	values := map[string][]string{}
	for k, v := range params {
		values[k] = []string{v}
	}
	return values
}
