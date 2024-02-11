package ginres

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrRes struct {
	Err string `json:"err"`
}

func NewErr(err string) *ErrRes {
	return &ErrRes{
		Err: err,
	}
}

type Error interface {
	Reply()
}

type InternalServerError struct {
	c   *gin.Context
	err error
}

func (e *InternalServerError) Reply() {
	e.c.AbortWithStatusJSON(http.StatusInternalServerError, NewErr(e.err.Error()))
}

func NewInternalServerError(c *gin.Context, err error) *InternalServerError {
	return &InternalServerError{
		c:   c,
		err: err,
	}
}

type ValidationError struct {
	c   *gin.Context
	err error
}

func (e *ValidationError) Reply() {
	e.c.AbortWithStatusJSON(http.StatusBadRequest, NewErr(e.err.Error()))
}

func NewValidationErr(c *gin.Context, err error) *ValidationError {
	return &ValidationError{
		c:   c,
		err: err,
	}
}
