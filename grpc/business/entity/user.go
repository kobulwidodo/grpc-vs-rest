package entity

import (
	pb "grpc/pb"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Username  string         `json:"username"`
	Password  string         `json:"password"`
	Name      string         `json:"name"`
}

func (u *User) ToGetListData() *pb.GetListData {
	return &pb.GetListData{
		Id:        int64(u.ID),
		CreatedAt: u.CreatedAt.Format(time.RFC3339),
		UpdatedAt: u.UpdatedAt.Format(time.RFC3339),
		DeletedAt: u.DeletedAt.Time.Format(time.RFC3339),
		Username:  u.Username,
		Password:  u.Password,
		Name:      u.Name,
	}
}
