package responses

import (
	"time"

	"github.com/google/uuid"
)

type BaseCategory struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CategoryDetail struct {
	BaseCategory
}
