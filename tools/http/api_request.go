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

type ApiResult struct {
	Err            error
	HttpStatusCode int
}

func (err *ApiResult) Ok() bool {
	return err.Err == nil && err.HttpStatusCode != 200
}

func (err *ApiResult) IsHttpError() bool {
	return err.HttpStatusCode != 200
}

func Get(url string, token string, respResult interface{}) ApiResult {
	return request(get, url, token, nil, API_TIMEOUT_SECS, respResult)
}

func Get_(url string, token string, timeOutSec int, respResult interface{}) ApiResult {
	return request(get, url, token, nil, timeOutSec, respResult)
}

func POST(url string, token string, postData interface{}, respResult interface{}) ApiResult {
	return request(post, url, token, postData, API_TIMEOUT_SECS, respResult)
}

func POST_(url string, token string, postData interface{}, timeOutSec int, respResult interface{}) ApiResult {
	return request(post, url, token, postData, timeOutSec, respResult)
}

func request(method string, url string, token string, postData interface{}, timeOutSec int, respResult interface{}) ApiResult {
	if respResult != nil {
		t := reflect.TypeOf(respResult).Kind()
		if t != reflect.Ptr && t != reflect.Slice && t != reflect.Map {
			return ApiResult{
				Err:            errors.New("value only support Pointer Slice and Map"),
				HttpStatusCode: 200,
			}
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
		return ApiResult{
			Err:            err,
			HttpStatusCode: 200,
		}
	}

	if resp.Response().StatusCode != 200 {
		return ApiResult{
			Err:            errors.New("network error"),
			HttpStatusCode: resp.Response().StatusCode,
		}
	}

	respData := &RespBody{
		Result: respResult,
	}
	err = resp.ToJSON(&respData)
	if err != nil {
		return ApiResult{
			Err:            err,
			HttpStatusCode: 200,
		}
	}
	if respData.Status <= 0 {
		return ApiResult{
			Err:            errors.New(respData.Msg),
			HttpStatusCode: 200,
		}
	}

	return ApiResult{
		Err:            nil,
		HttpStatusCode: 200,
	}
}
