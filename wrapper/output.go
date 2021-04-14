package wrapper

import (
	"net/http"

	"github.com/edgex-go-api/resp"
	"github.com/gin-gonic/gin"
)

type JsonOutput struct {
	context    *gin.Context
	HttpStatus int
	Resp       interface{}
}

func NewJsonOutput(c *gin.Context, httpStatus int, rsp interface{}) *JsonOutput {
	return &JsonOutput{
		context:    c,
		HttpStatus: httpStatus,
		Resp:       rsp,
	}
}

func SampleJson(c *gin.Context, p resp.ErrorCode, data interface{}) *JsonOutput {
	return NewJsonOutput(c, http.StatusOK, resp.NewStdResponse(p, data))
}

func (s *JsonOutput) Write() {
	s.context.JSON(s.HttpStatus, s.Resp)
}
