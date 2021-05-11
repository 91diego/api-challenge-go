package models

type Response struct {
	Code    int    `int:"code"`
	Status  string `string:"status"`
	Message string `string:"message"`
	Data    string `json:"data"`
}

type JsonResponse []Response
