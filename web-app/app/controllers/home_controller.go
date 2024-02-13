package controllers

import (
	"github.com/go-raptor/raptor"
)

type HomeController struct {
	raptor.Controller
}

func (hc *HomeController) Root(c *raptor.Context) error {
	return c.Render("home/root", nil)
}
