package paypal

import (
	"fmt"
	"github.com/satriahrh/oddes/connection/base"
	"net/http"
	"net/url"
)

type Gopay struct {
	Connection *base.Connection
}

func NewGopay() *Gopay {
	return &Gopay{
		Connection: &base.Connection{
			SetHeaderFunc:        setHeaderFunc,
			GetAuthorizationFunc: getAuthorization,
			Client:               &http.Client{},
		},
	}
}

func getAuthorization() (authorization string, err error) {
	if true {
		authorization = "taken from cache"
	} else {
		a := url.Values{
			"grant_type":      {"client_credentials"},
			"scope":           {"payment-create"},
		}
		authenticationUrl:= "https://gw.sandbox.gopay.com/api/oauth2/token"
		usernameBasic := "username"
		passwordBasic := "password"
		req := http.Request{}
		req.SetBasicAuth(usernameBasic, passwordBasic)
		response, err := http.PostForm(authenticationUrl, a)
		if err != nil {
			return
		}

		var result struct {
			AccessToken string `json:"access_token"`
		}
		if err = base.DecodeJSONResponse(response, result); err != nil {
			return
		}
		authorization = fmt.Sprintf("Bearer %v", result.AccessToken)
	}

	return
}

func setHeaderFunc(authorizationFunc base.GetAuthorizationFunc) (http.Header, error) {
	authorization, err := authorizationFunc()
	if err != nil {
		return nil, err
	}

	return http.Header{
		"Authorization": {authorization},
	}, nil
}

func (p *Gopay) SubmitForm() error {
	response, err := p.Connection.CallRestAPI("", "POST", nil)
	if err != nil {
		return err
	}
	if err := base.DecodeJSONResponse(response, nil); err != nil {
		return err
	}
	return response, err
}