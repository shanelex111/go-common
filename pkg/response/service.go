package response

import "github.com/gin-gonic/gin"

func Failed(c *gin.Context, e *Error) {
	c.AbortWithStatusJSON(e.Status, e)
}
