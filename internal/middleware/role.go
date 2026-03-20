package middleware

import (
	"net/http"

	apperrors "GolangTemplate/internal/errors"

	"github.com/gin-gonic/gin"
)

func RequireRole(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleVal, exists := c.Get("role")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": apperrors.ErrRoleNotFoundInToken.Error(),
			})
			return
		}

		role := roleVal.(string)

		for _, allowed := range roles {
			if role == allowed {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": apperrors.ErrForbidden.Error(),
		})
	}
}
