package middleware

import (
	"net/http"

	"github.com/Breeze0806/go-admin-core/sdk/api"
	"github.com/Breeze0806/go-admin-core/sdk/pkg/jwtauth"
	"github.com/gin-gonic/gin"
)

// AuthCheckRole 权限检查中间件
func AuthCheckRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := api.GetRequestLogger(c)
		data, _ := c.Get(jwtauth.JwtPayloadKey)
		v := data.(jwtauth.MapClaims)
		if v["rolekey"] == "admin" {
			c.Next()
			return
		}
		log.Warnf("role: %s method: %s path: %s message: %s", v["rolekey"], c.Request.Method, c.Request.URL.Path, "当前request无权限，请管理员确认！")
		c.JSON(http.StatusOK, gin.H{
			"code": 403,
			"msg":  "对不起，您没有该接口访问权限，请联系管理员",
		})
		c.Abort()
	}
}
