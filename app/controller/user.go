package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/matumoto1234/go-clean-architecture/app/usecase"
)

type UserController interface {
	GETUser(c *gin.Context)
	POSTUser(c *gin.Context)
}

type userControllerImpl struct {
	uu usecase.UserUseCase
}

func (uc userControllerImpl) GETUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	id := c.Param("id")
	user, err := uc.uu.GETUser(id)

	usecaseErr, ok := err.(*usecase.Error)
	if ok {
		switch usecaseErr.Kind {
		case usecase.ErrBadRequest:
			c.String(http.StatusBadRequest, "Bad Request")
			return
		case usecase.ErrInternalServerError:
			c.String(http.StatusInternalServerError, "Internal Server Error")
			return
		case usecase.ErrNotFound:
			c.String(http.StatusNotFound, "Not Found")
			return
		}
	}

	res, err := json.Marshal(user)
	if err != nil {
		c.String(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	c.JSON(http.StatusOK, res)
	return
}

func (uc userControllerImpl) POSTUser(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var body []byte
	c.Request.Body.Read(body)
	defer c.Request.Body.Close()

	// TODO: リクエストモデルをgoaかなんかで自動生成するようにする
	type userRequestModel struct {
		Name string `json:"name"`
	}

	var user userRequestModel
	err := json.Unmarshal(body, &user)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad Request")
		return
	}

	uc.uu.POSTUser(user.Name)

	c.String(http.StatusOK, "OK")
	return
}

func NewUserController(uu usecase.UserUseCase) UserController {
	return &userControllerImpl{
		uu: uu,
	}
}
