package handlers

import (
	"context"
	"library-web-api-go/api/helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Create[Ti any, To any](c *gin.Context, caller func(ctx context.Context, req *Ti) (*To, error)) {
	req := new(Ti)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helpers.GenerateBaseResponseWithValidationError(nil, false, helpers.ValidationError, err))
		return
	}

	res, err := caller(c, req)
	if err != nil {
		c.AbortWithStatusJSON(helpers.TranslateErrorToStatusCode(err),
			helpers.GenerateBaseResponseWithError(nil, false, helpers.InternalError, err))
		return
	}
	c.JSON(http.StatusCreated, helpers.GenerateBaseResponse(res, true, 0))
}

func Update[Ti any, To any](c *gin.Context, caller func(ctx context.Context, id int, req *Ti) (*To, error)) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	if id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound,
			helpers.GenerateBaseResponse(nil, false, helpers.ValidationError))
	}
	
	req := new(Ti)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helpers.GenerateBaseResponseWithValidationError(nil, false, helpers.ValidationError, err))
		return
	}

	res, err := caller(c, id, req)
	if err != nil {
		c.AbortWithStatusJSON(helpers.TranslateErrorToStatusCode(err),
			helpers.GenerateBaseResponseWithError(nil, false, helpers.InternalError, err))
		return
	}
	c.JSON(http.StatusOK, helpers.GenerateBaseResponse(res, true, 0))
}

func Delete(c *gin.Context, caller func(ctx context.Context, id int) error) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	if id == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound,
			helpers.GenerateBaseResponse(nil, false, helpers.ValidationError))
	}

	err := caller(c, id)
	if err != nil {
		c.AbortWithStatusJSON(helpers.TranslateErrorToStatusCode(err),
			helpers.GenerateBaseResponseWithError(nil, false, helpers.InternalError, err))
		return
	}
	c.JSON(http.StatusOK, helpers.GenerateBaseResponse(nil, true, 0))
}

func GetById[To any](c *gin.Context, caller func(ctx context.Context, id int) (*To, error)) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	if id == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, 
			helpers.GenerateBaseResponse(nil, false, helpers.ValidationError))
		return
	}

	res, err := caller(c, id)
	if err != nil {
		c.AbortWithStatusJSON(helpers.TranslateErrorToStatusCode(err),
			helpers.GenerateBaseResponseWithError(nil, false, helpers.InternalError, err))
		return
	}

	c.JSON(http.StatusOK, helpers.GenerateBaseResponse(res, true, 0))
}
