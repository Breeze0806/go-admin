package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/Breeze0806/go-admin-core/sdk/pkg/jwtauth"
	"go-admin/app/other/apis"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerFileRouter)
}

// 需认证的路由代码
func registerFileRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	var api = apis.File{}
	r := v1.Group("").Use(authMiddleware.MiddlewareFunc())
	{
		r.POST("/public/uploadFile", api.UploadFile)
	}
}
