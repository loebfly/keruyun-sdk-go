package network

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/loebfly/keruyun-sdk-go/internal/config"
	"github.com/loebfly/keruyun-sdk-go/internal/util"
	"github.com/loebfly/keruyun-sdk-go/kry_model"
)

type HttpMethod string

const (
	GET  HttpMethod = "GET"
	POST HttpMethod = "POST"
)

func JsonToResult[D any](options JsonOptions) kry_model.Result[D] {
	return toResult[D](JsonRequest(options))
}

func JsonToAny(options JsonOptions) (any, error) {
	res, err := JsonRequest(options)
	return string(res), err
}

// JsonRequest Json请求
func JsonRequest(options JsonOptions) ([]byte, error) {
	// 构建http请求
	req, err := http.NewRequest(string(options.Method), options.Url(), options.BodyBuffer())
	if err != nil {
		return nil, err
	}
	// 设置请求头
	for k, v := range options.Header {
		req.Header.Set(k, v)
	}
	// 设置content-type
	req.Header.Set("Content-Type", "application/json")

	// 设置请求参数
	req.URL.RawQuery = options.GetSignQueryStr()

	// 设置超时时间
	if options.timeout == 0 {
		options.timeout = 30 * time.Second
	}

	client := http.DefaultClient
	client.Timeout = options.timeout

	// 开始时间
	startTime := time.Now()
	reqTime := startTime.Format("2006-01-02 15:04:05.012")

	var reqQuery = make(map[string]string)
	for k, v := range req.URL.Query() {
		reqQuery[k] = v[0]
	}

	// 发送请求
	resp, doErr := client.Do(req)

	// 结束时间
	endTime := time.Now()
	respTime := endTime.Format("2006-01-02 15:04:05.012")

	respBuffer := new(bytes.Buffer)

	// 外部打印请求日志
	defer func() {
		if config.Global.PrintApiLogHandle == nil {
			return
		}

		var respBody = make(map[string]any)
		if respBuffer != nil {
			unmarshalErr := json.Unmarshal(respBuffer.Bytes(), &respBody)
			if unmarshalErr != nil {
				util.Any(kry_model.Result[any]{
					Code:    -999,
					Message: unmarshalErr.Error(),
				}).ToObject(&respBody)
			}
		}

		if doErr != nil {
			util.Any(kry_model.Result[any]{
				Code:    -999,
				Message: err.Error(),
			}).ToObject(&respBody)
		}
		if resp.StatusCode != http.StatusOK {
			util.Any(kry_model.Result[any]{
				Code:    -999,
				Message: fmt.Sprintf("http status code: %d", resp.StatusCode),
			}).ToObject(&respBody)
		}

		// 请求日志
		reqCtx := kry_model.ReqCtx{
			ReqTime:     reqTime,
			RespTime:    respTime,
			TTL:         int(endTime.Sub(startTime).Milliseconds()),
			Method:      string(options.Method),
			ContentType: "application/json",
			Host:        options.Host,
			URI:         options.Uri,
			ReqQuery:    reqQuery,
			ReqBody:     options.JSON,
			RespBody:    respBody,
		}
		config.Global.PrintApiLogHandle(reqCtx)
	}()

	if err != nil {
		return nil, err
	}

	// 读取响应
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status code: %d", resp.StatusCode)
	}

	_, err = respBuffer.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}
	// 返回响应
	return respBuffer.Bytes(), nil
}

// toResult 将字符串转换为Result
func toResult[D any](resp []byte, err error) kry_model.Result[D] {
	if err != nil {
		return kry_model.Result[D]{
			Code:    -999,
			Message: err.Error(),
		}
	}
	if resp == nil {
		return kry_model.Result[D]{
			Code:    -999,
			Message: "response is nil",
		}
	}
	var result kry_model.Result[D]
	err = json.Unmarshal(resp, &result)
	fmt.Println(err, "---")
	if err != nil {
		return kry_model.Result[D]{
			Code:    -999,
			Message: "response unmarshal failure",
		}
	} else {
		return result
	}
}
