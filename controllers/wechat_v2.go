package controllers

import (
	"api/configor"
	"api/model"
	"api/pay"
	"api/pay/wechat"
	"api/pkg/util"
	"api/pkg/xlog"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type WechatV2Validate struct {
	Types   string  `form:"type" validate:"required,oneof=h5 app min" label:"类型"`
	Subject string  `form:"subject" validate:"required" label:"商品名称"`
	Amount  float64 `form:"amount" validate:"required,gt=0" label:"金额"`
}

type WechatV2Data struct {
	Url     string `json:"url"`
	PaySign string `json:"pay_sign"`
}

type WechatV2PayController struct {
}

// Index wechat
// @Summary 微信支付v2
// @Description 微信支付下单
// @Tags 微信支付
// @Produce application/json
// @Param type formData string true "支付类型" default(h5)
// @Param subject formData string true "商品名称" default(测试)
// @Param amount formData number true "金额" default(0.01)
// @Success 200 {object} pay.Response{data=controllers.WechatV2Data}
// @Failure 500 {object} pay.ResponseError
// @Failure 422 {object} pay.ResponseVerificationErr
// @Router /wechat/v2/ [post]
func (t *WechatV2PayController) Index(c *gin.Context) {
	//验证
	var param WechatV2Validate
	err := c.ShouldBind(&param)
	if err != nil {
		xlog.Error(err)
		return
	}
	verificationErr := model.InitTrans(param)
	if verificationErr != nil {
		c.JSON(http.StatusUnprocessableEntity, pay.ResponseVerificationErr{Code: http.StatusUnprocessableEntity, Message: verificationErr})
		return
	}
	config := configor.Config.Wechat
	client := wechat.NewClient(config.AppId, config.MchId, config.ApiKey, config.IsProd)
	number := util.GetRandomString(32)
	//设置国家
	client.SetCountry(wechat.China)
	//初始化参数Map
	bm := make(pay.BodyMap)
	bm.Set("nonce_str", util.GetRandomString(32)).
		Set("body", param.Subject).
		Set("out_trade_no", number).
		Set("total_fee", param.Amount*100).
		Set("spbill_create_ip", "127.0.0.1").
		Set("notify_url", config.NotifyUrl).
		Set("trade_type", wechat.TradeType_H5).
		Set("device_info", "WEB").
		Set("sign_type", wechat.SignType_MD5).
		SetBodyMap("scene_info", func(bm pay.BodyMap) {
			bm.SetBodyMap("h5_info", func(bm pay.BodyMap) {
				bm.Set("type", "Wap")
				bm.Set("wap_url", config.ReturnUrl)
				bm.Set("wap_name", param.Subject)
			})
		}) /*.Set("openid", "xxx")*/

	//请求支付下单，成功后得到结果
	wxRsp, err := client.UnifiedOrder(c, bm)
	if err != nil {
		xlog.Error(err)
		return
	}
	xlog.Debug("Response：", wxRsp)

	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)

	data := WechatV2Data{}

	switch param.Types {

	case "app":
		paySign := wechat.GetAppPaySign(config.AppId, "", wxRsp.NonceStr, wxRsp.PrepayId, wechat.SignType_MD5, timeStamp, config.ApiKey)
		data.Url = wxRsp.MwebUrl
		data.PaySign = paySign
	case "h5":
		pac := "prepay_id=" + wxRsp.PrepayId
		paySign := wechat.GetJsapiPaySign(config.AppId, wxRsp.NonceStr, pac, wechat.SignType_MD5, timeStamp, config.ApiKey)
		data.Url = wxRsp.MwebUrl
		data.PaySign = paySign
	case "min":
		pac := "prepay_id=" + wxRsp.PrepayId
		paySign := wechat.GetMiniPaySign(config.AppId, wxRsp.NonceStr, pac, wechat.SignType_MD5, timeStamp, config.ApiKey)
		data.Url = wxRsp.MwebUrl
		data.PaySign = paySign
	default:
		c.JSON(http.StatusInternalServerError, pay.ResponseError{Code: http.StatusInternalServerError, Message: "not type"})
		return
	}

	c.JSON(http.StatusOK, pay.Response{Code: http.StatusOK, Message: pay.SUCCESS, Data: data})
}

func (t *WechatV2PayController) Notify(c *gin.Context) {
	rsp := new(wechat.NotifyResponse)

	// 解析参数
	bodyMap, err := wechat.ParseNotifyToBodyMap(c.Request)
	if err != nil {
		xlog.Debug("err:", err)
		rsp.ReturnCode = pay.FAIL
		c.XML(http.StatusOK, rsp.ToXmlString())
		return
	}
	xlog.Debug("bodyMap:", bodyMap)

	config := configor.Config.Wechat

	ok, err := wechat.VerifySign(config.ApiKey, wechat.SignType_MD5, bodyMap)
	if err != nil {
		xlog.Debug("err:", err)
		rsp.ReturnCode = pay.FAIL
		c.XML(http.StatusOK, rsp.ToXmlString())
		return
	}
	xlog.Debug("微信验签是否通过:", ok)

	rsp.ReturnCode = pay.SUCCESS
	rsp.ReturnMsg = pay.OK
	c.XML(http.StatusOK, rsp.ToXmlString())
}
