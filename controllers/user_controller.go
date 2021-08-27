package controllers

import (
	"fmt"
	"goginCasbin/model"
	"goginCasbin/repository"
	"goginCasbin/utils"
	"net/http"
	"strconv"

	"github.com/casbin/casbin/v2"

	"github.com/gin-gonic/gin"
)

// UserController : represent the user's controller contract
type UserController interface {
	AddUser(enforcer *casbin.Enforcer) gin.HandlerFunc
	GetUser(*gin.Context)
	GetAllUser(*gin.Context)
	Login(*gin.Context)
	UpdateUser(*gin.Context)
	DeleteUser(*gin.Context)
}

type userController struct {
	userRepo repository.UserRepository
}

//NewUserController -> returns new user controller
func NewUserController(repo repository.UserRepository) UserController {
	return userController{
		userRepo: repo,
	}
}

func (h userController) GetAllUser(ctx *gin.Context) {
	fmt.Println(ctx.Get("userID"))
	user, err := h.userRepo.GetAllUser()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, user)

}

func (h userController) GetUser(ctx *gin.Context) {
	id := ctx.Param("user")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.userRepo.GetUser(intID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, user)

}

func (h userController) Login(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		//ctx.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	//data := string(user.Email)
	dbUser, err := h.userRepo.GetByEmail(user.Email)

	if err != nil {

		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": dbUser})
		return
	}

	if isTrue := utils.ComparePassword(dbUser.Password, user.Password); isTrue {
		fmt.Println("user before", dbUser.ID)
		token := utils.GenerateToken(dbUser.ID)
		ctx.JSON(http.StatusOK, gin.H{"msg": "Successfully Login", "token": token})
		return
	}
	ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Password not matched"})
	return
}

func (h userController) AddUser(enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user model.User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		utils.HashPassword(&user.Password)
		user, err := h.userRepo.AddUser(user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		enforcer.AddGroupingPolicy(fmt.Sprint(user.ID), user.Role)
		user.Password = ""
		ctx.JSON(http.StatusOK, user)

	}
}

func (h userController) UpdateUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := ctx.Param("user")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user1, err := h.userRepo.GetUser(intID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}

	if user1 == (model.User{}){
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	//user.ID = uint(intID)
	//user, err = h.userRepo.UpdateUser(user1)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user1.Name = user.Name
	user1.Email = user.Email
	user1.Role = user.Role
	utils.HashPassword(&user.Password)
	user1.Password = user.Password
	h.userRepo.UpdateUser(user1)
	ctx.JSON(http.StatusOK, user1)

}

func (h userController) DeleteUser(ctx *gin.Context) {
	var user model.User
	id := ctx.Param("user")
	intID, _ := strconv.Atoi(id)
	user.ID = uint(intID)
	user, err := h.userRepo.DeleteUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, user)

}