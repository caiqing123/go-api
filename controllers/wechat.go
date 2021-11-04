package controllers

import (
	"api/di"
	"api/pay"
	"api/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type WechatPayController struct {
}

// wechat
// @Summary 微信支付
// @Description 微信支付下单
// @Tags 微信支付
// @Produce application/json
// @Param type formData string true "支付类型" default(h5)
// @Success 200 {object} pay.Response
// @Failure 500 {object} pay.ResponseError
// @Router /wechat/ [post]
func (t *WechatPayController) Index(c *gin.Context) {

	types := c.PostForm("type")
	//types := c.DefaultPostForm("type","h5")

	tradeNo := util.GetRandomString(32)
	di.Zap().Infof("tradeNo %s", tradeNo)

	switch types {

	case "h5":

	case "app":

	default:
		c.JSON(http.StatusInternalServerError, pay.ResponseError{Code: http.StatusInternalServerError, Message: "not type"})
		return
	}

	c.JSON(http.StatusOK, pay.Response{Code: http.StatusOK, Message: "wxpay"})
}
