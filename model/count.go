package model

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func SetSessionValue(r *http.Request, w http.ResponseWriter, key string, value interface{}) error {
	session, err := store.Get(r, "session-name")
	if err != nil {
		return err
	}
	session.Values[key] = value
	return session.Save(r, w)
}

func GetSessionValue(r *http.Request, key string) (interface{}, error) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		return nil, err
	}
	return session.Values[key], nil
}
