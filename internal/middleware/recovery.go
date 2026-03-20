package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	apperrors "GolangTemplate/internal/errors"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logrus.Errorf("panic recovered: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": apperrors.ErrInternalServerError.Error()})
				c.Abort()
			}
		}()
		c.Next()
	}
}
