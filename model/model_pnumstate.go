package model

import (
	"errors"
	"time"
)

var (
	// bucket이 없음
	ErrBucketNotFound = errors.New("error bucket not found")

	// Json Marshal이 생성이 안됌
	ErrJsonMarshalNotWorking = errors.New("error json marshal not working")

	// json unmarshal이 진행이 안됌
	ErrJsonUnmarshalNotWorking = errors.New("error json unmarhsal not working")
)

type PnumStatus struct {
	Cnt       int       `json:"cnt"`
	Timestamp time.Time `json:"timestamp"`
	Break     bool      `json:"break"`
	AuthCode  int       `json:"authCode"`
}
