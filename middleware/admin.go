package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AdminOnly hanya bisa diakses oleh user dengan status admin (is_admin = true)
func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		isAdmin, exists := c.Get("is_admin")

		if !exists || isAdmin != true {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Akses ditolak, hanya admin yang dapat mengakses resource ini",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
