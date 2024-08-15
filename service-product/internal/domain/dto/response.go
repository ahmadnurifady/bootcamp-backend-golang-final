package dto

import "service-product/internal/domain"

type MessageKafka struct {
	OrderType      string `json:"orderType"`
	FromService    string `json:"fromService"`
	TakenByService string `json:"takenByService"`
	TransactionId  string `json:"transactionId"`
	UserId         string `json:"userId"`
	ProductId      string `json:"productId"`
	RespStatus     string `json:"respStatus"`
	RespMessage    string `json:"respMessage"`
	RespCode       int    `json:"respCode"`
}

type BaseResponse struct {
	ResponseCode    int            `json:"responseCode"`
	ResponseMessage string         `json:"responseMessage"`
	Data            domain.Product `json:"data"`
}
