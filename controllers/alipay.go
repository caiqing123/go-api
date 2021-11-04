package controllers

import (
	"api/configor"
	"api/model"
	"api/pay"
	"api/pay/alipay"
	"api/pkg/util"
	"api/pkg/xlog"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AlipayValidate struct {
	Types   string  `form:"type" validate:"required" label:"类型"`
	Subject string  `form:"subject" validate:"required" label:"商品名称"`
	Amount  float64 `form:"amount" validate:"required,gt=0" label:"金额"`
}

type AlipayData struct {
	Url string `json:"url"`
}

type AlipayController struct {
}

// alipay
// @Summary 支付宝支付
// @Description 支付宝支付下单
// @Tags 支付宝支付
// @Produce application/json
// @Param type formData string true "支付类型" default(h5)
// @Param subject formData string true "商品名称" default(测试)
// @Param amount formData number true "金额" default(0.01)
// @Success 200 {object} pay.Response{data=controllers.AlipayData}
// @Failure 500 {object} pay.ResponseError
// @Failure 422 {object} pay.ResponseVerificationErr
// @Router /alipay/ [post]
func (t *AlipayController) Index(c *gin.Context) {
	//验证
	var param AlipayValidate
	err := c.ShouldBind(&param)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Error("param:", param.Types)
	verificationErr := model.InitTrans(param)
	if verificationErr != nil {
		c.JSON(http.StatusUnprocessableEntity, pay.ResponseVerificationErr{Code: http.StatusUnprocessableEntity, Message: verificationErr})
		return
	}

	config := configor.Config

	client, err := alipay.NewClient(config.Alipay.AppId, config.Alipay.PrivateKey, config.Alipay.IsProd)
	if err != nil {
		xlog.Error(err)
		return
	}
	//配置公共参数
	client.SetCharset("utf-8").
		SetSignType(alipay.RSA2).
		SetNotifyUrl(config.Alipay.NotifyUrl)

	subject := param.Subject
	amount := param.Amount
	tradeNo := util.GetRandomString(32)
	//data := make(map[string]interface{})
	data := AlipayData{}

	switch param.Types {

	case "h5":
		bm := make(pay.BodyMap)
		bm.Set("subject", subject)
		bm.Set("out_trade_no", tradeNo)
		bm.Set("quit_url", config.Alipay.ReturnUrl)
		bm.Set("total_amount", amount)
		bm.Set("product_code", "QUICK_WAP_WAY")
		//手机网站支付请求
		payUrl, err := client.TradeWapPay(c, bm)
		if err != nil {
			xlog.Error("err:", err)
			return
		}
		data.Url = payUrl
	case "app":
		bm := make(pay.BodyMap)
		bm.Set("subject", subject)
		bm.Set("out_trade_no", tradeNo)
		bm.Set("total_amount", amount)
		//手机APP支付参数请求
		payParam, err := client.TradeAppPay(c, bm)
		if err != nil {
			xlog.Error("err:", err)
			return
		}
		data.Url = payParam
	case "pc":
		bm := make(pay.BodyMap)
		bm.Set("subject", subject)
		bm.Set("out_trade_no", tradeNo)
		bm.Set("total_amount", amount)
		bm.Set("product_code", "FAST_INSTANT_TRADE_PAY")

		//电脑网站支付请求
		payUrl, err := client.TradePagePay(c, bm)
		if err != nil {
			xlog.Error("err:", err)
			return
		}
		data.Url = payUrl
	default:
		c.JSON(http.StatusInternalServerError, pay.ResponseError{Code: http.StatusInternalServerError, Message: "not type"})
		return
	}

	c.JSON(http.StatusOK, pay.Response{Code: http.StatusOK, Message: pay.SUCCESS, Data: data})
}
