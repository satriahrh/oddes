package base

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type GetAuthorizationFunc func() (string, error)
type SetHeaderFunc func(authorizationFunc GetAuthorizationFunc) (http.Header, error)

type Connection struct {
	GetAuthorizationFunc GetAuthorizationFunc
	SetHeaderFunc        SetHeaderFunc
	Client               *http.Client
}

func (c *Connection) CallRestAPI(url, method string, body *bytes.Buffer) (response *http.Response, err error) {
	header, err := c.SetHeaderFunc(c.GetAuthorizationFunc)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return
	}
	request.Header = header

	response, err = c.Client.Do(request)
	if err != nil {
		return
	}

	return
}

func DecodeJSONResponse(response *http.Response, result interface{}) error {
	return json.NewDecoder(response.Body).Decode(&result)
}
