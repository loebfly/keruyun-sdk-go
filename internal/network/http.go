package network

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/loebfly/keruyun-sdk-go/kry_result"
	"io"
	"net/http"
	"time"
)

type HttpMethod string

const (
	GET    HttpMethod = "GET"
	POST   HttpMethod = "POST"
	PUT    HttpMethod = "PUT"
	DELETE HttpMethod = "DELETE"
)

func JsonToResult[D any](options JsonOptions) kry_result.Result[D] {
	return toResult[D](JsonRequest(options))
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
	// 设置请求参数
	req.URL.RawQuery = options.GetSignQueryStr()

	// 设置超时时间
	if options.timeout == 0 {
		options.timeout = 30 * time.Second
	}

	client := http.DefaultClient
	client.Timeout = options.timeout

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	// 关闭请求
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	// 读取响应
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http status code: %d", resp.StatusCode)
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}
	// 返回响应
	return buf.Bytes(), nil
}

// toResult 将字符串转换为Result
func toResult[D any](resp []byte, err error) kry_result.Result[D] {
	if err != nil {
		return kry_result.Result[D]{
			Code:    -999,
			Message: err.Error(),
		}
	}
	if resp == nil {
		return kry_result.Result[D]{
			Code:    -999,
			Message: "response is nil",
		}
	}
	var result kry_result.Result[D]
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return kry_result.Result[D]{
			Code:    -999,
			Message: "response unmarshal failure",
		}
	} else {
		return result
	}
}
