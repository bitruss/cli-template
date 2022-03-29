package http

import (
	"errors"
	"reflect"
	"time"

	"github.com/imroc/req"
)

type RespBody struct {
	Status int         `json:"status" `
	Result interface{} `json:"result" `
	Msg    string      `json:"msg" `
}

const post = "POST"
const get = "GET"
const API_TIMEOUT_SECS = 10

type ApiRequest struct {
	Err            error
	HttpStatusCode int
	Body           interface{}
}

func (err *ApiRequest) Ok() bool {
	return err.Err == nil && err.HttpStatusCode != 200
}

func (err *ApiRequest) IsHttpError() bool {
	return err.HttpStatusCode != 200
}

func Get(url string, token string, apiq *ApiRequest) {
	request(get, url, token, nil, API_TIMEOUT_SECS, apiq)
}

func Get_(url string, token string, timeOutSec int, apiq *ApiRequest) {
	request(get, url, token, nil, timeOutSec, apiq)
}

func POST(url string, token string, postData interface{}, apiq *ApiRequest) {
	request(post, url, token, postData, API_TIMEOUT_SECS, apiq)
}

func POST_(url string, token string, postData interface{}, timeOutSec int, apiq *ApiRequest) {
	request(post, url, token, postData, timeOutSec, apiq)
}

func request(method string, url string, token string, postData interface{}, timeOutSec int, apiq *ApiRequest) {
	if apiq.Body != nil {
		t := reflect.TypeOf(apiq.Body).Kind()
		if t != reflect.Ptr && t != reflect.Slice && t != reflect.Map {
			apiq.Err = errors.New("value only support Pointer Slice and Map")
			apiq.HttpStatusCode = 200
			return
		}
	}

	authHeader := req.Header{
		"Accept": "application/json",
	}

	if token != "" {
		authHeader["Authorization"] = "Bearer " + token
	}

	r := req.New()
	r.SetTimeout(time.Duration(timeOutSec) * time.Second)

	var resp *req.Resp
	var err error

	switch method {
	case get:
		resp, err = r.Get(url, authHeader)
	case post:
		resp, err = r.Post(url, authHeader, req.BodyJSON(postData))
	default:
		// imposssible
	}

	if err != nil {
		apiq.Err = err
		apiq.HttpStatusCode = 200
		return
	}

	if resp.Response().StatusCode != 200 {
		apiq.Err = errors.New("network error")
		apiq.HttpStatusCode = resp.Response().StatusCode
		return
	}

	respData := &RespBody{
		Result: apiq.Body,
	}
	err = resp.ToJSON(&respData)
	if err != nil {
		apiq.Err = err
		apiq.HttpStatusCode = 200
		return
	}

	if respData.Status <= 0 {
		apiq.Err = errors.New(respData.Msg)
		apiq.HttpStatusCode = 200
		return
	}

	apiq.Err = nil
	apiq.HttpStatusCode = 200
}
