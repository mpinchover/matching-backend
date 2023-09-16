package route

import (
	"matching/src/types/requests"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Fn func(c *gin.Context) (interface{}, error)

func NewRoute(f Fn) gin.HandlerFunc {
	return func(c *gin.Context) {
		resp, err := f(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, &requests.APIErrorResponse{
				Message: err.Error(),
			})
			return
		}
		c.JSON(200, resp)
	}
}
