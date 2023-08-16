package network

import (
	"bytes"
	"encoding/json"
	"github.com/loebfly/keruyun-sdk-go/internal/util"
	"time"
)

const (
	/*
		公共参数(除去sign)按照参数名ASCII码进行自然排序,然后按照k1v1k2v2…拼接（appKey+shopIdenty/brandId+timestamp+version）,结果为parameters;
		拼接parameters+body(请求业务参数)+token;
		将所得字符串进行SHA-256运算,返回即为sign的值。
	*/
	SingPtrSign      = "sign"
	SignPtrAppKey    = "appKey"
	SignPtrSecretKey = "secretKey"
	SingPtrShopId    = "shopIdenty"
	SignPtrBrandId   = "brandId"
	SignPtrTimestamp = "timestamp"
	SignPtrVersion   = "version"
	SignPtrBody      = "body"
	SignPtrToken     = "token"
)

type JsonOptions struct {
	Host    string            // 请求主机
	Uri     string            // 请求路径
	Method  HttpMethod        // 请求方法
	timeout time.Duration     // 超时时间, 默认为30秒
	Header  map[string]string // 请求头
	Query   map[string]string // 请求参数
	JSON    any               // 请求体
}

func (receiver *JsonOptions) SetTimeout(timeout time.Duration) {
	receiver.timeout = timeout
}

// Url 请求地址
func (receiver *JsonOptions) Url() string {
	return receiver.Host + receiver.Uri
}

// BodyBuffer 请求体Buffer
func (receiver *JsonOptions) BodyBuffer() *bytes.Buffer {
	if receiver.JSON == nil {
		return bytes.NewBuffer([]byte{})
	}
	body, _ := json.Marshal(receiver.JSON)
	return bytes.NewBuffer(body)
}

// BodyString 请求体字符串
func (receiver *JsonOptions) BodyString() string {
	if receiver.JSON == nil {
		return ""
	}
	body, _ := json.Marshal(receiver.JSON)
	return string(body)
}

// GetSignQueryStr 带签名的请求参数
func (receiver *JsonOptions) GetSignQueryStr() string {
	query := receiver.Query
	queryStr := ""
	signStr := ""
	if _, ok := query[SignPtrAppKey]; ok {
		queryStr += SignPtrAppKey + "=" + query[SignPtrAppKey] + "&"
		signStr = SignPtrAppKey + query[SignPtrAppKey]
	}

	if _, ok := query[SingPtrShopId]; ok {
		queryStr += SingPtrShopId + "=" + query[SingPtrShopId] + "&"
		signStr += SingPtrShopId + query[SingPtrShopId]
	}
	if _, ok := query[SignPtrBrandId]; ok {
		queryStr += SignPtrBrandId + "=" + query[SignPtrBrandId] + "&"
		signStr += SignPtrBrandId + query[SignPtrBrandId]
	}

	if _, ok := query[SignPtrTimestamp]; ok {
		queryStr += SignPtrTimestamp + "=" + query[SignPtrTimestamp] + "&"
		signStr += SignPtrTimestamp + query[SignPtrTimestamp]
	}

	if _, ok := query[SignPtrVersion]; ok {
		queryStr += SignPtrVersion + "=" + query[SignPtrVersion] + "&"
		signStr += SignPtrVersion + query[SignPtrVersion]
	}

	if receiver.JSON != nil {
		signStr += SignPtrBody + util.Any(receiver.JSON).ToJson()
	}

	if _, ok := query[SignPtrSecretKey]; ok {
		signStr += query[SignPtrSecretKey]
	}

	if _, ok := query[SignPtrToken]; ok {
		signStr += query[SignPtrToken]
	}

	queryStr += SingPtrSign + "=" + util.Sha256(signStr)

	return queryStr
}
