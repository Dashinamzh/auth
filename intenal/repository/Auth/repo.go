package auth

import (
	"context"

	"github.com/Dashinamzh/auth/intenal/repository/Auth/model"
	desc "github.com/Dashinamzh/auth/pkg/auth_v1"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	tableName = "auth"

	idColumn        = "id"
	nameColumn      = "name"
	emailColumn     = "email"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db *pgxpool.Pool
}

func (r *repo) Create(ctx context.Context, Info *desc.UserInfo) {
	builder := sq.Insert(tableName).
		Columns(nameColumn, emailColumn).
		Values(Info.Name, Info.Email).
		Suffix("RETURNING id")
	query, args, err := builder.ToSql()
	if err != nil {
		return nil, error
	}

	var id int64
	err = r.db.QueryRow(ctx, query, args).Scan(&id)
	if err != nil {
		return
	}

	return
}

func (r *repo) Get(ctx context.Context, id int64) {
	builder := sq.Select(idColumn, nameColumn, emailColumn, createdAtColumn, updatedAtColumn).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return
	}

	var auth model.Auth
	err = r.db.QueryRow(ctx, query, args...).Scan(&auth.Id, &auth.Info.Name, &auth.Info.Email,
		&auth.CreateAt, &auth.UpdatedAt)
	if err != nil {
		return
	}
}
