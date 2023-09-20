package service

import (
	"server_frame/pkg/invoker"
)

var (
	User iUserService
)

func Init() error {
	User = NewUserService(invoker.MainDB)
	return nil
}
