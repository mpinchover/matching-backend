package handlers

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) TestAuth(c *gin.Context) (interface{}, error) {
	return map[string]interface{}{
		"success": true,
	}, nil
}
