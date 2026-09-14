package repositories

import (
	"backend-001/internal/models"
	"context"
)

func (r *sqlUserRepository) Create(ctx context.Context,u *models.User)