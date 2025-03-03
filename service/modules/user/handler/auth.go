package handler

import (
	"github.com/DeniesKresna/sined/service/extensions/terror"
	"github.com/DeniesKresna/sined/types/models"
	"github.com/gin-gonic/gin"
)

func (h UserHandler) AuthLogin(ctx *gin.Context) {
	var (
		loginReq models.LoginRequest
		terr     terror.ErrInterface
	)

	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		terr = terror.ErrParameter(err.Error())
		ResponseJson(ctx, terr)
		return
	}

	authResp, terr := h.userUsecase.AuthLogin(ctx, loginReq.Identifier, loginReq.Password)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, authResp)
}

func (h UserHandler) AuthSession(ctx *gin.Context) {
	var (
		terr terror.ErrInterface
	)

	authResp, terr := h.userUsecase.AuthGetFromContext(ctx)
	if terr != nil {
		ResponseJson(ctx, terr)
		return
	}
	ResponseJson(ctx, authResp)
}
