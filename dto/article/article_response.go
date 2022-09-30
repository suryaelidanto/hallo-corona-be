package articledto

import (
	"hallocorona/models"
	"time"
)

type ArticleResponse struct {
	ID          int         `json:"id"`
	Title       string      `json:"title"`
	Image       string      `json:"image"`
	Description string      `json:"description"`
	UserID      int         `json:"-"`
	User        models.User `json:"user"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
}
