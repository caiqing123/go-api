package controllers

import (
	"api/configor"
	"api/model"
	"api/pay"
	"api/pay/wechat/v3"
	"api/pkg/util"
	"api/pkg/xlog"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type WechatV3Validate struct {
	Types   string  `form:"type" validate:"required,oneof=h5 app min sm" label:"类型"`
	Subject string  `form:"subject" validate:"required" label:"商品名称"`
	Amount  float64 `form:"amount" validate:"required,gt=0" label:"金额"`
}

type WechatV3Data struct {
	Url     string `json:"url"`
	PaySign string `json:"pay_sign"`
}

type WechatV3PayController struct {
}

var PrivateKey = `-----BEGIN PRIVATE KEY-----
MIIEwAIBADANBgkqhkiG9w0BAQEFAASCBKowggSmAgEAAoIBAQDV523KVXZaaZI3
WxQiaz0J8o8nxAYsxBjrfcaKPnLo+r5uFME7GPV+4UHEZWG6ZogJ87yBt8L4IV8q
/2n0MPKV5qNtS0htG7G0Mtyw7lPmdXUXsA0ionbL2mzz0kgJ1S6FJwyZRRZNJ08Q
/GQE3TWqErbxL/2ITuzTeHrdTNL0i9oNxtB92EWFZ0gSL677zEiz41EVog24SyOd
TFqxjGFd9DR0CeRNU/oQPplFnM9YFseRuhEZ/jLATEvubH/U1qGqTlW0UHvIn14j
NqRxyAjDI/HfXl3Bo7Fx0QCJkVkqb+5ou8KFRchbcixRU0khbrxTy7dDJj60vSmr
PySqqZLFAgMBAAECggEBAKHPN9ZfX/B0/A6z7z86MCpeOryyJJmondFGi/H326Uy
SOus959k+hDJBZ8zsgH3neEpZ+gYwnxBgmRcYiI/BMMwfWAoGtmuoXbXIusU3pLv
N2x72PPiQktjKBgpciU+BrrjFzy6bmxe2AjZZC/pxrapAYrh6sA6NBykfwz5GHu0
DQmjHYqSlghDDljCzVR3Gcs/KicCMw6eQ0JlWDqtDEDoENlBkfn4spHwocV7HtSq
0bnUrQqqMtpZjbMJzZxJc39qkyNNDosuNy5GXYLQE7lr9RuRqLxEfg6KfGUS5bAZ
eJ5pizql7+c0viUtiXG17PYp8QR4c5G+54RlQd1pPuECgYEA9UBi5rFJzK0/n4aO
lsrp6BvUOSherp57SNYvpsRuBPU0odyH2/McLNphisKTxfSm0/hADaTmnzAnOUVg
cduc/5/5tVaaqyLL3SemxJhwqVsL3tE/KAN7HUBhhQrqD+H8r39TAoIkyfjCOHzS
74rygZ35x0kXNMavXQFB0RE2fEcCgYEA30dWaLddGmTvUXwhyTWcsiDfrsKbw8+n
MhAlSCXE42v9Uo3ULqD3/rpUQlMhoqaZb3cSyOyQwJvv0tp/g3hM7Q4usLxkdysc
KA9HmmZ4Q2P2838cqvNr/Dz1UAnfdDryMEnbiv1MUKYqFFTVX6oR0iH544JgDFCG
YLQA2M+3GpMCgYEAh+ax51v+pSirxN5vTSgMDc69/x5buS+g6W+m4CahQKYQEFGA
B2XkCwbIXngMIvm7KGK8O9NQ6I1qbtX+55jmmtAvM0lWU9boWRiL1Q0UAQSuwz34
XVfwdPkkEPFHWp3DxAwuF4m+kR0DowGocYzxbNn5e3EJJvmiW0tDCXMcWikCgYEA
tyNxWcUFBdBCh+i0YbCqzWSvdE3Fq8/YSPT7T3lDTHLYPu18W57Gq1Y0JI7BaQMT
mVzmuI1pkcKV7LIxoyl6l3ppi6eLFD/1AVq/FYL1I/mLpl/dqM6vBR8O686dTV3I
Jxl9jTyEayZQH4sR1TzPDze1GwpmM9Oc1RbwFuYRPycCgYEAzYaRKh6EQ+s37HDv
e/ZGMs3PI+CoA/x6lx4Owa7amRsWRKys45NV6gcC8pkbN4IeFaYXVHmJ1Yaef3xn
0VxHAzWI4BF+1pUwXzS2rAMBZR/VKS0XA856NauAC3mKHipoOWVVs+uFP3VMUQ79
hSImAa7UBzss6b6ie7AYxXtZBjY=
-----END PRIVATE KEY-----`

// Index wechat
// @Summary 微信支付v3
// @Description 微信支付下单
// @Tags 微信支付
// @Produce application/json
// @Param type formData string true "支付类型" default(h5)
// @Param subject formData string true "商品名称" default(测试)
// @Param amount formData number true "金额" default(0.01)
// @Success 200 {object} pay.Response{data=controllers.WechatV3Data}
// @Failure 500 {object} pay.ResponseError
// @Failure 422 {object} pay.ResponseVerificationErr
// @Router /wechat/v3/ [post]
func (t *WechatV3PayController) Index(c *gin.Context) {
	//验证
	var param WechatV3Validate
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
	// NewClientV3 初始化微信客户端 V3
	config := configor.Config.Wechat
	client, err := wechat.NewClientV3(config.MchId, config.SerialNo, config.ApiV3Key, PrivateKey)
	if err != nil {
		xlog.Error(err)
		return
	}
	// 启用自动同步返回验签，并定时更新微信平台API证书
	//err = client.AutoVerifySign()
	//if err != nil {
	//	xlog.Error(err)
	//	return
	//}
	// 打开Debug开关，输出日志
	client.DebugSwitch = pay.DebugOff

	number := util.GetRandomString(32)
	expire := time.Now().Add(10 * time.Minute).Format(time.RFC3339)
	//初始化参数Map
	bm := make(pay.BodyMap)
	bm.Set("description", param.Subject).
		Set("out_trade_no", number).
		Set("time_expire", expire).
		Set("notify_url", config.NotifyUrl)

	data := WechatV3Data{}

	switch param.Types {

	case "app":
		bm.Set("appid", config.AppId).
			SetBodyMap("amount", func(bm pay.BodyMap) {
				bm.Set("total", param.Amount*100).
					Set("currency", "CNY")
			})
		wxRsp, err := client.V3TransactionApp(c, bm)
		if err != nil {
			xlog.Error(err)
			return
		}
		xlog.Info("wxRsp:", wxRsp)
		data.Url = wxRsp.Response.PrepayId
	case "h5":
		bm.Set("appid", config.AppId).
			SetBodyMap("amount", func(bm pay.BodyMap) {
				bm.Set("total", param.Amount*100).
					Set("currency", "CNY")
			}).
			SetBodyMap("scene_info", func(bm pay.BodyMap) {
				bm.Set("payer_client_ip", c.ClientIP).
					SetBodyMap("h5_info", func(bm pay.BodyMap) {
						bm.Set("type", "Wap")
					})
			})
		wxRsp, err := client.V3TransactionH5(c, bm)
		if err != nil {
			xlog.Error(err)
			return
		}
		xlog.Info("wxRsp:", wxRsp)
		data.Url = wxRsp.Response.H5Url
	case "sm":
		bm.Set("appid", config.AppId).
			SetBodyMap("amount", func(bm pay.BodyMap) {
				bm.Set("total", param.Amount*100).
					Set("currency", "CNY")
			})
		wxRsp, err := client.V3TransactionNative(c, bm)
		if err != nil {
			xlog.Error(err)
			return
		}
		xlog.Info("wxRsp:", wxRsp)
		data.Url = wxRsp.Response.CodeUrl
	case "min":
		bm.Set("sp_appid", "xxx").
			Set("sp_mchid", "xxx").
			Set("sub_mchid", "xxx").
			SetBodyMap("amount", func(bm pay.BodyMap) {
				bm.Set("total", param.Amount*100).
					Set("currency", "CNY")
			}).
			SetBodyMap("payer", func(bm pay.BodyMap) {
				bm.Set("sp_openid", "asdas")
			})
		//text, err := wechat.V3EncryptText("张三")
		//if err != nil {
		//	xlog.Errorf("client.V3EncryptText(),err:%+v", err)
		//	err = nil
		//}
		//xlog.Debugf("加密text: %s", text)

		wxRsp, err := client.V3TransactionJsapi(c, bm)
		if err != nil {
			xlog.Error(err)
			return
		}
		xlog.Info("wxRsp:", wxRsp)

		data.Url = wxRsp.Response.PrepayId
	default:
		c.JSON(http.StatusInternalServerError, pay.ResponseError{Code: http.StatusInternalServerError, Message: "not type"})
		return
	}

	c.JSON(http.StatusOK, pay.Response{Code: http.StatusOK, Message: pay.SUCCESS, Data: data})
}

func (t *WechatV3PayController) Notify(c *gin.Context) {

	// NewClientV3 初始化微信客户端 V3
	config := configor.Config.Wechat
	client, err := wechat.NewClientV3(config.MchId, config.SerialNo, config.ApiV3Key, PrivateKey)
	if err != nil {
		xlog.Error(err)
		return
	}
	// 解析参数
	bodyMap, err := wechat.V3ParseNotify(c.Request)
	if err != nil {
		xlog.Debug("err:", err)
		return
	}
	xlog.Debug("bodyMap:", bodyMap)
	ok := bodyMap.VerifySignByPK(client.WxPublicKey())
	xlog.Debug("ok:", ok)
	ok1, err := bodyMap.DecryptCipherText(config.ApiV3Key)
	if err != nil {
		xlog.Debug("err:", err)
		return
	}
	xlog.Debug("ok1:", ok1)
}
