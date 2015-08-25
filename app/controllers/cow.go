package controllers

import "github.com/revel/revel"

type Cow struct {
	*revel.Controller
}

func (c Cow) Moo() revel.Result {
	return c.Render()
}
