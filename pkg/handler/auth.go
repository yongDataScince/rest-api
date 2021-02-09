package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yongDataScince/rest-api"
)

func (h *Handlers) signUp(c *gin.Context) {
	var input rest.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error()) // not correct user's data, stop all
		return
	}

	id, err := h.services.Autorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type SignInp struct {
	Usename  string  `json:"username" binding:"required"`
	Password string  `json:"password" binding:"required"`
}

func (h *Handlers) signIn(c *gin.Context) {
	var input SignInp

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error()) // not correct user's data, stop all
		return
	}

	token, err := h.services.Autorization.CreateTocken(input.Usename, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}