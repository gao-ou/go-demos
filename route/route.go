package route

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-demos/api"
	"net/http"
	"time"
)

func InitRoute() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		MaxAge:          12 * time.Hour,
	})).GET("/albums", api.PhotoAlbumList)

	{
		//r.GET("/albums/", api.PhotoAlbumList)
		r.POST("/albums", api.CreatePhotoAlbum)
	}

	fmt.Println("初始化路由")

	err := r.Run(":8080")
	if err != nil {
		fmt.Println("gin 启动错误:", err.Error())
	}
}

// Cors 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:63342")
		//c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Problem-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Expose-Headers", "Problem-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Problem-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法，因为有的模板是要请求两次的
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		// 处理请求
		c.Next()
	}
}
