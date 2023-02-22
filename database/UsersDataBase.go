package database

import "google.golang.org/protobuf/types/known/timestamppb"

type UsersUpdateData struct {
	ID        uint                  `json:"id" form:"id" binding:"required"`
	Username  string                `json:"username" form:"username" binding:"required"`
	Email     string                `json:"email" form:"email" binding:"required"`
	Password  string                `json:"password" form:"password" binding:"required" validate:"min:6"`
	CreatedAt timestamppb.Timestamp `json:"createdAt" form:"createdAt"`
	UpdatedAt timestamppb.Timestamp `json:"updatedAt" form:"updatedAt"`
}
