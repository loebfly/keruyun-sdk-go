package kry_model

type Result[D any] struct {
	Result      D      `json:"result"`
	Code        int    `json:"code"`
	MessageUuid string `json:"messageUuid,omitempty"`
	Message     string `json:"message"`
	ApiMessage  string `json:"apiMessage,omitempty"`
}

func (receiver Result[D]) IsSuccess() bool {
	return receiver.Code == 0
}
func (receiver Result[D]) ErrorMsg() string {
	if receiver.ApiMessage != "" {
		return receiver.ApiMessage
	} else {
		return receiver.Message
	}
}
