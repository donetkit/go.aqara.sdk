package entities

// 请求head参数
type HeaderV1Request struct {
	Appid       string // `json:"appid"`
	AccessToken string  // `json:"accesstoken"`
	Time        string  // `json:"time"`
}

type HeaderRequest struct {
	Appid string //`json:"appid"`
	//Sign string //`json:"sign"`
	Lang string //`json:"lang"`
	Time string //`json:"time"`
}

