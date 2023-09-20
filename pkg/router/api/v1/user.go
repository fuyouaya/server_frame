package v1

import (
	"github.com/gin-gonic/gin"
	"server_frame/pkg/errs"
	"server_frame/pkg/model/view"
	"server_frame/pkg/service"
)

func UserLogin(c *gin.Context) error {
	req := new(view.UserLoginReq)
	if err := c.ShouldBind(req); err != nil {
		return errs.ErrParamInvalid
	}

	if err := service.User.Login(c, req); err != nil {
		return err
	}

	c.JSON(200, "success")
	return nil
}
