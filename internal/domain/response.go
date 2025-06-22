package domain

type Response struct {
	StatusCode int         `json:"statusCode"`
	ResMessage string      `json:"responseMessage"`
	Data       interface{} `json:"data"`
}
