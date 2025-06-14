package response

import "github.com/gin-gonic/gin"

func Failed(c *gin.Context, e *response) {
	c.AbortWithStatusJSON(e.Status, e)
}
