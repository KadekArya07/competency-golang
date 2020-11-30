package model

import (
	"errors"

	"github.com/labstack/echo"
)

type CustomError struct {
	echo.Context
	Err error
}

func (c *CustomError) CathError() error {
	if c.Err != nil {
		return errors.New("Token Required ")
	}
	return nil
}
