package usecase

import (
	"github.com/DeniesKresna/sined/service/extensions/terror"
	"github.com/DeniesKresna/sined/types/models"
	"github.com/gin-gonic/gin"
)

func (u UserUsecase) RoleGetByID(ctx *gin.Context, id int64) (role models.Role, terr terror.ErrInterface) {
	return u.userRepo.RoleGetByID(ctx, id)
}
