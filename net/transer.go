package net

import "net/http"

type Transer struct {
	transporter *http.Client
	cookies     map[string]string
}

func (t *Transer) GetCookie(k string) string {
	if v, ok := t.cookies[k]; ok {
		return v
	}
	return ""
}

func (t *Transer) SetCookie(k string, v string) {
	t.cookies[k] = v
}

//get request
func (t *Transer) Get(url string) (*http.Response, error) {
}

//get request with referer
func (t *Transer) GetWithRef(url string, referer string) (*http.Response, error) {
}
