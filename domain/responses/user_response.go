package responses

import (
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/models"
	"github.com/google/uuid"
)

type UserLoginResponse struct {
	ID    uuid.UUID   `json:"id"`
	Name  string      `json:"name"`
	Email string      `json:"email"`
	Role  models.Role `json:"role"`
	Token string      `json:"token"`
}

type UserRegisterResponse struct {
	Message string `json:"message"`
}
