package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raudra/rate-limitor/rule-manager/src/services"
)

type Response struct {
	Data    interface{} `json:"data,omitempty"`
	Err     string      `json:"error,omitempty"`
	Success bool        `json:"success"`
	Status  int         `json:"-"`
}

func GetRules(c *gin.Context) {

	rules, err := services.GetRules()
	resp := generateResponse(rules, err)
	c.JSON(resp.Status, resp)
}

func generateResponse(data interface{}, err error) Response {
	var resp Response

	if err != nil {
		resp = Response{
			Success: false,
			Err:     fmt.Sprintf("%s", err),
			Status:  http.StatusUnprocessableEntity,
		}
	} else {
		resp = Response{
			Success: true,
			Data:    data,
			Status:  http.StatusOK,
		}
	}
	return resp

}
