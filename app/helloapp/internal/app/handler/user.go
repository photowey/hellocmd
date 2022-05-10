package handler

import (
	"net/http"
	"strconv"

	"github.com/photowey/hellocmd/app/helloapp/internal/app/service"
)

func GetUser(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "parameter error")
	}
	user, err := service.UserRepository.Get(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}
