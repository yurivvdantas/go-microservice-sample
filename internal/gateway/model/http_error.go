package model

type HTTPError struct {
	Cause  string `json:"cause"`
	Detail string `json:"detail"`
	Status int    `json:"status"`
}
