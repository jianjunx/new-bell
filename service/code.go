package service

import (
	"fmt"
	"net/http"
	"new-bell/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	CodeDefaultFault  = "3000"
	CodePasswdFault   = "3001"
	CodeUserNameFault = "3002"
	CodeUserEmpty     = "3003"
	CodeTokenInvalid  = "3004"
)

var faultTextMap = map[string]string{
	CodeDefaultFault:  "错误：",
	CodePasswdFault:   "密码错误",
	CodeUserNameFault: "用户名错误",
	CodeUserEmpty:     "用户信息为空",
	CodeTokenInvalid:  "Token失效",
}

func ErrorHandler(c *gin.Context, err error) {
	errCode := err.Error()
	text, ok := faultTextMap[errCode]
	if !ok {
		fmt.Println("ok")
		text = faultTextMap[CodeDefaultFault] + errCode
		errCode = CodeDefaultFault
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"code": errCode,
		"msg":  text,
	})
	zap.L().Error(text, zap.Error(fmt.Errorf(errCode)))
}

func SuccessList(c *gin.Context, p *models.PageListParams, list []interface{}, total int) {
	c.JSON(http.StatusOK, gin.H{
		"code":      200,
		"msg":       "ok",
		"list":      list,
		"total":     total,
		"pageSize":  p.PageSize,
		"pageIndex": p.PageIndex,
	})
}
func SuccessData(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": data,
	})
}
