package service

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"server_frame/pkg/model/mysql"
	"server_frame/pkg/model/view"
)

var _ iUserService = (*userService)(nil)

type iUserService interface {
	Login(c *gin.Context, req *view.UserLoginReq) error
}

type userService struct {
	db *gorm.DB
}

func (u *userService) Login(c *gin.Context, req *view.UserLoginReq) error {
	//TODO implement me
	if _, err := mysql.UserInfo(u.db, map[string]interface{}{
		"password": "123",
	}); err != nil {
		println("------------------------->", err.Error())
	}
	return nil
}

func NewUserService(db *gorm.DB) *userService {
	return &userService{db: db}
}
