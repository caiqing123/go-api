package controllers

import (
    "api/model"
    "github.com/dgrijalva/jwt-go"
    "github.com/gin-gonic/gin"
    "github.com/mix-go/dotenv"
    "net/http"
    "time"
)

type AuthController struct {
}



func (t *AuthController) Index(c *gin.Context) {
    var user model.User
    // 检查用户登录代码
    err := c.ShouldBind(&user)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status":  http.StatusInternalServerError,
            "message": err,
        })
        return
    }
    verificationErr :=  model.InitTrans(user)
    if verificationErr != nil {
        c.JSON(http.StatusUnprocessableEntity, gin.H{
            "status":  http.StatusUnprocessableEntity,
            "message": verificationErr,
        })
        return
    }
    // 创建 token
    now := time.Now().Unix()
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "iss": "http://example.org",                                  // 签发人
        "iat": now,                                                   // 签发时间
        "exp": now + int64(7200),                                     // 过期时间
        "nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(), // 什么时间之前不可用
        "uid": 100008,
    })
    tokenString, err := token.SignedString([]byte(dotenv.Getenv("HMAC_SECRET").String()))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status":  http.StatusInternalServerError,
            "message": "Creation of token fails",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "status":       http.StatusOK,
        "message":      "ok",
        "access_token": tokenString,
        "expire_in":    7200,
    })
}
