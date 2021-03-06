package helpers

import "github.com/gin-gonic/gin"

func SetCorsHeaders(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	c.Header("Content-Type", `application/json`)
	c.Header("Accept", `application/json`)
}
