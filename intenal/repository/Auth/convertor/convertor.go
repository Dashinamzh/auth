package convertor

import (
	"github.com/Dashinamzh/auth/intenal/repository/Auth/model"
	desc "github.com/Dashinamzh/auth/pkg/auth_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func toAuthFromRepo(auth model.Auth) *desc.User {
	var updated_at *timestamppb.Timestamp
	if auth.UpdatedAt.Valid {
		updated_at = timestamppb.New(auth.UpdatedAt.Time)
	}
	return &desc.User{
		Id:        auth.Id,
		Info:      ToAuthInfoFromRepo(*auth.Info),
		Role:      auth.Role,
		CreatedAt: timestamppb.New(auth.CreateAt),
		UpdatedAt: updated_at,
	}
}

func ToAuthInfoFromRepo(auth model.AuthInfo) *desc.UserInfo {
	return &desc.UserInfo{
		Name:  auth.Name,
		Email: auth.Email,
	}
}
