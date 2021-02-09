package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yongDataScince/rest-api"
)

func (h *Handlers) signUp(c *gin.Context) {
	var input rest.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// h.services.Autorization
}

func (h *Handlers) signIn(c *gin.Context) {
	
}