package handler

import (
	"net/http"

	"github.com/edgex-go-api/resp"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type UserInfoParams struct {
	UserID   int64  `from:"user_id" json:"user_id"`
	Password string `from:"password" json:"password"`
	CityName string `from:"city_name" json:"city_name"`
}

// func GetUserInfo(c *gin.Context) {
// 	logs.Error("arrive")
// 	var params = &UserInfoParams{}
// 	err := binding.Default(http.MethodPost, binding.MIMEJSON).Bind(c.Request, params)
// 	if err != nil {
// 		c.JSON(http.StatusOK, resp.NewStdResp(resp.RESP_CODE_PARAM_ERR, nil))
// 		return
// 	}
// 	params.CityName = c.Query("city_name")
// 	c.JSON(http.StatusOK, resp.NewStdResp(resp.RESP_CODE_SUCCESS, params))
// }

func GetUserInfo(c *gin.Context) {
	h := NewUserInfoHandler(c)

	err := h.CheckParams()
	if err != nil {
		c.JSON(http.StatusOK, resp.NewStdResp(resp.RESP_CODE_PARAM_ERR, nil))
		return
	}
	h.Pack()
	c.JSON(http.StatusOK, resp.NewStdResp(resp.RESP_CODE_SUCCESS, h.Resp))

}

type userInfoHandler struct {
	ReqCtx *gin.Context
	Params *UserInfoParams
	Resp   *UserInfoParams
}

func NewUserInfoHandler(c *gin.Context) *userInfoHandler {
	return &userInfoHandler{
		ReqCtx: c,
	}
}

func (s *userInfoHandler) CheckParams() error {
	var params = &UserInfoParams{}
	err := binding.Default(http.MethodPost, binding.MIMEJSON).Bind(s.ReqCtx.Request, params)
	if err != nil {
		return err
	}
	params.CityName = s.ReqCtx.Query("city_name")
	return nil
}

func (s *userInfoHandler) Pack() {
	s.Resp = s.Params
}
