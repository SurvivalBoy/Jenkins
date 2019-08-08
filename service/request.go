package service

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

// HTTPPostRequest 根据地址和参数，使用post请求返回Body
func HTTPPostRequest(uri string, data url.Values) (body Body, err error) {
	if len(uri) == 0 {
		return body, errors.New("uri is null")
	}

	c := GetHTTPClient()
	resp, err := c.PostForm(uri, data)
	if err != nil {
		return body, err
	}
	if resp != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode != 200 {
		return body, errors.New(resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(&body)
	return body, err
}

// HTTPGetRequest 根据get请求返回结果
func HTTPGetRequest(uri string) (body Body, err error) {
	if len(uri) == 0 {
		return body, errors.New("uri is null")
	}

	resp, err := http.Get(uri)
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return body, err
	}

	if resp.StatusCode != 200 {
		return body, errors.New(resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(&body)
	return body, err
}

// 获取即开即关链接
func GetHTTPClient() http.Client {
	tr := http.Transport{DisableKeepAlives: true}
	return http.Client{Transport: &tr}
}

// Body 返回结果结构体
type Body struct {
	Status    string      `json:"status,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Error     string      `json:"error,omitempty"`
	ErrorMsg  string      `json:"errorMsg,omitempty"`
	NewErrMsg string      `json:"newErrMsg,omitempty"`
}
