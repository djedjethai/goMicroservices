package domain

import (
	"encoding/json"
	"fmt"
	"github.com/djedjethai/bankingLib/logger"
	"net/http"
	"net/url"
)

type AuthRepository interface {
	IsAuthorized(token string, routeName string, vars map[string]string) bool
}

type RemoteAuthRepository struct{}

func NewAuthRepository() AuthRepository {
	return RemoteAuthRepository{}
}

func (f RemoteAuthRepository) IsAuthorized(token string, routeName string, vars map[string]string) bool {
	u := buildVerifyURL(token, routeName, vars)

	if response, err := http.Get(u); err != nil {
		logger.Error("Query to verify token failed" + err.Error())
		return false
	} else {
		logger.Info("response from auth authorize true.")

		m := map[string]bool{}
		if err = json.NewDecoder(response.Body).Decode(&m); err != nil {
			logger.Error("Error while decoding response from auth server:" + err.Error())
			return false
		}

		logger.Info("token has been authenticated true: " + fmt.Sprintf("%v", m))
		return m["IsAuthorized"]
	}
}

/*
  This will generate a url for token verification in the below format
  /auth/verify?token={token string}
              &routeName={current route name}
              &customer_id={customer id from the current route}
              &account_id={account id from current route if available}
  Sample: /auth/verify?token=aaaa.bbbb.cccc&routeName=MakeTransaction&customer_id=2000&account_id=95470
*/
func buildVerifyURL(token string, routeName string, vars map[string]string) string {
	u := url.URL{Host: "localhost:8181", Path: "/auth/verify", Scheme: "http"}
	q := u.Query()
	q.Add("token", token)
	q.Add("routeName", routeName)

	logger.Info("before range var: ")

	for k, v := range vars {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()

	logger.Info("build url to verify token: " + u.String())

	return u.String()
}
