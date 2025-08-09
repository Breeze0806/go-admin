package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerMonitorRouter)
}

// 需认证的路由代码
func registerMonitorRouter(v1 *gin.RouterGroup) {
	v1.GET("/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

}
