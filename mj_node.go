package goquery

import (
	"bytes"
	"code.google.com/p/mahonia"
	"exp/html"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func GbkToUtf8(s string) string {
	decoder := mahonia.NewDecoder("gb18030")
	return strings.TrimSpace(decoder.ConvertString(s))
}

func clear_space(s string) string {
	return strings.TrimSpace(s)
}

func NewDocumentFromPostUrl(uri string, data url.Values) (d *Document, err error) {
	resp, err := http.PostForm(uri, data)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	root, err := html.Parse(resp.Body)
	if err != nil {
		return
	}

	d = newDocument(root, resp.Request.URL)
	return
}

func NewDocumentGBK(uri string) (d *Document, err error) {
	resp, err := http.Get(uri)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	body_gbk := bytes.NewBufferString(GbkToUtf8(string(body)))
	resp.Body = ioutil.NopCloser(body_gbk)
	root, err := html.Parse(resp.Body)
	if err != nil {
		return
	}
	d = newDocument(root, resp.Request.URL)
	return
}
