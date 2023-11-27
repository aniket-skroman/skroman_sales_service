package middleware

import "github.com/gin-gonic/gin"

// check user has a permission or not to access a modules, manipulate a data
func Permission(ctx *gin.Context) {

	ctx.Next()
}
