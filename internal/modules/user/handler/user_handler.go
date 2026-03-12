package handler

import (
	"net/http"
	"strconv"

	"GolangTemplate/internal/modules/user/dto"
	"GolangTemplate/internal/modules/user/model"
	"GolangTemplate/internal/modules/user/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

// Create godoc
// @Summary Create user
// @Tags users
// @Accept json
// @Produce json
// @Param input body dto.CreateUserDTO true "User data"
// @Success 201 {object} model.User
// @Router /users [post]
func (h *UserHandler) Create(c *gin.Context) {
	var input dto.CreateUserRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		Name:  input.Name,
		Email: input.Email,
	}

	if err := h.service.Create(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	user, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
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
// @Param input body dto.UpdateUserDTO true "User data"
// @Success 200 {object} model.User
// @Router /users/{id} [put]
func (h *UserHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var input dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	if input.Name != nil {
		user.Name = *input.Name
	}

	if input.Email != nil {
		user.Email = *input.Email
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
