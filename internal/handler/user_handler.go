package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/ABHI002684/go-user-backend-api/internal/models"
	"github.com/ABHI002684/go-user-backend-api/internal/service"
)

type UserHandler struct {
	service   *service.UserService
	validator *validator.Validate
	logger    *zap.Logger
}

func NewUserHandler(s *service.UserService, l *zap.Logger) *UserHandler {
	return &UserHandler{
		service:   s,
		validator: validator.New(),
		logger:    l,
	}
}

func (h *UserHandler) Create(c *fiber.Ctx) error {
	var req models.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	if err := h.validator.Struct(req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	user, err := h.service.Create(c.Context(), req.Name, req.DOB)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	h.logger.Info("user created", zap.Int32("id", user.ID))
	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *UserHandler) Get(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	user, err := h.service.GetByID(c.Context(), int32(id))
	if err != nil {
		return fiber.ErrNotFound
	}
	return c.JSON(user)
}

func (h *UserHandler) List(c *fiber.Ctx) error {
	users, err := h.service.List(c.Context())
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.JSON(users)
}

func (h *UserHandler) Update(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var req models.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	user, err := h.service.Update(c.Context(), int32(id), req.Name, req.DOB)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.JSON(user)
}

func (h *UserHandler) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := h.service.Delete(c.Context(), int32(id)); err != nil {
		return fiber.ErrInternalServerError
	}
	return c.SendStatus(fiber.StatusNoContent)
}
