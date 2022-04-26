package dingding

import (
	"encoding/json"
	"fmt"

	"github.com/lack-io/dingding-sdk/api"
)

type DDErrorV1 struct {
	Code      string `json:"code"`
	RequestId string `json:"requestid"`
	Message   string `json:"message"`
}

func (e *DDErrorV1) Error() string {
	return fmt.Sprintf(`{"requestid": "%s", "code": "%s", "message": "%s"}`, e.RequestId, e.Code, e.Message)
}

func ParseErrV1(b []byte) (*DDErrorV1, bool) {
	err := &DDErrorV1{}
	e := json.Unmarshal(b, err)
	if e != nil {
		return err, false
	}
	return err, true
}

type DDErrorV0 struct {
	Code      string `json:"errcode"`
	Message   string `json:"errmsg"`
	RequestId string `json:"request_id"`
}

func (e *DDErrorV0) Error() string {
	return fmt.Sprintf(`{"code": "%s", "message": "%s", "request_id": "%s"}`, e.Code, e.Message, e.RequestId)
}

func ParseErrV0(b []byte) (*DDErrorV0, bool) {
	err := &DDErrorV0{}
	e := json.Unmarshal(b, err)
	if e != nil {
		return err, false
	}
	return err, true
}

func ParseResultV0(result *api.ResultV0) error {
	if result.ErrCode != 0 {
		if result.SubCode != "" || result.SubMsg != "" {
			return &DDErrorV0{Code: result.SubCode, Message: result.SubMsg, RequestId: result.RequestId}
		}
		return &DDErrorV0{Message: result.ErrMsg, RequestId: result.RequestId}
	}
	return nil
}
