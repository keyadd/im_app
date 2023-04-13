package middleware

import (
	"github.com/gin-gonic/gin"
	"im_app/global"
	"im_app/initialize"
	"im_app/utils"
)

//const CtxUserIDKey = "userID"

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// Authorization: Bearer xxxxxxx.xxx.xxx  / X-TOKEN: xxx.xxx.xx
		// 这里的具体实现方式要依据你的实际业务情况决定
		TokenHeaderName := global.GVA_CONFIG.JWT.TokenHeaderName
		//fmt.Println(TokenHeaderName)
		authHeader := c.Request.Header.Get(TokenHeaderName)
		if authHeader == "" {
			utils.ResponseError(c, global.CodeNeedLogin)
			c.Abort()
			return
		}

		// 获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := initialize.ParseToken(authHeader)
		if err != nil {
			utils.ResponseError(c, global.CodeInvalidToken)
			c.Abort()
			return
		}
		token, err := initialize.RefreshToken(mc)
		if err != nil {
			utils.ResponseError(c, global.CodeInvalidToken)
			c.Abort()
			return
		}
		//fmt.Println(token)
		c.Request.Header.Set("Authorization", token)

		// 将当前请求的userID信息保存到请求的上下文c上
		UserIdName := global.GVA_CONFIG.JWT.UserIdName
		//fmt.Println(UserIdName)
		c.Set(UserIdName, mc.UserID)

		c.Next() // 后续的处理请求的函数中 可以用过c.Get(CtxUserIDKey) 来获取当前请求的用户信息
	}
}
