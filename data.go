package main

type PhoneRequest struct {
	Phone string `json:"phone"`
}

type VertifyReqeus struct {
	Phone    string `json:"phone"`
	Authcode int    `json:"authcode"`
}
