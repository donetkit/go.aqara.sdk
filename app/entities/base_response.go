package entities

type BaseResponse struct {
	Code int `json:"code"`
	RequestId string `json:"requestId"`
	Message string `json:"message"`
	Result interface{} `json:"result"`
}


