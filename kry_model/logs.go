package kry_model

type ReqCtx struct {
	ReqTime     string            `json:"time"`         // 请求时间
	RespTime    string            `json:"responseTime"` // 响应时间
	TTL         int               `json:"ttl"`          // 请求耗时
	Method      string            `json:"method"`       // 请求方法
	ContentType string            `json:"contentType"`  // 请求类型
	Host        string            `json:"host"`         // 请求主机
	URI         string            `json:"uri"`          // 请求URI
	ReqQuery    map[string]string `json:"requestQuery"` // 请求参数
	ReqBody     any               `json:"requestBody"`  // 请求体
	RespBody    any               `json:"responseBody"` // 响应体
}
