package handler

import (
	"net/http"
	"strconv"

	apperrors "GolangTemplate/internal/errors"

	"GolangTemplate/internal/modules/user/dto"
	"GolangTemplate/internal/modules/user/service"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

// GetByID godoc
// @Summary Get user by id
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} model.User
// @Router /users/{id} [get]
func (h *UserHandler) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": apperrors.ErrInvalidID.Error()})
		return
	}

	user, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": apperrors.ErrUserNotFound.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Update godoc
// @Summary Update user
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param input body dto.UpdateUserRequest true "User data"
// @Success 200 {object} model.User
// @Router /users/{id} [put]
func (h *UserHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": apperrors.ErrInvalidID.Error()})
		return
	}

	userIDFromToken, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": apperrors.ErrUnauthorized.Error()})
		return
	}

	userRole, exists := c.Get("role")
	if uint(id) != userIDFromToken.(uint) && userRole != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": apperrors.ErrForbiddenCannotEditOtherUsersData.Error()})
		return
	}

	var input dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": apperrors.ErrUserNotFound.Error()})
		return
	}

	if input.Username != nil && *input.Username != user.Username {
		existing, err := h.service.GetByUsername(*input.Username)
		if err == nil && existing != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": apperrors.ErrUsernameAlreadyExists.Error()})
			return
		}
		user.Username = *input.Username
	}

	if input.Email != nil {
		user.Email = *input.Email
	}
	if input.Phone != nil {
		user.Phone = *input.Phone
	}
	if input.Password != nil {
		hash, err := bcrypt.GenerateFromPassword([]byte(*input.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": apperrors.ErrFailedToHashPassword.Error()})
			return
		}
		user.Password = string(hash)
	}

	if err := h.service.Update(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

// Delete godoc
// @Summary Delete user
// @Tags users
// @Param id path int true "User ID"
// @Success 204
// @Router /users/{id} [delete]
func (h *UserHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": apperrors.ErrInvalidID.Error()})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
