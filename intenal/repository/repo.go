package repository

import (
	"context"

	desc "github.com/Dashinamzh/auth/pkg/auth_v1"
)

type AuthRepository interface {
	Create(ctx context.Context, Info *desc.UserInfo) (int64, error)
	Get(ctx context.Context, id int64) (desc.UserInfo, error)
	Update(ctx context.Context, id int64, name string, email string) error
	Delete(ctx context.Context, id int64) error
}
