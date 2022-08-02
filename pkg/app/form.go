package app

import (
	"gin-learn/pkg/e"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, e.INVALID_PARAMS
	}
	vaild := validation.Validation{}
	check, err := vaild.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, e.ERROR
	}
	if !check {
		MarkErrors(vaild.Errors)
		return http.StatusBadRequest, e.INVALID_PARAMS
	}
	return http.StatusOK, e.SUCCESS
}
