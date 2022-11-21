package main

import (
	"github.com/gin-gonic/gin"
	"github.com/matumoto1234/go-clean-architecture/app/controller"
	"github.com/matumoto1234/go-clean-architecture/app/repository"
	"github.com/matumoto1234/go-clean-architecture/app/service"
	"github.com/matumoto1234/go-clean-architecture/app/usecase"
)

func main() {
	ur := repository.NewUserRepository()
	us := service.NewUserService(ur)
	uu := usecase.NewUserUseCase(us)
	uc := controller.NewUserController(uu)

	engine := gin.Default()

	v1Router := engine.Group("/v1")

	v1Router.GET("/user/:id", uc.GETUser)
	v1Router.POST("/user", uc.POSTUser)

	engine.Run(":8080")
}
