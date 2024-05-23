package response

import (
	"olin-assessment-muhammad-diffa/schemas"

	"github.com/gin-gonic/gin"
)

func BaseResponseWriter(c *gin.Context, statusCode int, result interface{}, status string, detail string) {
	response := schemas.ResponseBase{
		Result: result,
		Status: status,
		Detail: detail,
	}
	c.IndentedJSON(statusCode, response)
}
