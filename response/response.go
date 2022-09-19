package response

import (
	"google.golang.org/grpc/status"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

const (
	ErrorCodeNormal  = 200
	UnknownErrorCode = 500
)

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	if err == nil {
		body.Code = ErrorCodeNormal
		body.Msg = "success"
		body.Data = resp
		httpx.OkJson(w, body)
		return
	}
	errMap, ok := status.FromError(err)
	if ok == false {
		body.Code = UnknownErrorCode
		body.Msg = err.Error()
		body.Data = resp
		httpx.OkJson(w, body)
		return
	}

	if errMap == nil {
		body.Code = ErrorCodeNormal
		body.Msg = "success"
		body.Data = resp
		httpx.OkJson(w, body)
		return
	}

	body.Code = int(errMap.Code())
	body.Msg = errMap.Message()
	body.Data = resp
	httpx.OkJson(w, body)
}
