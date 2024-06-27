package dto

type GeneralResponse struct {
	Status   int    `json:"status"`
	Messsage string `json:"message"`
	Data     any    `json:"data"`
}
