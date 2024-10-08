package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kriserohalia/SI-COMPANY-PROFILE/server/driver/middleware"
	"github.com/kriserohalia/SI-COMPANY-PROFILE/server/helper"
	"github.com/kriserohalia/SI-COMPANY-PROFILE/server/model"
	"github.com/kriserohalia/SI-COMPANY-PROFILE/server/usecase"
)

type authController struct {
	am middleware.AuthMiddleware
	uc usecase.AuthUsecase
	rg *gin.RouterGroup
}

func (au *authController) login(ctx *gin.Context){
	var loginReq model.Login

	if err := ctx.ShouldBindJSON(&loginReq);err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message" :err.Error()})
		return 
	}
	fmt.Println("PRINT ", loginReq)
	rest,err := au.uc.Login(loginReq.Email,loginReq.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message" :err.Error()})
		return 
	}
	role:= helper.HelperGetRole(rest.TokenAccess)

	ctx.JSON(http.StatusOK,gin.H{"message" :"Login Success", "data" :rest,"role" :role})
}


func (au *authController) register (ctx *gin.Context){
	var regisReq model.Register

	if err := ctx.ShouldBindJSON(&regisReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message" :err.Error()})
	}

	err := au.uc.Register(regisReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message" :err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message" :"Register Success"})

}

func (au *authController) resetPassword (ctx *gin.Context) {
	var ResetReq struct {

		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	if err := ctx.ShouldBindJSON(&ResetReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message" : err.Error()})
	}
	 email,exist := ctx.Get("Email")
	 if !exist {
		ctx.JSON(http.StatusForbidden,gin.H{"message":"user not found"})
		return 
	 }

	 id,exist := ctx.Get("Id")
	 if !exist {
		ctx.JSON(http.StatusForbidden,gin.H{"message":"user not found"})
		return 
	 }

	err := au.uc.ResetPassword(ResetReq.OldPassword,ResetReq.NewPassword,email.(string),int(id.(float64)))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message" : err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message" :"Reset Password Success"})

}

func (au *authController) AuthRouter(){
	r:= au.rg.Group("auth")
	r.POST("/login",au.login)
	r.POST("/register",au.register)
	r.POST("/reset-password",au.am.JwtVerified("CUSTOMER","OWNER") ,au.resetPassword)

}


func NewAuthController(am middleware.AuthMiddleware, uc usecase.AuthUsecase, rg *gin.RouterGroup) *authController{
	return &authController{
		am: am,
		uc: uc,
		rg: rg,
	}
	
}