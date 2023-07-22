package main

import "github.com/gin-gonic/gin"

func main() {
	r := setupRouter()
	r.Run(":4000")
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(JSONMiddleware())

	r.POST("api/v1/anti-fraud/pre-analysis", preAuthHandler)
	r.POST("api/v1/anti-fraud/transactions", transactions)
	r.GET("api/v1/anti-fraud/transactions/:id", getTransaction)

	return r
}

func JSONMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Next()
	}
}
