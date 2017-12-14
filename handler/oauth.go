package handler

import (
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type OAuth struct {
	l    *log.Logger
	conf *oauth2.Config
}

func NewOAuth(id, code, redirect string) *OAuth {
	return &OAuth{
		l: log.New(ioutil.Discard, "", 0),
		conf: &oauth2.Config{
			ClientID:     id,
			ClientSecret: code,
			Scopes:       []string{"openid", "profile", "email"},
			Endpoint:     google.Endpoint,
			RedirectURL:  redirect,
		},
	}
}

func (oa *OAuth) SetLogger(l *log.Logger) {
	oa.l = l
}

func (oa *OAuth) GetURL() string {
	return oa.conf.AuthCodeURL("state")
}

func (oa *OAuth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		oa.l.Println("fail to parse form:", err)
		return
	}
	code := r.Form["code"]
	if len(code) != 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("missing code parameter"))
		return
	}
	w.Write([]byte("ok"))
}
