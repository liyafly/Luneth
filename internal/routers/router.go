package routers

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	apiV1 := r.Group("/api/v1")
	{
		apiV1.POST("tags")
		apiV1.DELETE("/tags/:id")
		apiV1.PUT("/tags/:id")
		apiV1.PATCH("/tags/:id/state")
		apiV1.GET("/tags")

		apiV1.POST("articles")
		apiV1.DELETE("/articles/:id")
		apiV1.PUT("/articles/:id")
		apiV1.PATCH("/articles/:id/state")
		apiV1.GET("/articles/:id")
		apiV1.GET("/articles")

	}
	return r
}
