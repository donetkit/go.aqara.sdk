package entities

/*

{
    "code": 0,
    "message": "Success",
    "requestId": "2379.74406.16088813148026799",
    "result": "EtuvencFksfaUesYTSiYw3axj5UXnuweg="
}

*/

type AiotResult struct {
	Code  int `json:"code"`
	Message string `json:"message"`
	RequestId string `json:"requestId"`
	Result string `json:"result"`
}